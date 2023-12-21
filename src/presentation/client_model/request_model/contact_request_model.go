package request_model

type ContactRequestModel struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}
