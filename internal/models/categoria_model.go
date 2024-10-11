package models

import "github.com/maty24/Goapi.git/pkg/globals"

type Categoria struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"type:varchar(100);not null;unique" json:"nombre" validate:"required,min=3,max=100"` // Añadido unique
}

type CategoriaResponse struct {
	ID     uint   `json:"id"`
	Nombre string `json:"nombre"`
}

func (CategoriaResponse) TableName() string {
	return "categorias"
}

// TableName es una función personalizada para devolver el nombre de la tabla
func (Categoria) TableName() string {
	return "categorias"
}

func ValidateCategoria(categoria *Categoria) error {
	return globals.Validate.Struct(categoria)
}
