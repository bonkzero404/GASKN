package dto

type ClientRequest struct {
	ClientName        string `json:"client_name" validate:"required,alphanum"`
	ClientDescription string `json:"client_description" validate:"alphanum"`
}
