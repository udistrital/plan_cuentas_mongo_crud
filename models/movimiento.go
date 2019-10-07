package models

// MovimientosCollection es el nombre de la colección en mongo.
const MovimientosCollection = "movimientos"

// MovimientoParameterCollection nombre de la coleccion para guardar los parametros de los movmientos.
const MovimientoParameterCollection = "movimientos_parametros"

// Movimiento es una estructura generica para los tipos de movimiento registados.
type Movimiento struct {
	ID                        string             `json:"_id" bson:"_id,omitempty"`
	IDPsql                    int                `json:"IDPsql" bson:"id_psql" validate:"required"`
	Tipo                      string             `json:"Tipo" bson:"tipo" validate:"required"`
	DocumentoPresupuestalUUID string             `json:"DocumentoPresupuestalUUID" bson:"documento_presupuestal_uuid"`
	Padre                     string             `json:"Padre" bson:"padre"`
	FechaRegistro             string             `json:"FechaRegistro" bson:"fecha_registro" validate:"required"`
	Movimientos               map[string]float64 `json:"Movimientos" bson:"movimientos"`
	Estado                    string             `json:"Estado" bson:"estado"`
	ValorActual               float64            `json:"ValorActual" bson:"valor_actual"`
	ValorInicial              float64            `json:"ValorInicial" bson:"valor_inicial"`
}

// MovimientoParameter this struct represent a "movmientos_parametros" collection item.
type MovimientoParameter struct {
	ID                   string `json:"_id" bson:"_id,omitempty"`
	TipoMovimientoHijo   string `json:"TipoMovimientoHijo" bson:"tipo_movimiento_hijo"`
	TipoMovimientoPadre  string `json:"TipoMovimientoPadre" bson:"tipo_movimiento_padre"`
	Multiplicador        int    `json:"Multiplicador" bson:"miltiplicador"`
	FatherCollectionName string `json:"FatherCollectionName" bson:"father_collection_name"`
	FatherUUIKeyName     string `json:"FatherUUIKeyName" bson:"father_uuikey_name"`
	Initial              bool   `json:"Initial" bson:"initial"`
}
