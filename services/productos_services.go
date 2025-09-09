package services

import (
	"abm_productos_categorias_go/dto"
	"abm_productos_categorias_go/repositories"
	"abm_productos_categorias_go/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductoInterface interface {
	CrearProducto(producto dto.ProductoPeticion) (dto.ProductoRespuesta, error)
	EliminarProducto(id string) error
	ObtenerProductoPorID(id string) (dto.ProductoRespuesta, error)
	ActualizarProducto(id string, producto dto.ProductoPeticion) (dto.ProductoRespuesta, error)
	ObtenerProductos(filtros dto.FiltroProducto) ([]dto.ProductoRespuesta, error)
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

func (service *ProductoService) EliminarProducto(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("el id ingresado no es formato correcto")
	}
	_, err = service.repositorioProducto.EliminarProducto(oid)
	return err
}
func (service *ProductoService) ObtenerProductoPorID(id string) (dto.ProductoRespuesta, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductoRespuesta{}, errors.New("el id ingresado no es formato correcto")
	}
	model, errorRepoProducto := service.repositorioProducto.ObtenerProductoPorId(oid)
	if errorRepoProducto != nil {
		return dto.ProductoRespuesta{}, errors.New("Producto NO encontrado")
	}
	modelCat, errorRepoCategoria := service.repositorioCategoria.ObtenerCategoriaPorId(model.CategoriaId)
	if errorRepoCategoria != nil {
		return dto.ProductoRespuesta{}, errors.New("el producto no tiene categoria")
	}
	categoriaRespuesta := utils.ConvertirCategoriaModelADto(modelCat)
	productoRespuesta := utils.ConvertirModelADto(model)
	productoRespuesta.Categoria = categoriaRespuesta
	return productoRespuesta, nil
}

func (service *ProductoService) ActualizarProducto(id string, producto dto.ProductoPeticion) (dto.ProductoRespuesta, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductoRespuesta{}, errors.New("Ocurrio un error con el ID ingresado :" + id)
	}
	model, err := utils.ConvertirDtoAModel(producto)
	if err != nil {
		return dto.ProductoRespuesta{}, err
	}
	model.ID = oid
	_, err = service.repositorioProducto.ActualizarProducto(model)
	if err != nil {
		return dto.ProductoRespuesta{}, errors.New("No se pudo actualizar el producto")
	}
	productoRespuesta := utils.ConvertirModelADto(model)
	categoriaModel, err := service.repositorioCategoria.ObtenerCategoriaPorId(model.CategoriaId)
	if err != nil {
		return dto.ProductoRespuesta{}, errors.New("no se pudo encontrar la categoria perteneciente al producto")
	}
	productoRespuesta.Categoria = utils.ConvertirCategoriaModelADto(categoriaModel)
	return productoRespuesta, nil
}

func (service *ProductoService) ObtenerProductos(filtros dto.FiltroProducto) ([]dto.ProductoRespuesta, error) {
	productos, err := service.repositorioProducto.ObtenerProductos(filtros)
	if err != nil {
		return []dto.ProductoRespuesta{}, errors.New("Ocurrio un error al obtener los productos")
	}
	var resultado []dto.ProductoRespuesta
	for _, producto := range productos {
		categoriaModel, err := service.repositorioCategoria.ObtenerCategoriaPorId(producto.CategoriaId)
		if err != nil {
			return []dto.ProductoRespuesta{}, nil
		}
		productoRespuesta := utils.ConvertirModelADto(producto)
		productoRespuesta.Categoria = utils.ConvertirCategoriaModelADto(categoriaModel)
		resultado = append(resultado, productoRespuesta)
	}
	return resultado, nil
}
