package models

type ResponseModel struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}
