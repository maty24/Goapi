package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maty24/Goapi.git/internal/controller"
	"github.com/maty24/Goapi.git/internal/services"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {

	//Autor
	autorService := services.NewAutorService(db)
	autorController := controller.NewAutorController(autorService)

	//Categoria
	categoriaService := services.NewCategoriaService(db)
	categoriaController := controller.NewCategoriaController(categoriaService)

	//Libro
	libroService := services.NewLibroService(db)
	libroController := controller.NewLibroController(libroService)

	//Usuario
	usuarioService := services.NewUsuarioService(db)
	usuarioController := controller.NewUsuarioController(usuarioService)

	//Prestamos
	prestamoService := services.NewPrestamoService(db)
	prestamoController := controller.NewPrestamoController(prestamoService)

	api := r.Group("/api/v1")
	{
		autor := api.Group("/autor")
		{
			autor.GET("/", autorController.GetAllAutores)
			autor.GET("/:id", autorController.GetAutorByID)
			autor.POST("/", autorController.CreateAutor)
			autor.PATCH("/:id", autorController.UpdateAutor)
		}

		categoria := api.Group("/categoria")
		{
			categoria.GET("/", categoriaController.GetAllCategorias)
			categoria.GET("/:id", categoriaController.GetCategoriaByID)
			categoria.POST("/", categoriaController.CreateCategoria)
			categoria.PATCH("/:id", categoriaController.UpdateCategoria)
		}

		libro := api.Group("/libro")
		{
			libro.GET("/", libroController.GetAllLibros)
			libro.GET("/:id", libroController.GetLibroByID)
			libro.POST("/", libroController.CreateLibro)
			libro.PATCH("/:id", libroController.UpdateLibro)
		}

		usuario := api.Group("/usuario")
		{
			usuario.GET("/", usuarioController.GetAllUsuarios)
			usuario.GET("/:id", usuarioController.GetUsuarioByID)
			usuario.POST("/", usuarioController.CreateUsuario)
			usuario.POST("/login", usuarioController.LoginUsuario)
		}

		prestamo := api.Group("/prestamo")
		{
			prestamo.GET("/:id", prestamoController.GetActivePrestamosByLectorID)
			prestamo.POST("/", prestamoController.CreatePrestamo)
			prestamo.PATCH("/:id", prestamoController.ReturnPrestamo)
		}
	}

}
