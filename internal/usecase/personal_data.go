package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
)

func (uc *UseCase) GetPersonalData(ctx context.Context, obj internal.PersonalDataRequest) (internal.PersonalDataResponse, error) {
	id, err := uuid.Parse(obj.DialogId)
	if err != nil {
		return internal.PersonalDataResponse{}, err
	}

	dialog, err := uc.DialogGet(ctx, id)
	if err != nil {
		return internal.PersonalDataResponse{}, err
	}

	user, err := uc.UserGet(ctx, internal.UserGetPars{
		ID: dialog.ClientID,
	})
	if err != nil {
		return internal.PersonalDataResponse{}, err
	}

	obj.PhoneNumber = user.PhoneNumber
	obj.CallID = user.PhoneNumber

	return uc.personalData.GetPersonalDataForResponse(ctx, obj)
}
