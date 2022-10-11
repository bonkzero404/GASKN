package dto

type ClientResponse struct {
	ID                string `json:"id"`
	ClientName        string `json:"client_name"`
	ClientDescription string `json:"client_description"`
	ClientSlug        string `json:"client_slug"`
	IsActive          bool   `json:"is_active"`
}
