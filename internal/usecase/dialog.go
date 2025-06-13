package usecase

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"

	aihack "github.com/mrbelka12000/ai_hack"
	"github.com/mrbelka12000/ai_hack/internal"
)

func (uc *UseCase) DialogGet(ctx context.Context, id uuid.UUID) (internal.Dialog, error) {
	return uc.dialogService.Get(ctx, id)
}

func (uc *UseCase) DialogDelete(ctx context.Context, id uuid.UUID) error {
	return uc.dialogService.Delete(ctx, id)
}

func (uc *UseCase) DialogUpdate(ctx context.Context, obj internal.Dialog) error {
	return uc.dialogService.Update(ctx, obj)
}

func (uc *UseCase) DialogCreate(ctx context.Context, obj internal.DialogCU) (out internal.DialogMessageResponse, err error) {
	dialogID, err := uc.dialogService.Create(ctx, obj)
	if err != nil {
		return out, err
	}

	dmObj := internal.DialogMessage{
		DialogID:  dialogID,
		Role:      aihack.RoleClient,
		Message:   obj.Message,
		CreatedAt: time.Now().UTC(),
	}

	response, err := uc.dialogsMessagesService.AddMessage(ctx, dmObj, true)
	if err != nil {
		return out, err
	}

	response.DialogID = dialogID

	return response, nil
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

func (uc *UseCase) DialogAddMessage(ctx context.Context, obj internal.DialogMessage) (internal.DialogMessageResponse, error) {
	return uc.dialogsMessagesService.AddMessage(ctx, obj, obj.Role == aihack.RoleClient)
}

func (uc *UseCase) DialogFull(ctx context.Context, obj internal.DialogFull) (out internal.DialogMessageResponse, err error) {
	user, err := uc.userService.Get(ctx, internal.UserGetPars{
		PhoneNumber: obj.PhoneNumber,
	})
	if err != nil {
		return out, err
	}

	dialogID, err := uc.dialogService.Create(ctx, internal.DialogCU{
		ClientID:  user.ID,
		Status:    aihack.DialogStatusOpen,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return out, err
	}

	dialogMessages := parseFullDialog(obj.Message, dialogID)

	var response internal.DialogMessageResponse
	for _, message := range dialogMessages {
		response, err = uc.dialogsMessagesService.AddMessage(ctx, message, message.Role == aihack.RoleClient)
		if err != nil {
			return out, err
		}
		response.DialogID = message.DialogID
	}

	return response, nil
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
			DialogID:  dialogID,
			Role:      role,
			Message:   strings.TrimSpace(message[start:end]),
			CreatedAt: time.Now().UTC(),
		})
	}

	return result
}
