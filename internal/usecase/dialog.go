package usecase

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	aihack "github.com/mrbelka12000/ai_hack"
	"github.com/mrbelka12000/ai_hack/internal"
)

const (
	dialogCachePrefix = "dialog_"
	defaultTTL        = time.Minute * 5
)

func (uc *UseCase) DialogGet(ctx context.Context, id uuid.UUID) (obj internal.Dialog, err error) {
	rawObject, ok := uc.cache.Get(dialogCachePrefix + id.String())
	if ok {

		if err = json.Unmarshal([]byte(rawObject), &obj); err != nil {
			uc.log.With("error", err).Error("failed to unmarshal dialog")
			return obj, err
		}

		return obj, nil
	}

	obj, err = uc.dialogService.Get(ctx, id)
	if err != nil {
		return internal.Dialog{}, err
	}

	obj.DialogsMessages, err = uc.dialogsMessagesService.GetMessagesByDialogID(ctx, id)
	if err != nil {
		return internal.Dialog{}, err
	}

	raw, err := json.Marshal(obj)
	if err == nil {
		if err = uc.cache.Set(dialogCachePrefix+id.String(), string(raw), defaultTTL); err != nil {
			uc.log.With("error", err).Info("cache set")
		}
	}

	return obj, nil
}

func (uc *UseCase) DialogDelete(ctx context.Context, id uuid.UUID) error {
	return uc.dialogService.Delete(ctx, id)
}

func (uc *UseCase) DialogUpdate(ctx context.Context, obj internal.Dialog) error {
	uc.cache.Delete(dialogCachePrefix + obj.ID.String())
	return uc.dialogService.Update(ctx, obj)
}

func (uc *UseCase) DialogCreate(ctx context.Context, obj internal.DialogCU) (id uuid.UUID, err error) {
	obj.ID = uuid.New()

	err = uc.dialogService.Create(ctx, obj)
	if err != nil {
		return uuid.Nil, err
	}

	response, err := uc.dialogsMessagesService.AddMessage(ctx, internal.DialogMessage{
		DialogID:   obj.ID,
		Role:       aihack.RoleClient,
		Message:    obj.Message,
		IsLoggedIn: obj.ClientID != 1,
		CreatedAt:  time.Now().UTC(),
	}, true)
	if err != nil {
		return uuid.Nil, err
	}

	rawData, err := json.Marshal(response)
	if err != nil {
		return uuid.Nil, err
	}

	dialog := internal.Dialog{
		ID:      obj.ID,
		RawData: rawData,
	}

	if err := uc.dialogService.Update(ctx, dialog); err != nil {
		return uuid.Nil, err
	}

	return obj.ID, nil
}

func (uc *UseCase) DialogList(ctx context.Context, pars internal.DialogPars) ([]internal.Dialog, error) {
	response, err := uc.dialogService.List(ctx, pars)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(response); i++ {
		dialogMessages, err := uc.dialogsMessagesService.GetMessagesByDialogID(ctx, response[i].ID)
		if err != nil {
			return nil, err
		}

		response[i].DialogsMessages = dialogMessages
	}

	return response, nil
}

func (uc *UseCase) DialogAddMessage(ctx context.Context, obj internal.DialogMessage) error {
	uc.cache.Delete(dialogCachePrefix + obj.DialogID.String())
	response, err := uc.dialogsMessagesService.AddMessage(ctx, obj, obj.Role == aihack.RoleClient)
	if err != nil {
		return err
	}

	if obj.Role == aihack.RoleOperator {
		return nil
	}

	rawData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	dialogObj := internal.Dialog{
		ID:      obj.DialogID,
		RawData: rawData,
	}

	if err = uc.dialogService.Update(ctx, dialogObj); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DialogFull(ctx context.Context, obj internal.DialogFull) (out uuid.UUID, err error) {
	user, err := uc.userService.Get(ctx, internal.UserGetPars{
		PhoneNumber: obj.PhoneNumber,
	})
	if err != nil {
		err = uc.UserCreate(ctx, internal.UserCU{
			PhoneNumber: obj.PhoneNumber,
			Code:        "1324",
			Role:        aihack.RoleClient,
			CreatedAt:   time.Now().UTC(),
		})
		if err != nil {
			return out, err
		}
	}

	user, err = uc.userService.Get(ctx, internal.UserGetPars{
		PhoneNumber: obj.PhoneNumber,
	})
	if err != nil {
		return out, err
	}

	dialogID := uuid.New()
	dialogCU := internal.DialogCU{
		ID:        dialogID,
		ClientID:  user.ID,
		Status:    aihack.DialogStatusOpen,
		CreatedAt: time.Now().UTC(),
	}

	if err = uc.dialogService.Create(ctx, dialogCU); err != nil {
		return out, err
	}

	dialogMessages := parseFullDialog(obj.Message, dialogID)

	for i, message := range dialogMessages {
		var needGenerate bool
		if i == len(dialogMessages)-1 {
			needGenerate = true
		}
		_, err = uc.dialogsMessagesService.AddMessage(ctx, message, needGenerate)
		if err != nil {
			return out, err
		}
	}

	return dialogID, nil
}

func parseFullDialog(message string, dialogID uuid.UUID) []internal.DialogMessage {
	var (
		result []internal.DialogMessage
	)

	// Use a regexp that matches **Role**: and splits content accordingly
	re := regexp.MustCompile(`\*\*(Оператор|Клиент)\*\*:\s*`)

	// Find all indices where roles are marked
	locs := re.FindAllStringIndex(message, -1)
	matches := re.FindAllStringSubmatch(message, -1)

	for i := 0; i < len(locs); i++ {
		start := locs[i][1]
		end := len(message)
		if i+1 < len(locs) {
			end = locs[i+1][0]
		}

		roleStr := strings.ToLower(matches[i][1])
		var role aihack.Role
		if roleStr == "оператор" {
			role = aihack.RoleOperator
		} else {
			role = aihack.RoleClient
		}

		result = append(result, internal.DialogMessage{
			DialogID:   dialogID,
			Role:       role,
			IsLoggedIn: true,
			Message:    strings.TrimSpace(message[start:end]),
			CreatedAt:  time.Now().UTC(),
		})
	}

	return result
}
