package controllers

import (
	"net/http"
	"strconv"

	"github.com/MarcosVerse/nami/internal/dto"
	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// CriarUsuario cria um novo usuário
// @Summary Cria um usuário
// @Tags Usuários
// @Accept json
// @Produce json
// @Param usuario body dto.CreateUsuarioInput true "Usuário"
// @Success 201 {object} models.Usuario
// @Failure 400 {object} dto.ResponseMessage
// @Failure 500 {object} dto.ResponseMessage
// @Router /usuarios [post]
func CriarUsuario(c *gin.Context) {
	var input dto.CreateUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.BadRequest})
		return
	}

	usuario := models.Usuario{
		Nome:  input.Nome,
		Email: input.Email,
		Senha: input.Senha, // futuramente criptografar
	}

	if err := repository.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

// AtualizarUsuario atualiza os dados de um usuário
// @Summary Atualiza um usuário
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param usuario body dto.UpdateUsuarioInput true "Dados para atualização"
// @Success 200 {object} dto.ResponseMessage
// @Failure 400 {object} dto.ResponseMessage
// @Failure 404 {object} dto.ResponseMessage
// @Failure 500 {object} dto.ResponseMessage
// @Router /usuarios/{id} [put]
func AtualizarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.InvalidUserID})
		return
	}

	var usuario models.Usuario
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ResponseMessage{Message: config.UserNotFound})
		return
	}

	var input dto.UpdateUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.BadRequest})
		return
	}

	usuario.Nome = input.Nome
	usuario.Email = input.Email
	if input.Senha != "" {
		usuario.Senha = input.Senha // futuramente criptografar
	}

	if err := repository.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Message: config.UserUpdated})
}

// DeletarUsuario remove um usuário
// @Summary Deleta um usuário
// @Tags Usuários
// @Param id path int true "ID do usuário"
// @Success 200 {object} dto.ResponseMessage
// @Failure 404 {object} dto.ResponseMessage
// @Failure 500 {object} dto.ResponseMessage
// @Router /usuarios/{id} [delete]
func DeletarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := repository.DB.Delete(&models.Usuario{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, dto.ResponseMessage{Message: config.UserNotFound})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Message: config.UserDeleted})
}
