package models

import "time"

// General estructura general de una entidad en plan de cuentas
type General struct {
	Vigencia          int       `json:"Vigencia" bson:"vigencia"`
	Nombre            string    `json:"Nombre" bson:"nombre"`
	Descripcion       string    `json:"-" bson:"descripcion"`
	FechaCreacion     time.Time `json:"-" bson:"fechaCreacion"`
	FechaModificacion time.Time `json:"-" bson:"fechaModificacion"`
	Activo            bool      `json:"-" bson:"activo"`
}

// GeneralAfectationFileds .. main structure for balance.
type GeneralAfectationFileds struct {
	ValorActual  float64 `json:"ValorActual" bson:"ValorActual"`
	ValorInicial float64 `json:"ValorInicial" bson:"ValorInicial"`
}
