package dto

type UpdateUsuarioInput struct {
	Nome   string `json:"nome" example:"string"`
	Email  string `json:"email" example:"string"`
	Senha  string `json:"senha" example:"string"`
}
