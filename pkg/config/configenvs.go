package config

import (
	"errors"
	"fmt"
	_ "log"
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

// GetEnvStrict obtiene el valor de la variable de entorno si existe, de lo contrario, retorna un error.
func GetEnvStrict(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("la variable de entorno %s es obligatoria pero no está configurada", key)
	}
	return value, nil
}

// LoadDBConfig carga la configuración de la base de datos desde las variables de entorno.
func LoadDBConfig() (Config, error) {
	requiredEnvVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	var missingVars []string

	for _, envVar := range requiredEnvVars {
		if _, exists := os.LookupEnv(envVar); !exists {
			missingVars = append(missingVars, envVar)
		}
	}

	if len(missingVars) > 0 {
		return Config{}, fmt.Errorf("las siguientes variables de entorno son obligatorias: %v", missingVars)
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}, nil
}

// Validate valida que todas las variables de configuración estén presentes.
func (c *Config) Validate() error {
	if c.DBHost == "" || c.DBUser == "" || c.DBPassword == "" || c.DBName == "" || c.DBPort == "" {
		return errors.New("faltan variables de entorno necesarias para la configuración de la base de datos")
	}
	return nil
}
