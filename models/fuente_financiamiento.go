package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// rubroFuente Relación entre un rubro y una fuente
type rubroFuente struct {
	Valor        float64 `json:"Valor" bson"valor"`
	Productos    []int   `json:"Productos" bson:"productos`
	Dependencias []int   `json:"Dependencias" bson:"dependencias"`
}

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	*General
	ID             string                             `json:"_id" bson:"_id,omitempty"`
	TipoFuente     interface{}                        `json:"TipoFuente" bson"tipoFuente"`
	ValorOriginal  float64                            `json:"ValorOriginal" bson:"valorOriginal"`
	ValorAcumulado float64                            `json:"ValorAcumulado" bson"valorAcumulado"`
	Rubros         map[string]map[string]*rubroFuente `json:"Rubros" bson:"rubros"`
}

// ArbolRubroApropiacion2018Collection constante para la colección
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
	defer session.Close()

	c := db.Cursor(session, FuenteFinanciamientoCollection)
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

// GetFuenteFinanciamientoByIDPsql Obtener un documento por el idpsql
func GetFuenteFinanciamientoByIDPsql(session *mgo.Session, id int) *FuenteFinanciamiento {
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	var fuenteFinanciamiento *FuenteFinanciamiento
	err := c.Find(bson.M{"idpsql": id}).One(&fuenteFinanciamiento)
	if err != nil {
		return nil
	}
	return fuenteFinanciamiento
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
