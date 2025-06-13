package ml

type (
	PersonalDataRequest struct {
		DialogId string `json:"dialog_id"`
		DataType string `json:"data_type"`
	}

	PersonalData struct {
		Balance string `json:"balance"`
	}
)
