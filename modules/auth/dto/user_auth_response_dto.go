package dto

type UserAuthResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
	Token    string `json:"token"`
	Exp      int64  `json:"expires"`
}
