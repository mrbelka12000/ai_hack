package ml

import (
	"context"
	"fmt"
	"net/http"
)

const (
	analyzeEndPoint = "/analyze"
)

type (
	AnalyzeRequest struct {
		DialogId string `json:"dialog_id"`
		Dialog   string `json:"dialog"`
		LoggedIn bool   `json:"logged_in"`
	}

	AnalyzeResponse struct {
		DialogId          string   `json:"dialog_id"`
		Message           string   `json:"message"`
		RelativeQuestions []string `json:"relative_questions"`
		DatabaseFile      string   `json:"database_file"`
		DatabaseFilePart  string   `json:"database_file_part"`
		Confidence        int      `json:"confidence"`
	}
)

func (c *Client) Analyze(ctx context.Context, req AnalyzeRequest) (out AnalyzeResponse, err error) {

	statusCode, err := c.do(ctx, req, &out, c.apiURL+analyzeEndPoint, http.MethodPost)
	if err != nil {
		return out, err
	}

	if statusCode >= 400 {
		return out, fmt.Errorf("status code %d", statusCode)
	}

	return out, nil
}
