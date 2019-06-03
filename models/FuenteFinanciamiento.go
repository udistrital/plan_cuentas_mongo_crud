package models

import (
	"github.com/udistrital/financiera_mongo_crud/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// FuenteFinaciamientoPadre ...
type FuenteFinaciamientoPadre struct {
	ID              string      `json:"_id" bson:"_id,omitempty"`
	UnidadEjecutora int         `json:"unidad_ejecutora"`
	Descripcion     string      `json:"descripcion"`
	IDPsql          int         `json:"idpsql"`
	Nombre          string      `json:"nombre"`
	TipoFuente      interface{} `json:"tipo_fuente"`
	ValorOriginal   float64     `json:"valor_original"`
}

// FuenteMovimiento ...
type FuenteMovimiento struct {
	ID                string  `orm:"size(128)"`
	IDPsql            string  `json:"idpsql"`
	Rubro             string  `json:"rubro"`
	DependenciaIDPsql string  `json:"dependencia_idpsql"`
	Saldo             float64 `json:"saldo"`
}

// ArbolRubroApropiacion2018Collection constante para la colección
const fuenteFinanciamientoPadre = "fuente_financiamiento_padre"
const fuenteMovimiento = "fuente_movimiento"

// InsertFuenteMovimiento función para registrar un documento de tipo fuente_movimiento
func InsertFuenteMovimiento(session *mgo.Session, j *FuenteMovimiento) {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	c.Insert(&j)
}

// InsertFuentFinanciamientoPadre función para registrar un documento de tipo fuente_financiamiento_padre
func InsertFuentFinanciamientoPadre(session *mgo.Session, j *FuenteFinaciamientoPadre) {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	c.Insert(&j)
}

// GetFuenteFinanciamientoPadreByID Obtener un documento por el id
func GetFuenteFinanciamientoPadreByID(session *mgo.Session, id string) *FuenteFinaciamientoPadre {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	var fuenteFinaciamientoPadre *FuenteFinaciamientoPadre
	err := c.Find(bson.M{"_id": id}).One(&fuenteFinaciamientoPadre)
	if err != nil {
		return nil
	}
	return fuenteFinaciamientoPadre
}

// GetFuenteFinanciamientoPadreByIDPsql Obtener un documento por el idpsql
func GetFuenteFinanciamientoPadreByIDPsql(session *mgo.Session, id int) *FuenteFinaciamientoPadre {
	c := db.Cursor(session, fuenteFinanciamientoPadre)
	var fuenteFinaciamientoPadre *FuenteFinaciamientoPadre
	err := c.Find(bson.M{"idpsql": id}).One(&fuenteFinaciamientoPadre)
	if err != nil {
		return nil
	}
	return fuenteFinaciamientoPadre
}

// EstructaRegistroFuentePadreTransaccion crea una estructura para FuenteFinanciamientoPadre de tipo registro.
func EstructaRegistroFuentePadreTransaccion(session *mgo.Session, estructura *FuenteFinaciamientoPadre) (op txn.Op, err error) {
	estructura.ID = bson.NewObjectId().Hex()
	op = txn.Op{
		C:      fuenteFinanciamientoPadre,
		Id:     estructura.ID,
		Assert: "d-",
		Insert: estructura,
	}
	return op, err
}
