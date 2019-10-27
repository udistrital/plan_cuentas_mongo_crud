package models

// VigenciaCollectionName nombre de la colleccion para guardar las agrupaciones de vigencias.
const VigenciaCollectionName = "vigencia"

// Vigencia estructura para acceder de forma mas rápida a la información de las vigencias registradas.
type Vigencia struct {
	ID           string `json:"Id" bson:"_id,omitempty"`
	NameSapce    string `json:"NameSpace" bson:"name_space"`
	CentroGestor string `json:"CentroGestor" bson:"centro_gestor"`
	Valor        int    `json:"Valor" bson:"valor"`
	Estado       string `json:"Estado" bson:"estado"`
}
