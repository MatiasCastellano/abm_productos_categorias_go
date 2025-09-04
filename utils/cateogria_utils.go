package utils

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/model"
)

func ConvertirCategoriaModelADto(categoria model.Categoria) dto.CategoriaRespuesta {
	return dto.CategoriaRespuesta{
		Nombre:      categoria.Nombre,
		Descripcion: categoria.Descripcion,
		ID:          categoria.ID.Hex(),
	}
}
