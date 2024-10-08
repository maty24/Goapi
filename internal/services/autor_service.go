package services

import (
	"errors"
	"fmt"
	"github.com/maty24/Goapi.git/internal/models"
	"gorm.io/gorm"
)

type AutorService struct {
	DB *gorm.DB
}

func NewAutorService(db *gorm.DB) *AutorService {
	return &AutorService{DB: db}
}

// GetAllAutores obtiene todos los autores
func (s *AutorService) GetAllAutores() ([]models.Autor, error) {
	var autores []models.Autor

	if err := s.DB.Find(&autores).Error; err != nil {
		// Aquí podrías agregar un log si lo consideras necesario
		return nil, fmt.Errorf("error al obtener autores: %w", err)
	}

	return autores, nil
}

// GetAutorByID obtiene un autor por su ID
func (s *AutorService) GetAutorByID(id uint) (*models.Autor, error) {
	var autor models.Autor

	if err := s.DB.First(&autor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("autor con ID %d no encontrado", id)
		}
		return nil, fmt.Errorf("error al obtener autor con ID %d: %w", id, err)
	}

	return &autor, nil
}

// CreateAutor crea un nuevo autor
func (s *AutorService) CreateAutor(autor *models.Autor) error {
	// Validar el autor
	if err := models.ValidateAutor(autor); err != nil {
		return fmt.Errorf("datos del autor inválidos: %w", err)
	}

	if err := s.DB.Create(autor).Error; err != nil {
		return fmt.Errorf("error al crear autor: %w", err)
	}

	return nil
}

// UpdateAutor actualiza un autor
func (s *AutorService) UpdateAutor(autor *models.Autor) error {
	existingAutor := &models.Autor{}
	if err := s.DB.First(existingAutor, autor.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("autor con ID %d no encontrado", autor.ID)
		}
		return fmt.Errorf("error al buscar autor con ID %d: %w", autor.ID, err)
	}

	// Validar el autor
	if err := models.ValidateAutor(autor); err != nil {
		return fmt.Errorf("datos del autor inválidos: %w", err)
	}

	if err := s.DB.Model(existingAutor).Updates(autor).Error; err != nil {
		return fmt.Errorf("error al actualizar autor con ID %d: %w", autor.ID, err)
	}

	return nil
}
