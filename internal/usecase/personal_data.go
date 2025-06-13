package usecase

import (
	"context"

	"github.com/mrbelka12000/ai_hack/internal"
)

func (uc *UseCase) GetPersonalData(ctx context.Context, obj internal.PersonalDataRequest) (internal.PersonalDataResponse, error) {
	dialog, err := uc.DialogGet(ctx, obj.DialogId)
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
