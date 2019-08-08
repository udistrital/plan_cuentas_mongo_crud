package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

type codigoRubro string

// dependenciaRubro Relación entre dependencia y rubro
type dependenciaRubro struct {
	ID    int     `json:"Id" bson:"idDepdencia"`
	Valor float64 `json:"ValorDependencia" bson:"valorDependencia"`
}

// rubroFuente Relación entre un rubro y una fuente
type rubroFuente struct {
	Dependencias []*dependenciaRubro `json:"Dependencias" bson:"dependencias"`
	Productos    []int               `json:"Productos" bson:"productos`
}

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	*General
	ID             string                       `json:"Codigo" bson:"_id,omitempty"`
	TipoFuente     interface{}                  `json:"TipoFuente" bson"tipoFuente"`
	ValorOriginal  float64                      `json:"ValorOriginal" bson:"valorOriginal"`
	ValorAcumulado float64                      `json:"ValorAcumulado" bson"valorAcumulado"`
	Rubros         map[codigoRubro]*rubroFuente `json:"Rubros" bson:"rubros"`
}

// FuenteFinanciamientoCollection constante para la colección
const FuenteFinanciamientoCollection = "fuente_financiamiento"

// // InsertFuenteMovimiento función para registrar un documento de tipo fuente_movimiento
// func InsertFuenteMovimiento(session *mgo.Session, j *FuenteMovimiento) {
// 	c := db.Cursor(session, fuenteFinanciamiento)
// 	c.Insert(&j)
// }

// InsertFuenteFinanciamiento función para registrar un documento de tipo fuente_financiamiento
func InsertFuenteFinanciamiento(j *FuenteFinanciamiento) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, FuenteFinanciamientoCollection)

	defer session.Close()
	return c.Insert(&j)
}

// GetFuenteFinanciamientoByID Obtener un documento por el id
func GetFuenteFinanciamientoByID(session *mgo.Session, id string) *FuenteFinanciamiento {
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	var fuenteFinanciamiento *FuenteFinanciamiento
	err := c.Find(bson.M{"_id": id}).One(&fuenteFinanciamiento)
	if err != nil {
		return nil
	}
	return fuenteFinanciamiento
}

// UpdateFuenteFinanciamiento actualiza una fuente de financiamiento
func UpdateFuenteFinanciamiento(j *FuenteFinanciamiento, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, FuenteFinanciamientoCollection)

	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}

// DeleteFuenteFinanciamiento elimina una fuente de financiamiento con su ID
func DeleteFuenteFinanciamiento(id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, FuenteFinanciamientoCollection)
	defer session.Close()

	return c.RemoveId(id)
}

// PostFuentePadreTransaccion crea una estructura para FuenteFinanciamiento de tipo registro.
func PostFuentePadreTransaccion(session *mgo.Session, estructura *FuenteFinanciamiento) (op txn.Op, err error) {
	estructura.ID = bson.NewObjectId().Hex()
	op = txn.Op{
		C:      FuenteFinanciamientoCollection,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}
