package types

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginOutput struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
