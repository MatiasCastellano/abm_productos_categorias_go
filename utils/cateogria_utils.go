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
func ConvertirCategoriaDTOAModel(categoria dto.CategoriaPeticion) model.Categoria {
	return model.Categoria{
		Nombre:      categoria.Nombre,
		Descripcion: categoria.Descripcion,
	}
}
