package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/models"
	"github.com/maty24/Goapi.git/internal/services"
	"strconv"
)

type UsuarioController struct {
	userService *services.UsuarioService
}

func NewUsuarioController(userService *services.UsuarioService) *UsuarioController {
	return &UsuarioController{userService: userService}
}

// GetAllUsuarios returns all users
func (c *UsuarioController) GetAllUsuarios(ctx *gin.Context) {
	users, err := c.userService.GetAllUsuarios()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	ctx.JSON(200, users)
}

// GetUsuarioByID returns a user by its ID
func (c *UsuarioController) GetUsuarioByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := c.userService.GetUsuarioByID(uint(id))
	if err != nil {
		if err.Error() == "usuario con ID no encontrado" {
			ctx.JSON(404, gin.H{"error": "Usuario no encontrado"})
			return
		}
		ctx.JSON(500, gin.H{"error": "Error al obtener usuario"})
		return
	}

	ctx.JSON(200, usuario)
}

// CreateUsuario creates a new user
func (c *UsuarioController) CreateUsuario(ctx *gin.Context) {
	var usuario models.Usuario
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	if err := c.userService.CreateUsuario(&usuario); err != nil {
		ctx.JSON(500, gin.H{"error": "Error al crear usuario"})
		return
	}

	ctx.JSON(201, usuario)
}

// LoginUsuario logs in a user and returns a JWT token
func (c *UsuarioController) LoginUsuario(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	token, err := c.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}
