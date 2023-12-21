package response_model

type ResponseModel struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
