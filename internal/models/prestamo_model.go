package models

import (
	"time"
)

type Prestamo struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LibroID         uint      `gorm:"not null" json:"libro_id"`   // Relación con la tabla libros
	UsuarioID       uint      `gorm:"not null" json:"usuario_id"` // Relación con la tabla usuarios (lector)
	FechaPrestamo   time.Time `gorm:"type:date;default:CURRENT_DATE" json:"fecha_prestamo"`
	FechaDevolucion time.Time `gorm:"type:date" json:"fecha_devolucion" validate:"gtefield=FechaPrestamo"` // Validación para evitar devoluciones antes del préstamo
	FechaDevuelto   time.Time `gorm:"type:date" json:"fecha_devuelto"`
	Estado          string    `gorm:"type:varchar(50);default:pendiente;check:estado IN ('pendiente','devuelto')" json:"estado" validate:"required,oneof=pendiente devuelto"`
}

// TableName es una función personalizada para devolver el nombre de la tabla
func (Prestamo) TableName() string {
	return "prestamos"
}

type PrestamoResponse struct {
	ID              uint          `json:"id"`
	LibroID         uint          `json:"libro_id"`
	Libro           LibroResponse `json:"libro"`
	FechaPrestamo   time.Time     `json:"fecha_prestamo"`
	FechaDevolucion time.Time     `json:"fecha_devolucion"`
	FechaDevuelto   time.Time     `json:"fecha_devuelto"`
	Estado          string        `json:"estado"`
}

func (PrestamoResponse) TableName() string {
	return "prestamos"
}
