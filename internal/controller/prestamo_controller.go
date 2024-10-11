package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/models"
	"github.com/maty24/Goapi.git/internal/services"
	"net/http"
	"strconv"
)

type PrestamoController struct {
	prestamoService *services.PrestamoService
}

func NewPrestamoController(prestamoService *services.PrestamoService) *PrestamoController {
	return &PrestamoController{prestamoService: prestamoService}
}

// GetActivePrestamosByLectorID returns active loans filtered by lectorID
func (c *PrestamoController) GetActivePrestamosByLectorID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	prestamos, err := c.prestamoService.GetActivePrestamosByLectorID(uint(id))
	if err != nil {
		if err.Error() == "lector con ID no encontrado" {
			ctx.JSON(404, gin.H{"error": "Lector no encontrado"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, prestamos)
}

// CreatePrestamo creates a new loan if the book is available
func (c *PrestamoController) CreatePrestamo(ctx *gin.Context) {
	var prestamo models.Prestamo
	if err := ctx.ShouldBindJSON(&prestamo); err != nil {
		ctx.JSON(422, gin.H{"error": "Datos inválidos", "details": err.Error()})
		return
	}

	if err := c.prestamoService.CreatePrestamo(&prestamo); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, prestamo)
}

// getPrestamoById
func (c *PrestamoController) GetPrestamoByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	prestamo, err := c.prestamoService.GetPrestamoByID(uint(id))
	if err != nil {
		if err.Error() == "préstamo con ID no encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Préstamo no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, prestamo)
}

// ReturnPrestamo handles the request to return a loan
func (c *PrestamoController) ReturnPrestamo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.prestamoService.ReturnPrestamo(uint(id)); err != nil {
		if err.Error() == "no se encontró un préstamo pendiente para este libro" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Préstamo devuelto con éxito"})
}
