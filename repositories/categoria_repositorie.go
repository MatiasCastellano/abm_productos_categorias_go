package repositories

import (
	"abm_productos_categorias_go/database"
	"abm_productos_categorias_go/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoriaRepositorioInterface interface {
	CrearCategoria(categoria model.Categoria) (*mongo.InsertOneResult, error)
	EliminarCategoria(Id primitive.ObjectID) (*mongo.DeleteResult, error)
	ActualizarCategoria(categoria model.Categoria) (*mongo.UpdateResult, error)
	ObtenerCategorias(nombre string) ([]model.Categoria, error)
	ObtenerCategoriaPorId(id primitive.ObjectID) (model.Categoria, error)
}

type CategoriaRepositorio struct {
	db database.DB
}

func NuevoCategoriaRepositorio(db database.DB) *CategoriaRepositorio {
	return &CategoriaRepositorio{db: db}
}

func (repositorio CategoriaRepositorio) CrearCategoria(categoria model.Categoria) (*mongo.InsertOneResult, error) {
	colecion := repositorio.db.GetClient().Database("abm_productos").Collection("categorias")
	resultado, err := colecion.InsertOne(context.TODO(), categoria)
	return resultado, err
}

func (repositorio CategoriaRepositorio) EliminarCategoria(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	colecion := repositorio.db.GetClient().Database("abm_productos").Collection("categorias")
	filtro := bson.M{"_id": id}
	resultado, err := colecion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}

func (repositorio CategoriaRepositorio) ActualizarCategoria(categoria model.Categoria) (*mongo.UpdateResult, error) {
	colecion := repositorio.db.GetClient().Database("abm_productos").Collection("categorias")
	filtro := bson.M{"_id": categoria.ID}
	actualizacion := bson.M{
		"$set": bson.M{
			"nombre":      categoria.Nombre,
			"descripcion": categoria.Descripcion}}
	resultado, err := colecion.UpdateOne(context.TODO(), filtro, actualizacion)
	return resultado, err
}

func (repositorio CategoriaRepositorio) ObtenerCategorias(nombre string) ([]model.Categoria, error) {
	colecion := repositorio.db.GetClient().Database("abm_productos").Collection("categorias")
	filtro := bson.M{}
	if nombre != "" {
		filtro[nombre] = bson.M{"$regex": nombre, "$options": "i"}
	}
	cursor, err := colecion.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var categorias []model.Categoria
	for cursor.Next(context.Background()) {
		var categoria model.Categoria
		err := cursor.Decode(&categoria)
		if err != nil {
			continue
		}
		categorias = append(categorias, categoria)
	}
	return categorias, nil
}

func (repositorio ProductoRepositorio) ObtenerCategoriaPorId(id primitive.ObjectID) (model.Categoria, error) {
	colecion := repositorio.db.GetClient().Database("abm_productos").Collection("categorias")
	filtro := bson.M{"_id": id}
	var categoria model.Categoria
	err := colecion.FindOne(context.TODO(), filtro).Decode(&categoria)
	return categoria, err
}
