package dto

type LoginInput struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"` //uso futuro
}
