package model

type ResponseMessage struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	DisplayMessage string `json:"display_message"`
}
