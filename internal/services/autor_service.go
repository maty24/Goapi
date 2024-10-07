package services

import (
	"github.com/maty24/Goapi.git/internal/models"
	"gorm.io/gorm"
)

type AutorService struct {
	DB *gorm.DB
}

func newAutorService(db *gorm.DB) *AutorService {
	return &AutorService{DB: db}
}

// GetAllAutores obtiene todos los autores
func (s *AutorService) GetAllAutores() (*[]models.Autor, error) {
	var autores []models.Autor

	if err := s.DB.Find(&autores).Error; err != nil {
		return nil, err
	}

	return &autores, nil
}
