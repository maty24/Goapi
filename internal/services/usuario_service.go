package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maty24/Goapi.git/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

var jwtSecretKey = []byte("my_secret_key") // Debe ser almacenada de manera segura

type UsuarioService struct {
	db *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{db: db}
}

func (s *UsuarioService) GetAllUsuarios() (*[]models.Usuario, error) {
	var usuarios []models.Usuario

	if err := s.db.Find(&usuarios).Error; err != nil {
		return nil, err
	}

	return &usuarios, nil
}

func (s *UsuarioService) GetUsuarioByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario

	if err := s.db.First(&usuario, id).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (s *UsuarioService) CreateUsuario(usuario *models.Usuario) error {
	// Validar los datos del usuario
	if err := models.ValidateUsuario(usuario); err != nil {
		return err
	}

	// Hashear la contraseña
	hashedPassword, err := hashPassword(usuario.PasswordHash)
	if err != nil {
		return err
	}

	// Asignar el hash a la estructura del usuario
	usuario.PasswordHash = hashedPassword

	// Crear el usuario en la base de datos
	if err := s.db.Create(usuario).Error; err != nil {
		return err
	}

	return nil
}

// Login Función de login que devuelve un token JWT
func (s *UsuarioService) Login(email, password string) (string, error) {
	var usuario models.Usuario

	// Buscar el usuario por email
	if err := s.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return "", errors.New("usuario no encontrado")
	}

	// Verificar si la contraseña es correcta
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	// Crear el token JWT
	token, err := generateJWT(usuario.ID, usuario.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Función para generar un token JWT
func generateJWT(userID uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // El token expira en 24 horas

	claims := &models.Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Fecha de expiración del token
		},
	}

	// Crear el token con los claims definidos
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Función para hashear la contraseña
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
