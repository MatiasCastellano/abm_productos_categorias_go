package services

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/repositories"
	"abm_productos_categorias_go/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductoInterface interface {
}

type ProductoService struct {
	repositorioProducto  repositories.ProductoRepositorioInterface
	repositorioCategoria repositories.CategoriaRepositorioInterface
}

func NuevoProductoService(repositorioProducto repositories.ProductoRepositorioInterface, repositorioCategoria repositories.CategoriaRepositorioInterface) *ProductoService {
	return &ProductoService{
		repositorioProducto:  repositorioProducto,
		repositorioCategoria: repositorioCategoria,
	}
}

func (service *ProductoService) CrearProducto(producto dto.ProductoPeticion) (dto.ProductoRespuesta, error) {
	model, err := utils.ConvertirDtoAModel(producto)
	if err != nil {
		return dto.ProductoRespuesta{}, err
	}
	resultado, err := service.repositorioProducto.CrearProducto(model)
	if err != nil {
		return dto.ProductoRespuesta{}, err
	}
	modelCategoria, error := service.repositorioCategoria.ObtenerCategoriaPorId(model.CategoriaId)
	if error != nil {
		return dto.ProductoRespuesta{}, error
	}
	categoriaRespuesta := utils.ConvertirCategoriaModelADto(modelCategoria)
	oid, ok := resultado.InsertedID.(primitive.ObjectID)
	if ok {
		model.ID = oid
		productoRespuesta := utils.ConvertirModelADto(model)
		productoRespuesta.Categoria = categoriaRespuesta
		return productoRespuesta, nil
	}
	return dto.ProductoRespuesta{}, errors.New("ocurrio un error al procesar el ID")
}
