package utils

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertirDtoAModel(producto dto.ProductoPeticion) (model.Producto, error) {
	cat, err := primitive.ObjectIDFromHex(producto.CategoriaID)
	if err != nil {
		return model.Producto{}, err
	}
	return model.Producto{
		Descripcion: producto.Descripcion,
		Nombre:      producto.Nombre,
		CategoriaId: cat,
		Precio:      producto.Precio,
	}, err
}

func ConvertirModelADto(producto model.Producto) dto.ProductoRespuesta { //ESTO IRIA EN EL SERVICE, EL CARGAR LA CATEGORIA EN LA RESPUESTA
	return dto.ProductoRespuesta{
		Nombre:      producto.Nombre,
		Precio:      producto.Precio,
		Descripcion: producto.Descripcion,
		ID:          producto.ID.Hex(),
	}
}
