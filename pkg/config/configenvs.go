package config

import (
	"errors"
	"log"
	"os"
)

// Config estructura de configuración para la base de datos.
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

// GetEnvStrict GetEnv retorna el valor de la variable de entorno si existe, de lo contrario, retorna el valor por defecto.
func GetEnvStrict(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("La variable de entorno %s es obligatoria pero no está configurada", key)
	}
	return value
}

// LoadDBConfig carga la configuración de la base de datos desde las variables de entorno.
func LoadDBConfig() Config {
	return Config{
		DBHost:     GetEnvStrict("DB_HOST"),
		DBUser:     GetEnvStrict("DB_USER"),
		DBPassword: GetEnvStrict("DB_PASSWORD"),
		DBName:     GetEnvStrict("DB_NAME"),
		DBPort:     GetEnvStrict("DB_PORT"),
	}
}

// Validate valida que todas las variables de configuración estén presentes.
func (c *Config) Validate() error {
	if c.DBHost == "" || c.DBUser == "" || c.DBPassword == "" || c.DBName == "" || c.DBPort == "" {
		return errors.New("faltan variables de entorno necesarias para la configuración de la base de datos")
	}
	return nil
}
