package controllers

import (
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/MarcosVerse/nami/internal/database"
    "github.com/MarcosVerse/nami/internal/models"
    "golang.org/x/crypto/bcrypt"
)

// CriarUsuario cria um novo usuário
// @Summary Cria um usuário
// @Tags usuários
// @Accept json
// @Produce json
// @Param usuario body models.Usuario true "Usuário"
// @Success 201 {object} models.Usuario
// @Router /usuarios [post]
func CriarUsuario(c *gin.Context) {
    var usuario models.Usuario
    if err := c.ShouldBindJSON(&usuario); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 🔒 Criptografar a senha
    hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
        return
    }
    usuario.Senha = string(hash)

    // Definir timestamps
    usuario.CriadoEm = time.Now()
    usuario.AtualizadoEm = time.Now()

    if err := database.DB.Create(&usuario).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, usuario)
}

// DeletarUsuario remove um usuário
// @Summary Deleta um usuário
// @Description Remove usuário do banco pelo ID
// @Tags usuários
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /usuarios/{id} [delete]
func DeletarUsuario(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := database.DB.Delete(&models.Usuario{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
