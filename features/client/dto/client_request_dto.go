package dto

type ClientRequest struct {
	ClientName        string `json:"client_name" validate:"required,alphanum_extra"`
	ClientDescription string `json:"client_description" validate:"omitempty,alphanum_extra"`
}
