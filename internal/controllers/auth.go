package controllers

import (
	"net/http"
	"time"

	"github.com/MarcosVerse/nami/internal/dto"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login autentica um usuário (simplificado)
// @Summary Realiza login do usuário
// @Description Autentica o usuário com email e senha e retorna um token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginInput true "Credenciais de login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.LoginResponse
// @Failure 401 {object} dto.LoginResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.LoginResponse{Message: "Dados inválidos"})
		return
	}

	var usuario models.Usuario 
	if err := repository.DB.Where("email = ? AND senha = ?", input.Email, input.Senha).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, dto.LoginResponse{Message: "Email ou senha incorretos"})
		return
	}

	// aqui gera o token jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  usuario.ID,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SUA_CHAVE_SECRETA"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.LoginResponse{Message: "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login realizado com sucesso",
		Token:   tokenString,
	})
}
