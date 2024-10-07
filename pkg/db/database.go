package db

import (
	"context"
	"fmt"
	"github.com/maty24/Goapi.git/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"runtime"
	"time"
)

// InitDB inicializa la conexión a la base de datos.
func InitDB(ctx context.Context) (*gorm.DB, error) {
	config := config.LoadDBConfig()

	// Validar la configuración
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Construir el DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Santiago",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort,
	)

	// Conectar a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	// Configurar el pool de conexiones
	if err := configureConnectionPool(db); err != nil {
		return nil, err
	}

	// Verificar conexión inicial
	if err := verifyConnection(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

// configureConnectionPool configura el pool de conexiones SQL.
func configureConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error al obtener la conexión SQL: %v", err)
	}

	numCPU := runtime.NumCPU()
	sqlDB.SetMaxIdleConns(numCPU)
	sqlDB.SetMaxOpenConns(numCPU * 2)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * 10)

	return nil
}

// verifyConnection verifica que la conexión a la base de datos sea exitosa..

func verifyConnection(ctx context.Context, db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error al obtener la conexión SQL: %v", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos exitosa")
	return nil
}
