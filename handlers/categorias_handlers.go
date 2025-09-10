package handlers

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriaHandler struct {
	service services.CategoriaInterface
}

func NuevaCategoriaHandler(ser services.CategoriaInterface) *CategoriaHandler {
	return &CategoriaHandler{
		service: ser,
	}
}

func (handler *CategoriaHandler) CrearCategoria(c *gin.Context) {
	var solicitud dto.CategoriaPeticion
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	categoria, error := handler.service.CrearCategoria(solicitud)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusCreated, categoria)
}

func (handler *CategoriaHandler) ObtenerCategorias(c *gin.Context) {
	var solicitud dto.FiltroProducto
	solicitud.Nombre = c.Query("nombre")
	categorias, err := handler.service.ObtenerCategorias(solicitud)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categorias)
}
