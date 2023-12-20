package request_model

import "encoding/json"

type ContactRequestModel struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

func ConvertContactRequestModel(inputs string) (*ContactRequestModel, error) {
	var req ContactRequestModel
	err := json.Unmarshal([]byte(inputs), &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
