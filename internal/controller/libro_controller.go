package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/models"
	"github.com/maty24/Goapi.git/internal/services"
	"strconv"
)

type LibroController struct {
	libroService *services.LibroService
}

func NewLibroController(libroService *services.LibroService) *LibroController {
	return &LibroController{libroService: libroService}
}

// GetAllLibros returns all books
func (c *LibroController) GetAllLibros(ctx *gin.Context) {
	libros, err := c.libroService.GetAllLibros()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al obtener libros"})
		return
	}

	ctx.JSON(200, libros)
}

// GetLibroByID returns a book by its ID
func (c *LibroController) GetLibroByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	libro, err := c.libroService.GetLibroByID(uint(id))
	if err != nil {
		if err.Error() == "libro con ID no encontrado" {
			ctx.JSON(404, gin.H{"error": "Libro no encontrado"})
			return
		}
		ctx.JSON(500, gin.H{"error": "Error al obtener libro"})
		return
	}

	ctx.JSON(200, libro)
}

// CreateLibro creates a new book
func (c *LibroController) CreateLibro(ctx *gin.Context) {
	var libro models.Libro
	if err := ctx.ShouldBindJSON(&libro); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	if err := c.libroService.CreateLibro(&libro); err != nil {
		// Log the error details
		fmt.Printf("Error details: %v\n", err)
		ctx.JSON(500, gin.H{"error": "Error al crear libro", "details": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Libro creado exitosamente",
		"libro":   libro,
	})
}

// UpdateLibro updates a book
func (c *LibroController) UpdateLibro(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var libro models.Libro
	if err := ctx.ShouldBindJSON(&libro); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	libro.ID = uint(id)
	if err := c.libroService.UpdateLibro(&libro); err != nil {
		ctx.JSON(500, gin.H{"error": "Error al actualizar libro"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Libro actualizado exitosamente",
		"libro":   libro,
	})
}
