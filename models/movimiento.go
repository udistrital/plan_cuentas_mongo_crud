package models

// MovimientosCollection es el nombre de la colección en mongo.
const MovimientosCollection = "movimientos"

// MovimientoParameterCollection nombre de la coleccion para guardar los parametros de los movmientos.
const MovimientoParameterCollection = "movimientos_parametros"

// Movimiento es una estructura generica para los tipos de movimiento registados.
type Movimiento struct {
	ID                        string             `json:"_id" bson:"_id,omitempty"`
	IDPsql                    int                `json:"IDPsql" bson:"IDPsql" validate:"required"`
	Tipo                      string             `json:"Tipo" bson:"Tipo" validate:"required"`
	DocumentoPresupuestalUUID string             `json:"DocumentoPresupuestalUUID" bson:"DocumentoPresupuestalUUID"`
	Padre                     string             `json:"Padre" bson:"Padre"`
	FechaRegistro             string             `json:"FechaRegistro" bson:"FechaRegistro" validate:"required"`
	Movimientos               map[string]float64 `json:"Movimientos" bson:"Movimientos"`
	Estado                    string             `json:"Estado" bson:"Estado"`
	ValorActual               float64            `json:"ValorActual" bson:"ValorActual"`
	ValorInicial              float64            `json:"ValorInicial" bson:"ValorInicial"`
}

// MovimientoParameter this struct represent a "movmientos_parametros" collection item.
type MovimientoParameter struct {
	ID                   string `json:"_id" bson:"_id,omitempty"`
	TipoMovimientoHijo   string `json:"TipoMovimientoHijo" bson:"TipoMovimientoHijo"`
	TipoMovimientoPadre  string `json:"TipoMovimientoPadre" bson:"TipoMovimientoPadre"`
	Multiplicador        int    `json:"Multiplicador" bson:"Multiplicador"`
	FatherCollectionName string `json:"FatherCollectionName" bson:"FatherCollectionName"`
	FatherUUIKeyName     string `json:"FatherUUIKeyName" bson:"FatherUUIKeyName"`
}
