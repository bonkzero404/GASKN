package dto

type UserClient struct {
	ClientId        string `json:"client_id"`
	ClientShortName string `json:"client_short_name"`
	ClientName      string `json:"client_name"`
	RoleName        string `json:"role_name"`
}

type UserAuthProfileResponse struct {
	ID       string       `json:"id"`
	FullName string       `json:"full_name"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	IsActive bool         `json:"is_active"`
	Clients  []UserClient `json:"clients"`
}
