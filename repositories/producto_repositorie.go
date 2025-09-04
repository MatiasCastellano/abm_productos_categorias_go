package repositories

import (
	"abm_productos_categorias_go/database"
	"abm_productos_categorias_go/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductoRepositorioInterface interface {
	CrearProducto(producto model.Producto) (*mongo.InsertOneResult, error)
	EliminarProducto(Id primitive.ObjectID) (*mongo.DeleteResult, error)
	ActualizarProducto(producto model.Producto) (*mongo.UpdateResult, error)
	ObtenerProductos(nombre string, precio float64) ([]model.Producto, error)
	ObtenerProductoPorId(id primitive.ObjectID) (model.Producto, error)
}

type ProductoRepositorio struct {
	db database.DB
}

func NuevoProductoRepositorio(db database.DB) *ProductoRepositorio {
	return &ProductoRepositorio{db: db}
}

func (repositorio ProductoRepositorio) CrearProducto(producto model.Producto) (*mongo.InsertOneResult, error) {
	colleccion := repositorio.db.GetClient().Database("abm_productos").Collection("productos")
	resultado, err := colleccion.InsertOne(context.TODO(), producto)
	return resultado, err
}

func (repositorio ProductoRepositorio) EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	colleccion := repositorio.db.GetClient().Database("abm_productos").Collection("productos")
	filtro := bson.M{"_id": id}
	resultado, err := colleccion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}

func (repositorio ProductoRepositorio) ActualizarProducto(producto model.Producto) (*mongo.UpdateResult, error) {
	coleccion := repositorio.db.GetClient().Database("abm_productos").Collection("productos")
	filtro := bson.M{"_id": producto.ID}
	actualizacion := bson.M{
		"$set": bson.M{
			"categoria_id": producto.CategoriaId,
			"nombre":       producto.Nombre,
			"precio":       producto.Precio,
			"descripcion":  producto.Descripcion}}
	resultado, err := coleccion.UpdateOne(context.TODO(), filtro, actualizacion)
	return resultado, err
}

func (repositorio ProductoRepositorio) ObtenerProductos(nombre string, precio float64) ([]model.Producto, error) {
	coleccion := repositorio.db.GetClient().Database("abm_productos").Collection("productos")
	filtro := bson.M{}
	if nombre != "" {
		filtro[nombre] = bson.M{"$regex": nombre, "$options": "i"}
	}
	if precio >= 0 {
		filtro["precio"] = bson.M{"$gt": precio}
	}
	cursor, err := coleccion.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var productos []model.Producto
	for cursor.Next(context.Background()) {
		var producto model.Producto
		err := cursor.Decode(&producto)
		if err != nil {
			continue
		}
		productos = append(productos, producto)
	}
	return productos, nil
}

func (repositorio ProductoRepositorio) ObtenerProductoPorId(id primitive.ObjectID) (model.Producto, error) {
	coleccion := repositorio.db.GetClient().Database("abm_productos").Collection("productos")
	filtro := bson.M{"_id": id}
	var producto model.Producto
	err := coleccion.FindOne(context.TODO(), filtro).Decode(&producto)
	return producto, err
}
