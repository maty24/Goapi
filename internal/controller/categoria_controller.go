package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/models"
	"github.com/maty24/Goapi.git/internal/services"
	"strconv"
)

type CategoriaController struct {
	categoriaService *services.CategoriaService
}

func NewCategoriaController(categoriaService *services.CategoriaService) *CategoriaController {
	return &CategoriaController{categoriaService: categoriaService}
}

// GetAllCategorias returns all categories
func (c *CategoriaController) GetAllCategorias(ctx *gin.Context) {
	categorias, err := c.categoriaService.GetAllCategorias()

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al obtener categorias"})
		return
	}

	ctx.JSON(200, categorias)
}

// GetCategoriaByID returns a category by its ID
func (c *CategoriaController) GetCategoriaByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	categoria, err := c.categoriaService.GetCategoriaByID(uint(id))
	if err != nil {
		if err.Error() == "categoria con ID no encontrada" {
			ctx.JSON(404, gin.H{"error": "Categoria no encontrada"})
			return
		}
		ctx.JSON(500, gin.H{"error": "Error al obtener categoria"})
		return
	}

	ctx.JSON(200, categoria)

}

// CreateCategoria creates a new category
func (c *CategoriaController) CreateCategoria(ctx *gin.Context) {
	var categoria models.Categoria
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	if err := c.categoriaService.CreateCategoria(&categoria); err != nil {
		ctx.JSON(500, gin.H{"error": "Error al crear categoria"})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Categoria creada exitosamente",
	})

}

// UpdateCategoria updates a category
func (c *CategoriaController) UpdateCategoria(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var categoria models.Categoria
	if err := ctx.ShouldBindJSON(&categoria); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	categoria.ID = uint(id)
	if err := c.categoriaService.UpdateCategoria(&categoria); err != nil {
		ctx.JSON(500, gin.H{"error": "Error al actualizar categoria"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Categoria actualizada exitosamente",
	})
}
