package models

import (
	"time"
)

type Prestamo struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LibroID         uint      `gorm:"not null" json:"libro_id"`
	LectorID        uint      `gorm:"not null" json:"lector_id"`
	FechaPrestamo   time.Time `gorm:"type:date;default:CURRENT_DATE" json:"fecha_prestamo"`
	FechaDevolucion time.Time `gorm:"type:date" json:"fecha_devolucion"`
	FechaDevuelto   time.Time `gorm:"type:date" json:"fecha_devuelto"`
	Estado          string    `gorm:"type:varchar(50);default:pendiente;check:estado IN ('pendiente','devuelto')" json:"estado" validate:"required,oneof=pendiente devuelto"`
}

func (Prestamo) TableName() string {
	return "prestamos"
}
