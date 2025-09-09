package dto

type ProductoPeticion struct {
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion,omitempty"`
	Precio      float64 `json:"precio"`
	CategoriaID string  `json:"categoriaID"`
}

type ProductoRespuesta struct {
	ID          string             `json:"id"`
	Nombre      string             `json:"nombre"`
	Descripcion string             `json:"descripcion,omitempty"`
	Precio      float64            `json:"precio"`
	Categoria   CategoriaRespuesta `json:"categoria"`
}

type FiltroProducto struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
