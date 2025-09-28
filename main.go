package main

import (
	"abm_productos_categorias_go/database"
	"abm_productos_categorias_go/handlers"
	"abm_productos_categorias_go/repositories"
	"abm_productos_categorias_go/services"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	productoHandler  *handlers.ProductosHandlres
	categoriaHandler *handlers.CategoriaHandler
	router           *gin.Engine
)

func main() {
	router = gin.Default()
	dependencies()
	mappingRoutes()
	log.Println("Iniciando en el servidor")
	router.Run(":8080")

}

func mappingRoutes() {
	productos := router.Group("/productos")
	{
		productos.POST("", productoHandler.CrearProducto)
		productos.GET("", productoHandler.ObtenerProductos)
		productos.GET("/:id", productoHandler.ObtenerProductoPorId)
		productos.PUT("/:id", productoHandler.ActualizarProducto)
		productos.DELETE("/:id", productoHandler.EliminarProducto)
	}

	categorias := router.Group("/categorias")
	{
		categorias.POST("", categoriaHandler.CrearCategoria)
		categorias.GET("/:id", categoriaHandler.ObtenerCategorias)
	}
}

func dependencies() {
	var db database.DB
	var ProductoRepositorio repositories.ProductoRepositorioInterface
	var ProductoService services.ProductoInterface
	var CategoriaRepositorio repositories.CategoriaRepositorioInterface
	var CategoriaService services.CategoriaInterface

	db = database.NewMongoDB()
	CategoriaRepositorio = repositories.NuevoCategoriaRepositorio(db)
	CategoriaService = services.NuevaCategoriaService(CategoriaRepositorio)
	categoriaHandler = handlers.NuevaCategoriaHandler(CategoriaService)

	ProductoRepositorio = repositories.NuevoProductoRepositorio(db)
	ProductoService = services.NuevoProductoService(ProductoRepositorio, CategoriaRepositorio)
	productoHandler = handlers.NuevoProductosHandlers(ProductoService)

}
