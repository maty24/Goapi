package services

import (
	"errors"
	"fmt"
	"github.com/maty24/Goapi.git/internal/models"
	"gorm.io/gorm"
)

type LibroService struct {
	db *gorm.DB
}

func NewLibroService(db *gorm.DB) *LibroService {
	return &LibroService{db: db}
}

// GetAllLibros obtiene todos los libros de la base de datos
func (s *LibroService) GetAllLibros() (*[]models.LibroResponse, error) {
	var libros []models.LibroResponse

	if err := s.db.Preload("Categoria").Find(&libros).Error; err != nil {
		return nil, err
	}

	return &libros, nil
}

// GetLibroByID obtiene un libro por su ID, incluyendo la relación con Categoria
func (s *LibroService) GetLibroByID(id uint) (*models.LibroResponse, error) {
	var libro models.LibroResponse

	if err := s.db.Preload("Categoria").First(&libro, id).Error; err != nil {
		return nil, err
	}

	return &libro, nil
}

// CreateLibro crea un nuevo libro
func (s *LibroService) CreateLibro(libro *models.Libro) error {
	if err := models.ValidateLibro(libro); err != nil {
		return fmt.Errorf("datos del libro inválidos: %w", err)
	}

	if err := s.db.Create(libro).Error; err != nil {
		fmt.Printf("Error details: %v\n", err) // Log the error details
		return fmt.Errorf("error al crear libro xd: %w", err)
	}

	return nil
}

// UpdateLibro actualiza un libro
func (s *LibroService) UpdateLibro(libro *models.Libro) error {
	existingLibro := &models.Libro{}

	if err := s.db.First(existingLibro, libro.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("libro con ID %d no encontrado", libro.ID)
		}
		return fmt.Errorf("error al obtener libro con ID %d: %w", libro.ID, err)
	}

	if err := s.db.Model(existingLibro).Updates(libro).Error; err != nil {
		return fmt.Errorf("error al actualizar libro: %w", err)
	}

	if err := models.ValidateLibro(libro); err != nil {
		return fmt.Errorf("datos del libro inválidos: %w", err)
	}

	return nil

}
