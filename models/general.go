package models

// General estructura general de una entidad en plan de cuentas
type General struct {
	Vigencia    int                           `json:"Vigencia"`
	Nombre      string                        `json:"Nombre"`
	Descripcion string                        `json:"Descripcion"`
	IDPsql      int                           `json:"IdPsql"`
}
