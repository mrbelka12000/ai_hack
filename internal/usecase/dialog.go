package usecase

import (
	"context"
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

func (uc *UseCase) DialogCreate(ctx context.Context, obj internal.DialogCU) (uuid.UUID, error) {
	dialogID, err := uc.dialogService.Create(ctx, obj)
	if err != nil {
		return uuid.Nil, err
	}

	dmObj := internal.DialogMessage{
		DialogID:  dialogID,
		Role:      aihack.RoleClient,
		Message:   obj.Message,
		CreatedAt: time.Now().UTC(),
	}

	_, err = uc.dialogsMessagesService.AddMessage(ctx, dmObj)
	if err != nil {
		return uuid.Nil, err
	}

	return dialogID, nil
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
	return uc.dialogsMessagesService.AddMessage(ctx, obj)
}
