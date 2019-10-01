package models

// DocumentoPresupuestalCollection ... Nombre de la collección
const DocumentoPresupuestalCollection = "documento_presupuestal"

// DocumentoPresupuestal ... estructura para guardar información de documentos presupuestales.
type DocumentoPresupuestal struct {
	ID            string       `json:"_id" bson:"_id,omitempty"`
	Data          interface{}  `json:"Data" bson:"Data" validate:"required"`
	Tipo          string       `json:"Tipo" bson:"Tipo" validate:"required"`
	AfectacionIds []string     `json:"AfectacionIds" bson:"AfectacionIds"`
	Afectacion    []Movimiento `bson:"-" validate:"required"`
	FechaRegistro string       `json:"FechaRegistro" bson:"FechaRegistro" validate:"required"`
	Estado        string       `json:"Estado" bson:"Estado"`
	ValorActual   float64      `json:"ValorActual" bson:"ValorActual"`
	ValorInicial  float64      `json:"ValorInicial" bson:"ValorInicial"`
}
