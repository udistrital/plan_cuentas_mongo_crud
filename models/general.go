package models

import "time"

// General estructura general de una entidad en plan de cuentas
type General struct {
	Vigencia          int       `json:"Vigencia" bson:"vigencia"`
	Nombre            string    `json:"Nombre" bson:"nombre"`
	Descripcion       string    `json:"Descripcion" bson:"descripcion"`
	FechaCreacion     time.Time `json:"FechaCreacion" bson:"fechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion" bson:"fechaModificacion", `
	Activo            bool      `json:"Activo" bson:"activo"`
}
