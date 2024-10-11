package db

import (
	"context"
	"fmt"
	"github.com/maty24/Goapi.git/pkg/config"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(ctx context.Context) (*gorm.DB, error) {
	config, err := config.LoadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error cargando la configuración de la base de datos: %v", err)
	}

	// Configurar valores opcionales (sslmode y TimeZone)
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "America/Santiago" // Valor por defecto si no se especifica
	}

	// Construir el DSN de forma más segura
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DBHost, url.QueryEscape(config.DBUser), url.QueryEscape(config.DBPassword),
		config.DBName, config.DBPort, sslMode, timeZone,
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

	// Verificar la conexión inicial
	if err := verifyConnection(ctx, db); err != nil {
		return nil, err
	}

	log.Println("Conexión a la base de datos exitosa")
	return db, nil
}

// configureConnectionPool configura el pool de conexiones SQL.
func configureConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error al obtener la conexión SQL: %v", err)
	}

	// Cargar valores de entorno opcionales para el pool de conexiones
	maxIdleConns := getEnvInt("DB_MAX_IDLE_CONNS", runtime.NumCPU())
	maxOpenConns := getEnvInt("DB_MAX_OPEN_CONNS", runtime.NumCPU()*2)
	connMaxLifetime := getEnvDuration("DB_CONN_MAX_LIFETIME", time.Hour)
	connMaxIdleTime := getEnvDuration("DB_CONN_MAX_IDLE_TIME", time.Minute*10)

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)

	return nil
}

// verifyConnection verifica que la conexión a la base de datos sea exitosa.
func verifyConnection(ctx context.Context, db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error al obtener la conexión SQL: %v", err)
	}

	// Verificar con PingContext (verifica el tiempo de respuesta de la DB)
	if err := sqlDB.PingContext(ctx); err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %v", err)
	}

	return nil
}

// Funciones auxiliares para leer valores de entorno
func getEnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if exists {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	value, exists := os.LookupEnv(key)
	if exists {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}
