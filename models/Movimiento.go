package models

import (
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// MovimientosCollection es el nombre de la colección en mongo.
const MovimientosCollection = "movimientos"

// MovimientoCdp es una estructura generica para los tipos de movimiento registados.
type MovimientoCdp struct {
	ID              string                   `json:"_id" bson:"_id,omitempty"`
	IDPsql          string                   `json:"idpsql"`
	RubrosAfecta    []map[string]interface{} `json:"rubros_afecta"`
	ValorOriginal   float64                  `json:"valor_original"`
	Tipo            string                   `json:"tipo"`
	Vigencia        string                   `json:"vigencia"`
	DocumentoPadre  string                   `json:"documento_padre"`
	FechaRegistro   string                   `json:"fecha_registro"`
	UnidadEjecutora string                   `json:"unidad_ejecutora"`
}

// GetMovimientoByPsqlId Obtener un documento por el idpsql
func GetMovimientoByPsqlId(session *mgo.Session, id, tipo string) (*MovimientoCdp, error) {
	c := db.Cursor(session, MovimientosCollection)
	defer session.Close()
	var movimientoCdp *MovimientoCdp
	err := c.Find(bson.M{"idpsql": id, "tipo": tipo}).One(&movimientoCdp)
	return movimientoCdp, err
}

// EstrctTransaccionMov crea una transacción para MovimientoCdp de tipo registro.
func EstrctTransaccionMov(session *mgo.Session, estructura *MovimientoCdp) (ops txn.Op, err error) {
	estructura.ID = bson.NewObjectId().Hex()
	op := txn.Op{
		C:      MovimientosCollection,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}

// EstrctUpdateTransaccionMov crea una transacción para MovimientoCdp de tipo update.
func EstrctUpdateTransaccionMov(session *mgo.Session, estructura *MovimientoCdp) (ops txn.Op, err error) {
	op := txn.Op{
		C:      MovimientosCollection,
		Id:     estructura.ID,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"rubrosafecta", estructura.RubrosAfecta}}}},
	}
	return op, err
}
