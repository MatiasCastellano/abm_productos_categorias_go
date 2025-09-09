package handlers

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductosHandlres struct {
	service services.ProductoInterface
}

func NuevoProductosHandlers(ser services.ProductoInterface) *ProductosHandlres {
	return &ProductosHandlres{service: ser}
}

func (Handler *ProductosHandlres) CrearProducto(c *gin.Context) {
	var solicitud dto.ProductoPeticion
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	producto, err := Handler.service.CrearProducto(solicitud)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}
	c.JSON(http.StatusCreated, producto)
}
