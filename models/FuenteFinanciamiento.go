package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	ID              string      `json:"_id" bson:"_id,omitempty"`
	UnidadEjecutora int         `json:"unidad_ejecutora"`
	Descripcion     string      `json:"descripcion"`
	IDPsql          int         `json:"idpsql"`
	Nombre          string      `json:"nombre"`
	TipoFuente      interface{} `json:"tipo_fuente"`
	ValorOriginal   float64     `json:"valor_original"`
	ValorAcumulado  float64     `json:"valor_acumulado"`
}

// FuenteMovimiento ...
type FuenteMovimiento struct {
	ID                string  `orm:"size(128)"`
	IDPsql            string  `json:"idpsql"`
	Rubro             string  `json:"rubro"`
	DependenciaIDPsql string  `json:"dependencia_idpsql"`
	ValorAcumulado    float64 `json:"valor_acumulado"`
}

// ArbolRubroApropiacion2018Collection constante para la colección
const fuenteFinanciamiento = "fuente_financiamiento"
const fuenteMovimiento = "fuente_movimiento"

// InsertFuenteMovimiento función para registrar un documento de tipo fuente_movimiento
func InsertFuenteMovimiento(session *mgo.Session, j *FuenteMovimiento) {
	c := db.Cursor(session, fuenteFinanciamiento)
	c.Insert(&j)
}

// InsertFuentFinanciamientoPadre función para registrar un documento de tipo fuente_financiamiento
func InsertFuentFinanciamientoPadre(session *mgo.Session, j *FuenteFinanciamiento) {
	c := db.Cursor(session, fuenteFinanciamiento)
	c.Insert(&j)
}

// GetFuenteFinanciamientoByID Obtener un documento por el id
func GetFuenteFinanciamientoByID(session *mgo.Session, id string) *FuenteFinanciamiento {
	c := db.Cursor(session, fuenteFinanciamiento)
	var fuenteFinanciamiento *FuenteFinanciamiento
	err := c.Find(bson.M{"_id": id}).One(&fuenteFinanciamiento)
	if err != nil {
		return nil
	}
	return fuenteFinanciamiento
}

// GetFuenteFinanciamientoByIDPsql Obtener un documento por el idpsql
func GetFuenteFinanciamientoByIDPsql(session *mgo.Session, id int) *FuenteFinanciamiento {
	c := db.Cursor(session, fuenteFinanciamiento)
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
		C:      fuenteFinanciamiento,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}
