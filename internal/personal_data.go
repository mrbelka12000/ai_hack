package internal

import "github.com/google/uuid"

type (
	PersonalData struct {
		CallID        string `gorm:"column:call_id"`
		PhoneNumber   string `gorm:"column:phone_number"`
		Br            string `gorm:"column:br"`
		Currency      string `gorm:"column:currency"`
		BegDate       string `gorm:"column:beg_date"`
		EndDate       string `gorm:"column:end_date"`
		ProlDate      string `gorm:"column:prol_date"`
		ProlCount     string `gorm:"column:prol_count"`
		Amt           string `gorm:"column:amt"`
		AmtTng        string `gorm:"column:amt_tng"`
		Od            string `gorm:"column:od"`
		PrOd          string `gorm:"column:pr_od"`
		DayPrOd       string `gorm:"column:day_pr_od"`
		Pog           string `gorm:"column:pog"`
		Stav          string `gorm:"column:stav"`
		Sht           string `gorm:"column:sht"`
		BrVyd         string `gorm:"column:br_vyd"`
		FlWork        string `gorm:"column:flwork"`
		RateEffective string `gorm:"column:rate_effective"`
	}

	PersonalDataRequest struct {
		DialogId    uuid.UUID `json:"dialog_id"`
		DataType    string    `json:"data_type"`
		CallID      string    `json:"-"`
		PhoneNumber string    `json:"-"`
	}

	PersonalDataPars struct {
		CallID      string `gorm:"column:call_id"`
		PhoneNumber string `gorm:"column:phone_number"`
	}

	PersonalDataResponse struct {
		Result any `json:"result"`
	}
)
