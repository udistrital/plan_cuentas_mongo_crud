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
	Productos    []string            `json:"Productos" bson:"productos`
	ValorTotal   float64             `json:"ValorTotal" bson:"ValorTotal"`
}

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	*General
	ID              string                       `json:"Codigo" bson:"_id,omitempty"`
	TipoFuente      interface{}                  `json:"TipoFuente" bson"tipoFuente"`
	ValorInicial    float64                      `json:"ValorInicial" bson:"valor_inicial"`
	ValorActual     float64                      `json:"ValorActual" bson:"valor_actual"`
	Estado          string                       `json:"Estado" bson:"estado"`
	Rubros          map[codigoRubro]*rubroFuente `json:"Rubros" bson:"rubros"`
	NumeroDocumento string                       `json:"NumeroDocumento" bson:"numeroDocumento"`
	TipoDocumento   string                       `json:"TipoDocumento" bson:"tipoDocumento"`
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
func GetFuenteFinanciamientoByID(id string) (*FuenteFinanciamiento, error) {
	var fuenteFinanciamiento FuenteFinanciamiento

	session, err := db.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	c := db.Cursor(session, FuenteFinanciamientoCollection)

	err = c.FindId(id).One(&fuenteFinanciamiento)

	return &fuenteFinanciamiento, err
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

// GetAllFuenteFinanciamiento obtiene todos los registros de fuente de financiamiento
func GetAllFuenteFinanciamiento(query map[string]interface{}) ([]FuenteFinanciamiento, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, err
	}

	c := db.Cursor(session, FuenteFinanciamientoCollection)
	defer session.Close()

	var fuentesFinanciamiento []FuenteFinanciamiento

	err = c.Find(query).All(&fuentesFinanciamiento)

	return fuentesFinanciamiento, err
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

// GetFuentesByRubroApropiacion devuelve todas las fuentes que tenga un rubro idRubroApropiacion
func GetFuentesByRubroApropiacion(idRubroApropiacion string) (fuentes []FuenteFinanciamiento, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}

	c := db.Cursor(session, FuenteFinanciamientoCollection)
	defer session.Close()

	c.Find(bson.M{"rubros." + idRubroApropiacion: bson.M{"$exists": "true"}}).All(&fuentes)
	return
}
