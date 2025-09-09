package services

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/repositories"
	"abm_productos_categorias_go/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoriaInterface interface {
	CrearCategoria(categoria dto.CategoriaPeticion) (dto.CategoriaRespuesta, error)
}

type CategoriaService struct {
	repositories repositories.CategoriaRepositorioInterface
}

func NuevaCategoriaService(repositorio repositories.CategoriaRepositorioInterface) *CategoriaService {
	return &CategoriaService{repositories: repositorio}
}

func (service *CategoriaService) CrearCategoria(categoria dto.CategoriaPeticion) (dto.CategoriaRespuesta, error) {
	model := utils.ConvertirCategoriaDTOAModel(categoria)
	resultado, err := service.repositories.CrearCategoria(model)
	if err != nil {
		return dto.CategoriaRespuesta{}, err
	}
	oid, ok := resultado.InsertedID.(primitive.ObjectID)
	if ok {
		model.ID = oid
		return utils.ConvertirCategoriaModelADto(model), nil
	}
	return dto.CategoriaRespuesta{}, errors.New("Ocurrio un error al procesar el ID")
}
func (service *CategoriaService) ObtenerCategoriaPorId(id string) (dto.CategoriaRespuesta, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.CategoriaRespuesta{}, errors.New("Ocurrio un error con el ID ingresado :" + id)
	}
	model, err := service.repositories.ObtenerCategoriaPorId(oid)
	if err != nil {
		return dto.CategoriaRespuesta{}, errors.New("Categoria NO encontrada con el siguiente ID :" + id)
	}
	return utils.ConvertirCategoriaModelADto(model), nil
}
