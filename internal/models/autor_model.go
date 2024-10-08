package models

import (
	"github.com/go-playground/validator/v10"
)

type Autor struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"type:varchar(100);not null" json:"nombre" validate:"required,min=3,max=100"`
}

// TableName Custom function to return table name
func (Autor) TableName() string {
	return "autores"
}

var validate = validator.New()

// ValidateAutor Function to validate Autor
func ValidateAutor(autor *Autor) error {
	return validate.Struct(autor)
}
