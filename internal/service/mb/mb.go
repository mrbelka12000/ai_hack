package mb

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

func (s *Service) StartParse(filePath string) error {
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

		val := strings.Split(record[0], ";")

		mb := internal.MB{
			ID:         trimQuotes(getValue(val, 0)),
			CustID:     trimQuotes(getValue(val, 9)),
			Acct:       trimQuotes(getValue(val, 12)),
			Br:         trimQuotes(getValue(val, 13)),
			Segment:    trimQuotes(getValue(val, 15)),
			Product:    trimQuotes(getValue(val, 16)),
			ContCode:   trimQuotes(getValue(val, 17)),
			ContType:   trimQuotes(getValue(val, 18)),
			DocNum:     ptr(trimQuotes(getValue(val, 19))),
			SubsLoanTo: ptr(trimQuotes(getValue(val, 25))),
			LineType:   ptr(trimQuotes(getValue(val, 29))),
			EndDate:    ptr(trimQuotes(getValue(val, 33))),
			AmtTng:     ptr(trimQuotes(getValue(val, 37))),
			OdTng:      ptr(trimQuotes(getValue(val, 41))),
			Stav:       ptr(trimQuotes(getValue(val, 48))),
			DayPrPr:    ptr(trimQuotes(getValue(val, 51))),
		}

		if err := s.repo.Create(context.Background(), mb); err != nil {
			fmt.Println("Skipping row due to error:", err)
			continue
		}
	}

	return nil
}

func trimQuotes(s string) string {
	return strings.Trim(s, `'`)
}

func ptr(s string) *string {
	return &s
}

func getValue(arr []string, ind int) string {
	if ind >= len(arr) {
		return ""
	}

	return arr[ind]
}
