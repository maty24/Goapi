package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/services"
)

type AutorController struct {
	autorService *services.AutorService
}

func NewAutorController(autorService *services.AutorService) *AutorController {
	return &AutorController{autorService: autorService}
}

// GetAllAutores returns all authors
func (c *AutorController) GetAllAutores(ctx *gin.Context) {
	autorCompleto, err := c.autorService.GetAllAutores()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if autorCompleto == nil {
		ctx.JSON(404, gin.H{"error": "not found"})
		return
	}

	ctx.JSON(200, autorCompleto)
}
