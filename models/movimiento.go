package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

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
	DocumentosPresGenerados   *[]string          `json:"DocumentosPresGenerados" bson:"documentos_pres_generados,omitempty"`
	RubroDetalle              *NodoRubroReducido
}

// MovimientoParameter this struct represent a "movmientos_parametros" collection item.
type MovimientoParameter struct {
	ID                     string  `json:"_id" bson:"_id,omitempty"`
	TipoMovimientoHijo     string  `json:"TipoMovimientoHijo" bson:"tipo_movimiento_hijo"`
	TipoMovimientoPadre    string  `json:"TipoMovimientoPadre" bson:"tipo_movimiento_padre"`
	Multiplicador          int     `json:"Multiplicador" bson:"miltiplicador"`
	FatherCollectionName   string  `json:"FatherCollectionName" bson:"father_collection_name"`
	FatherUUIKeyName       string  `json:"FatherUUIKeyName" bson:"father_uuikey_name"`
	Initial                bool    `json:"Initial" bson:"initial"`
	WithOutChangeState     bool    `json:"WithOutChangeState" bson:"without_change_state"`
	TipoDocumentoGenerado  *string `json:"TipoDocumentoGenerado" bson:"tipo_documento_generado,omitempty"`
	NoBalanceLeftStateName *string `json:"NoBalanceLeftStateName" bson:"no_balance_left_state_name,omitempty"`
}

// AfectacionMovimiento this struct will save modification's movs for test their behaivor before save on db
type AfectacionMovimiento struct {
	CuentaCredito       *NodoRubroApropiacion `json:"cuenta_credito"`
	CuentaContraCredito *NodoRubroApropiacion `json:"cuenta_contra_credito"`
	Tipo                string                `json:"tipo"`
	Valor               float64               `json:"valor"`
}

// GetMovmientoByParentId obtiene todos los movimientos relacionados
func GetMovmientoByParentId(vigencia, areaFuncional, id string) (Movimiento, error) {
	var movimiento Movimiento

	session, err := db.GetSession()
	if err != nil {
		return movimiento, err
	}

	collectionFixedName := MovimientosCollection + "_" + vigencia + "_" + areaFuncional

	c := db.Cursor(session, collectionFixedName)
	err = c.Find(bson.M{"padre": id}).One(&movimiento)

	defer session.Close()
	return movimiento, err
}
