package models

type Categoria struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nombre string `gorm:"type:varchar(100);not null" json:"nombre" validate:"required,min=3,max=100"`
}

func (Categoria) TableName() string {
	return "categorias"
}
