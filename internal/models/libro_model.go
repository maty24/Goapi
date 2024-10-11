package models

import (
	"github.com/maty24/Goapi.git/pkg/globals"
	"time"
)

type Libro struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Titulo           string    `gorm:"type:varchar(150);not null;uniqueIndex:idx_titulo_autor" json:"titulo" validate:"required,min=3,max=150"`
	AutorID          uint      `gorm:"not null;index:idx_titulo_autor" json:"autor_id" validate:"required"` // Relación con autores
	CategoriaID      uint      `gorm:"index" json:"categoria_id"`                                           // Relación con categorías
	FechaPublicacion time.Time `gorm:"type:date;check:fecha_publicacion <= current_date" json:"fecha_publicacion"`
	Disponible       bool      `gorm:"default:true" json:"disponible"`
}

// TableName es una función personalizada para devolver el nombre de la tabla
func (Libro) TableName() string {
	return "libros"
}

// LibroResponse incluye la información del libro, el autor y la categoría
type LibroResponse struct {
	ID               uint              `json:"id"`
	Titulo           string            `json:"titulo" validate:"required,min=3,max=150"`
	Autor            AutorResponse     `json:"autor"`     // Devolver los datos del autor en lugar de solo el ID
	Categoria        CategoriaResponse `json:"categoria"` // Devolver los datos de la categoría en lugar de solo el ID
	FechaPublicacion time.Time         `json:"fecha_publicacion"`
	Disponible       bool              `json:"disponible"`
}

// TableName es una función personalizada para devolver el nombre de la tabla (opcional)
func (LibroResponse) TableName() string {
	return "libros"
}

// ValidateLibro valida los datos de un libro
func ValidateLibro(libro *Libro) error {
	return globals.Validate.Struct(libro)
}
