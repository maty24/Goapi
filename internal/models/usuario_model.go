package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Usuario struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre             string    `gorm:"type:varchar(100);not null" json:"nombre" validate:"required,min=3,max=100"`
	Email              string    `gorm:"type:varchar(100);unique;not null" json:"email" validate:"required,email"`
	PasswordHash       string    `gorm:"type:varchar(255);not null" json:"password_hash" validate:"required,min=6"`
	TipoUsuario        string    `gorm:"type:varchar(50);not null;check:tipo_usuario IN ('lector','bibliotecario')" json:"tipo_usuario" validate:"required,oneof=lector bibliotecario"`
	Estado             string    `gorm:"type:varchar(50);default:activo;check:estado IN ('activo','inactivo')" json:"estado" validate:"required,oneof=activo inactivo"`
	FechaRegistro      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"fecha_registro"`
	UltimoInicioSesion time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"ultimo_inicio_sesion"`
}

// Claims define la estructura para los claims del JWT
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func (Usuario) TableName() string {
	return "usuarios"
}

// ValidateUsuario valida los datos de un usuario
func ValidateUsuario(usuario *Usuario) error {
	return validate.Struct(usuario)
}
