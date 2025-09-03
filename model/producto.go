package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Producto struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Nombre      string             `bson:"nombre"`
	Descripcion string             `bson:"descripcion"`
	Precio      float64            `bson:"precio"`
	CategoriaId primitive.ObjectID `bson:"categoria_id"`
}
