package models

import (
	"time"
)

type Libro struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Titulo           string    `gorm:"type:varchar(150);not null" json:"titulo" validate:"required,min=3,max=150"`
	AutorID          uint      `gorm:"not null" json:"autor_id"`
	CategoriaID      uint      `json:"categoria_id"`
	FechaPublicacion time.Time `gorm:"type:date" json:"fecha_publicacion"`
	Disponible       bool      `gorm:"default:true" json:"disponible"`
}

func (Libro) TableName() string {
	return "libros"
}
