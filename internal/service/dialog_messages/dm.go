package dialog_messages

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/internal/client/ml"
	"github.com/mrbelka12000/ai_hack/pkg/validator"
)

const (
	defaultTTL = 30 * time.Minute
)

type Service struct {
	repo     repo
	mlClient mlClient
	cache    cache
}

func NewService(repo repo, mlClient mlClient, cache cache) *Service {
	return &Service{
		repo:     repo,
		mlClient: mlClient,
		cache:    cache,
	}
}

func (s *Service) AddMessage(ctx context.Context, obj internal.DialogMessage, needToGenerate bool) (out internal.DialogMessageResponse, err error) {
	if err = validator.ValidateStruct(obj); err != nil {
		return out, err
	}

	if err = s.repo.AddMessage(ctx, obj); err != nil {
		return out, err
	}

	messageFromCache, _ := s.cache.Get(obj.DialogID.String())
	messageToSave := getMessageToSaveInRedis(obj)

	fullDialog := constructDialog(messageFromCache, messageToSave)

	var resp ml.AnalyzeResponse

	if needToGenerate {
		resp, err = s.mlClient.Analyze(ctx, ml.AnalyzeRequest{
			DialogId: obj.DialogID.String(),
			Dialog:   fullDialog,
			LoggedIn: obj.IsLoggedIn,
		})
		if err != nil {
			return out, err
		}
	}

	if err = s.cache.Set(obj.DialogID.String(), fullDialog, defaultTTL); err != nil {
		return out, err
	}

	if resp.Error != "" {
		resp.Message = resp.Error
	}

	return internal.DialogMessageResponse{
		Message:           resp.Message,
		RelativeQuestions: resp.RelativeQuestions,
		DatabaseFile:      resp.DatabaseFile,
		DatabaseFilePart:  resp.DatabaseFilePart,
		Confidence:        resp.Confidence,
	}, nil
}

func (s *Service) GetMessagesByDialogID(ctx context.Context, dialogID uuid.UUID) ([]internal.DialogMessage, error) {
	return s.repo.GetMessagesByDialogID(ctx, dialogID)
}

func getMessageToSaveInRedis(obj internal.DialogMessage) string {
	return fmt.Sprintf("[%v] %v", obj.Role, obj.Message)
}

func constructDialog(oldMsg, newMsg string) string {
	if oldMsg == "" {
		return newMsg
	}

	return oldMsg + newMsg
}
