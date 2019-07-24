package models

import "time"

// General estructura general de una entidad en plan de cuentas
type General struct {
	Vigencia          int       `json:"Vigencia"`
	Nombre            string    `json:"Nombre"`
	Descripcion       string    `json:"Descripcion"`
	FechaCreacion     time.Time `json:"FechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion"`
	Activo bool `json:"Activo"`
}
