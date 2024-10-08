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
	dbHost, err := GetEnvStrict("DB_HOST")
	if err != nil {
		return Config{}, err
	}

	dbUser, err := GetEnvStrict("DB_USER")
	if err != nil {
		return Config{}, err
	}

	dbPassword, err := GetEnvStrict("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	dbName, err := GetEnvStrict("DB_NAME")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := GetEnvStrict("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		DBHost:     dbHost,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBPort:     dbPort,
	}, nil
}

// Validate valida que todas las variables de configuración estén presentes.
func (c *Config) Validate() error {
	if c.DBHost == "" || c.DBUser == "" || c.DBPassword == "" || c.DBName == "" || c.DBPort == "" {
		return errors.New("faltan variables de entorno necesarias para la configuración de la base de datos")
	}
	return nil
}
