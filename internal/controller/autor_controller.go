package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/models"
	"github.com/maty24/Goapi.git/internal/services"
	"strconv"
)

type AutorController struct {
	autorService *services.AutorService
}

func NewAutorController(autorService *services.AutorService) *AutorController {
	return &AutorController{autorService: autorService}
}

// GetAllAutores returns all authors
func (c *AutorController) GetAllAutores(ctx *gin.Context) {
	autores, err := c.autorService.GetAllAutores()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al obtener autores"})
		return
	}

	ctx.JSON(200, autores)
}

// GetAutorByID returns an author by its ID
func (c *AutorController) GetAutorByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	autor, err := c.autorService.GetAutorByID(uint(id))
	if err != nil {
		if err.Error() == "autor con ID no encontrado" {
			ctx.JSON(404, gin.H{"error": fmt.Sprintf("Autor con ID %d no encontrado", id)})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, autor)
}

// CreateAutor creates a new author
func (c *AutorController) CreateAutor(ctx *gin.Context) {
	var autor models.Autor
	if err := ctx.ShouldBindJSON(&autor); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	if err := c.autorService.CreateAutor(&autor); err != nil {
		ctx.JSON(500, gin.H{"error": "Error al crear autor"})
		return
	}

	// Retornar la URL del nuevo recurso con el código 201
	ctx.JSON(201, gin.H{
		"message": "Autor creado con éxito",
		"autor":   autor,
	})
}

// UpdateAutor updates an author
func (c *AutorController) UpdateAutor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var autor models.Autor
	if err := ctx.ShouldBindJSON(&autor); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	autor.ID = uint(id) // Asegurar que el ID viene de la URL y no del JSON
	if err := c.autorService.UpdateAutor(&autor); err != nil {
		if err.Error() == fmt.Sprintf("autor con ID %d no encontrado", id) {
			ctx.JSON(404, gin.H{"error": fmt.Sprintf("Autor con ID %d no encontrado", id)})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Autor actualizado con éxito", "autor": autor})
}
