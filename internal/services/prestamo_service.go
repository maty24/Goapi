package services

import (
	"errors"
	"github.com/maty24/Goapi.git/internal/models"
	"gorm.io/gorm"
	"time"
)

type PrestamoService struct {
	DB *gorm.DB
}

func NewPrestamoService(db *gorm.DB) *PrestamoService {
	return &PrestamoService{DB: db}
}

// GetActivePrestamosByLectorID returns active loans filtered by lectorID
func (s *PrestamoService) GetActivePrestamosByLectorID(lectorID uint) ([]models.PrestamoResponse, error) {
	var prestamos []models.PrestamoResponse

	if err := s.DB.Preload("Libro").Preload("Libro.Categoria").
		Where("lector_id = ? AND estado = ?", lectorID, "pendiente").
		Find(&prestamos).Error; err != nil {
		return nil, err
	}

	return prestamos, nil
}

// GetPendingPrestamos returns loans that are in a pending state, with a default limit of 20
func (s *PrestamoService) GetPendingPrestamos(limit int) ([]models.Prestamo, error) {
	if limit <= 0 {
		limit = 20
	}

	var prestamos []models.Prestamo
	if err := s.DB.Preload("Libro").Preload("Libro.Categoria").Preload("Lector").
		Where("estado = ?", "pendiente").
		Limit(limit).
		Find(&prestamos).Error; err != nil {
		return nil, err
	}

	return prestamos, nil
}

// GetPrestamoByID returns a loan by its ID
func (s *PrestamoService) GetPrestamoByID(id uint) (*models.Prestamo, error) {
	var prestamo models.Prestamo

	if err := s.DB.Preload("Libro").Preload("Lector").First(&prestamo, id).Error; err != nil {
		return nil, err
	}

	return &prestamo, nil
}

// CreatePrestamo crea un nuevo préstamo si el libro está disponible
func (s *PrestamoService) CreatePrestamo(prestamo *models.Prestamo) error {
	var libro models.Libro

	// Verificar si el libro está disponible
	if err := s.DB.First(&libro, prestamo.LibroID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("libro no encontrado")
		}
		return err
	}

	if !libro.Disponible {
		return errors.New("el libro no está disponible")
	}

	// Crear el préstamo
	if err := s.DB.Create(prestamo).Error; err != nil {
		return err
	}

	// Marcar el libro como no disponible
	libro.Disponible = false
	if err := s.DB.Save(&libro).Error; err != nil {
		return err
	}

	return nil
}

// ReturnPrestamo updates the loan status to returned and the book status to available
func (s *PrestamoService) ReturnPrestamo(libroID uint) error {
	var prestamo models.Prestamo

	// Verificar si el préstamo está en estado pendiente
	if err := s.DB.Where("libro_id = ? AND estado = ?", libroID, "pendiente").First(&prestamo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no se encontró un préstamo pendiente para este libro")
		}
		return err
	}

	// Actualizar el estado del préstamo a devuelto
	prestamo.Estado = "devuelto"
	prestamo.FechaDevuelto = time.Now()
	if err := s.DB.Save(&prestamo).Error; err != nil {
		return err
	}

	// Actualizar el estado del libro a disponible
	var libro models.Libro
	if err := s.DB.First(&libro, libroID).Error; err != nil {
		return err
	}
	libro.Disponible = true
	if err := s.DB.Save(&libro).Error; err != nil {
		return err
	}

	return nil
}
