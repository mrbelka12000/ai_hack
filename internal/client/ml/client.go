package ml

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type Client struct {
	hc     *http.Client
	apiURL string
	log    *slog.Logger
}

func NewClient(apiURL string, log *slog.Logger) *Client {
	return &Client{
		apiURL: apiURL,
		hc: &http.Client{
			Timeout: 30 * time.Second,
		},
		log: log,
	}
}

func (c *Client) do(ctx context.Context, in, out any, reqURL, reqMethod string) (int, error) {

	var (
		requestBody  []byte
		responseBody []byte
		err          error
		statusCode   int
	)

	defer func() {
		fields := map[string]interface{}{
			"request_body":   string(requestBody),
			"response_body":  string(responseBody),
			"status_code":    statusCode,
			"request_url":    reqURL,
			"request_method": reqMethod,
		}

		if err != nil {
			fields["error"] = err.Error()
		}

		c.log.Info("execute request", "fields", fields)
	}()

	requestBody, err = json.Marshal(in)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, reqMethod, reqURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if out != nil {
		if err := json.Unmarshal(responseBody, out); err != nil {
			return 0, err
		}
	}

	return statusCode, nil
}
