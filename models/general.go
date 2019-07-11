package models

// General estructura general de una entidad en plan de cuentas
type General struct {
	ID          string                        `json:"_id"`
	Vigencia    int                           `json:"Vigencia"`
	Nombre      string                        `json:"Nombre"`
	Descripcion string                        `json:"Descripcion"`
	IDPsql      int                           `json:"IdPsql"`
	Movimientos map[string]map[string]float64 `json:"Movimientos"`
}
