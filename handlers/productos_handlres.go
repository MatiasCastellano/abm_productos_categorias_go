package handlers

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/services"
	"net/http"
	"strconv"

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

func (handlres *ProductosHandlres) ObtenerProductos(c *gin.Context) {
	var solicitud dto.FiltroProducto
	solicitud.Nombre = c.Query("nombre")
	PreciosStr := c.Query("precio")
	if PreciosStr != "" {
		precio, err := strconv.ParseFloat(PreciosStr, 64)
		if err == nil {
			solicitud.Precio = precio
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
	}
	productos, err := handlres.service.ObtenerProductos(solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

func (handlers *ProductosHandlres) ObtenerProductoPorId(c *gin.Context) {
	id := c.Param("id")
	producto, err := handlers.service.ObtenerProductoPorID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (handler *ProductosHandlres) ActualizarProducto(c *gin.Context) {
	var solicitud dto.ProductoPeticion
	id := c.Param("id")
	if err := c.ShouldBindJSON(&solicitud); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	Producto, err := handler.service.ActualizarProducto(id, solicitud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Producto)
}

func (handler *ProductosHandlres) EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	err := handler.service.EliminarProducto(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Resultado": "Se elimino correctamente"})
}
