package dto

type ClientRequest struct {
	ClientName        string `json:"client_name" validate:"required"`
	ClientDescription string `json:"client_description"`
}
