package main

import (
	"context"
	"errors"
	"github.com/maty24/Goapi.git/pkg/db"
	"github.com/maty24/Goapi.git/pkg/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Cambiar al modo de producción
	gin.SetMode(gin.ReleaseMode)

	// Crear un contexto con cancelación para manejar el cierre de la aplicación
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Manejador de señales para cierre controlado
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Conectar a la base de datos
	database, err := db.InitDB(ctx)
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	// Verificar conexión
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión SQL: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error al cerrar la conexión SQL: %v", err)
		}
	}()

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	// Crear instancia de Gin
	r := gin.Default()

	// Configurar las rutas con la instancia de Gin y la conexión a la base de datos
	router.SetupRouter(r, database)

	// Configurar el puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001" // Puerto por defecto si no se especifica
	}

	// Iniciar el servidor en una goroutine para que podamos escuchar las señales del sistema
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Servidor iniciado en el puerto %s", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar señales de cierre
	<-stop
	log.Println("Apagando servidor...")

	// Crear un nuevo contexto con un tiempo de espera para el cierre del servidor
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	// Intentar un cierre ordenado del servidor
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error al apagar el servidor: %v", err)
	}

	log.Println("Servidor apagado correctamente")

}
