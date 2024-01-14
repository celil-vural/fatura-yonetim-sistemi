package auth

type LoginDto struct {
	IdentityNumber string `json:"identity_number"`
	Password       string `json:"password"`
}
