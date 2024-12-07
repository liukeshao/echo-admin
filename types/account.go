package types

type LoginInput struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}
type LoginOutput struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}
