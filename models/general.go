package models

import "time"

// General estructura general de una entidad en plan de cuentas
type General struct {
	Vigencia          int       `json:"Vigencia" bson:"vigencia"`
	Nombre            string    `json:"Nombre" bson:"nombre"`
	Descripcion       string    `json:"Descripcion" bson:"descripcion"`
	FechaCreacion     time.Time `json:"FechaCreacion" bson:"fechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion" bson:"fechaModificacion"`
	Activo            bool      `json:"Activo" bson:"activo"`
}

// GeneralAfectationFileds .. main structure for balance.
type GeneralAfectationFileds struct {
	ValorActual  float64 `json:"ValorActual" bson:"ValorActual"`
	ValorInicial float64 `json:"ValorInicial" bson:"ValorInicial"`
}

// GeneralReducida estructura general reducida de una entidad en plan de cuentas
type GeneralReducida struct {
	Nombre string `json:"Nombre" bson:"nombre"`
}
