package models

import (
	"github.com/maty24/Goapi.git/pkg/globals"
)

type Autor struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"type:varchar(100);not null" json:"nombre" validate:"required,min=3,max=100"`
}

type AutorResponse struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
}

// TableName Custom function to return table name
func (Autor) TableName() string {
	return "autores"
}

func (AutorResponse) TableName() string {
	return "autores"
}

// ValidateAutor Function to validate Autor
func ValidateAutor(autor *Autor) error {
	return globals.Validate.Struct(autor)
}
