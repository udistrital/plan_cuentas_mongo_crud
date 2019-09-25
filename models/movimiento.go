package models

// MovimientosCollection es el nombre de la colección en mongo.
const MovimientosCollection = "movimientos"
const MovimientoParameterCollection = "movimientos_parametros"

// Movimiento es una estructura generica para los tipos de movimiento registados.
type Movimiento struct {
	ID            string             `json:"_id" bson:"_id,omitempty"`
	IDPsql        int                `json:"IDPsql" bson:"IDPsql" validate:"required"`
	Valor         float64            `json:"Valor" bson:"Valor" validate:"required"`
	Tipo          string             `json:"Tipo" bson:"Tipo" validate:"required"`
	Padre         string             `json:"Padre" bson:"Padre"`
	FechaRegistro string             `json:"FechaRegistro" bson:"FechaRegistro" validate:"required"`
	Movimientos   map[string]float64 `json:"Movimientos" bson:"Movimientos"`
}

// MovimientoParameter this struct represent a "movmientos_parametros" collection item.
type MovimientoParameter struct {
	ID                  string `json:"_id" bson:"_id,omitempty"`
	TipoMovimientoHijo  string `json:"TipoMovimientoHijo" bson:"TipoMovimientoHijo"`
	TipoMovimientoPadre string `json:"TipoMovimientoPadre" bson:"TipoMovimientoPadre"`
	Multiplicador       int    `json:"Multiplicador" bson:"Multiplicador"`
}
