package suggestions

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

		if len(val) < 51 {
			fmt.Println("Skipping row due to invalid len:", val)
			continue
		}

		mb := internal.MB{
			ID:         trimQuotes(val[0]),
			CustID:     trimQuotes(val[9]),
			Acct:       trimQuotes(val[12]),
			Br:         trimQuotes(val[13]),
			Segment:    trimQuotes(val[15]),
			Product:    trimQuotes(val[16]),
			ContCode:   trimQuotes(val[17]),
			ContType:   trimQuotes(val[18]),
			DocNum:     ptr(trimQuotes(val[19])),
			SubsLoanTo: ptr(trimQuotes(val[25])),
			LineType:   ptr(trimQuotes(val[29])),
			EndDate:    ptr(trimQuotes(val[33])),
			AmtTng:     ptr(trimQuotes(val[37])),
			OdTng:      ptr(trimQuotes(val[41])),
			Stav:       ptr(trimQuotes(val[48])),
			DayPrPr:    ptr(trimQuotes(val[51])),
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
