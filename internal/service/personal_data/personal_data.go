package personal_data

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mrbelka12000/ai_hack/internal"
)

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) StartParseMB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.Read()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Skipping row due to error:", err)
			continue
		}

		row := strings.Split(record[0], ";")

		mb := internal.PersonalData{
			CallID:        trimQuotes(getValue(row, 0)),
			PhoneNumber:   "-",
			Br:            trimQuotes(getValue(row, 13)),
			Currency:      trimQuotes(getValue(row, 31)),
			BegDate:       trimQuotes(getValue(row, 32)),
			EndDate:       trimQuotes(getValue(row, 33)),
			ProlDate:      trimQuotes(getValue(row, 34)),
			ProlCount:     trimQuotes(getValue(row, 35)),
			Amt:           trimQuotes(getValue(row, 36)),
			AmtTng:        trimQuotes(getValue(row, 37)),
			Od:            trimQuotes(getValue(row, 38)),
			PrOd:          trimQuotes(getValue(row, 39)),
			DayPrOd:       trimQuotes(getValue(row, 40)),
			Pog:           trimQuotes(getValue(row, 47)),
			Stav:          trimQuotes(getValue(row, 48)),
			Sht:           trimQuotes(getValue(row, 57)),
			BrVyd:         trimQuotes(getValue(row, 84)),
			FlWork:        trimQuotes(getValue(row, 88)),
			RateEffective: trimQuotes(getValue(row, 139)),
		}

		if err := s.repo.Create(context.Background(), mb); err != nil {
			fmt.Println("Skipping row due to error:", err)
			continue
		}
	}

	return nil
}

func (s *Service) StartParseRB(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1
	reader.Read()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Skipping row due to error:", err)
			continue
		}

		row := strings.Split(record[0], ";")

		mb := internal.PersonalData{
			CallID:        "-",
			PhoneNumber:   trimQuotes(getValue(row, 0)),
			Br:            trimQuotes(getValue(row, 13)),
			Currency:      trimQuotes(getValue(row, 31)),
			BegDate:       trimQuotes(getValue(row, 32)),
			EndDate:       trimQuotes(getValue(row, 33)),
			ProlDate:      trimQuotes(getValue(row, 34)),
			ProlCount:     trimQuotes(getValue(row, 35)),
			Amt:           trimQuotes(getValue(row, 36)),
			AmtTng:        trimQuotes(getValue(row, 37)),
			Od:            trimQuotes(getValue(row, 38)),
			PrOd:          trimQuotes(getValue(row, 39)),
			DayPrOd:       trimQuotes(getValue(row, 40)),
			Pog:           trimQuotes(getValue(row, 47)),
			Stav:          trimQuotes(getValue(row, 48)),
			Sht:           trimQuotes(getValue(row, 57)),
			BrVyd:         trimQuotes(getValue(row, 84)),
			FlWork:        trimQuotes(getValue(row, 88)),
			RateEffective: trimQuotes(getValue(row, 139)),
		}

		if err := s.repo.Create(context.Background(), mb); err != nil {
			fmt.Println("Skipping row due to error:", err)
			continue
		}
	}

	return nil
}

func (s *Service) GetPersonalDataForResponse(ctx context.Context, obj internal.PersonalDataRequest) (internal.PersonalDataResponse, error) {
	result, err := s.repo.GetPersonalDataForResponse(ctx, obj)
	if err != nil {
		return internal.PersonalDataResponse{}, err
	}

	return internal.PersonalDataResponse{
		Message: result,
	}, nil
}

func trimQuotes(s string) string {
	return strings.Trim(s, `'`)
}

func getValue(arr []string, ind int) string {
	if ind >= len(arr) {
		return ""
	}

	return arr[ind]
}
