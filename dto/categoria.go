package dto

type CategoriaPeticion struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion,omitempty"`
}

type CategoriaRespuesta struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion,omitempty"`
}
