package models

// DocumentoPresupuestalCollection ... Nombre de la collección
const DocumentoPresupuestalCollection = "documento_presupuestal"

// DocumentoPresupuestal ... estructura para guardar información de documentos presupuestales.
type DocumentoPresupuestal struct {
	ID            string       `json:"_id" bson:"_id,omitempty"`
	Data          interface{}  `json:"Data" bson:"data" validate:"required"`
	Tipo          string       `json:"Tipo" bson:"tipo" validate:"required"`
	AfectacionIds []string     `json:"AfectacionIds" bson:"afectacion_ids"`
	Afectacion    []Movimiento `bson:"-" validate:"required"`
	FechaRegistro string       `json:"FechaRegistro" bson:"fecha_registro" validate:"required"`
	Estado        string       `json:"Estado" bson:"estado"`
	ValorActual   float64      `json:"ValorActual" bson:"valor_actual"`
	ValorInicial  float64      `json:"ValorInicial" bson:"valor_inicial"`
	Vigencia      int          `json:"Vigencia" bson:"vigencia" validate:"required"`
	CentroGestor  string       `json:"CentroGestor" bson:"centro_gestor" validate:"required"`
}
