package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/MarcosVerse/nami/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
    // rota de teste
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    usuarios := r.Group("/usuarios")
    {
        usuarios.POST("/", controllers.CriarUsuario)    
        usuarios.PUT("/:id", controllers.AtualizarUsuario) 
        usuarios.DELETE("/:id", controllers.DeletarUsuario) 
    }

    // login
    r.POST("/login", controllers.Login)
}