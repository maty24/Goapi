package services

import (
	"errors"
	"fmt"
	"github.com/maty24/Goapi.git/internal/models"
	"gorm.io/gorm"
)

type CategoriaService struct {
	db *gorm.DB
}

func NewCategoriaService(db *gorm.DB) *CategoriaService {
	return &CategoriaService{db: db}
}

func (s *CategoriaService) GetAllCategorias() (*[]models.CategoriaResponse, error) {
	var categorias []models.CategoriaResponse

	if err := s.db.Find(&categorias).Error; err != nil {
		return nil, fmt.Errorf("error al obtener categorias: %w", err)
	}

	return &categorias, nil
}

func (s *CategoriaService) GetCategoriaByID(id uint) (*models.CategoriaResponse, error) {
	var categoria models.CategoriaResponse

	if err := s.db.First(&categoria, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("categoria con ID %d no encontrada", id)
		}
		return nil, fmt.Errorf("error al obtener categoria con ID %d: %w", id, err)
	}

	return &categoria, nil
}

func (s *CategoriaService) CreateCategoria(categoria *models.Categoria) error {
	if err := models.ValidateCategoria(categoria); err != nil {
		return fmt.Errorf("datos de la categoria inválidos: %w", err)
	}

	if err := s.db.Create(categoria).Error; err != nil {
		return fmt.Errorf("error al crear categoria: %w", err)
	}

	return nil
}

func (s *CategoriaService) UpdateCategoria(categoria *models.Categoria) error {
	existingCategoria := &models.Categoria{}

	if err := s.db.First(existingCategoria, categoria.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("categoria con ID %d no encontrada", categoria.ID)
		}
		return fmt.Errorf("error al obtener categoria con ID %d: %w", categoria.ID, err)
	}

	if err := models.ValidateCategoria(categoria); err != nil {
		return fmt.Errorf("datos de la categoria inválidos: %w", err)
	}

	if err := s.db.Model(existingCategoria).Updates(categoria).Error; err != nil {
		return fmt.Errorf("error al actualizar categoria: %w", err)
	}

	return nil
}
