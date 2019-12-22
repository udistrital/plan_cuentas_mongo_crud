package models

import (
	"errors"
	"regexp"
	"strconv"

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

// rubroFuente Relación entre un rubro y una fuente discriminada por un tipo(ingreso,gasto)
type rubroFuente struct {
	Dependencias []*dependenciaRubro `json:"Dependencias" bson:"dependencias"`
	Productos    []string            `json:"Productos" bson:"productos`
	ValorTotal   float64             `json:"ValorTotal" bson:"ValorTotal"`
	Tipo         string              `json:"Tipo" bson:"tipo`
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
	UnidadEjecutora string                       `json:"UnidadEjecutora" bson:"unidad_ejecutora"`
	Movimientos     map[string]interface{}       `json:"Movimientos" bson:"movimientos"`
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
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	if strconv.Itoa(j.Vigencia) != "0" {
		fuente, _ := GetFuenteFinanciamientoByID(j.ID, j.UnidadEjecutora, strconv.Itoa(j.Vigencia))
		if fuente.ID != "" {
			return errors.New("La Fuente de Financiamiento " + j.ID + " ya existe para la vigencia " + strconv.Itoa(j.Vigencia))
		}
		c = db.Cursor(session, FuenteFinanciamientoCollection+"_"+strconv.Itoa(j.Vigencia)+"_"+j.UnidadEjecutora)
	} else {
		fuente, _ := GetFuenteFinanciamientoByID(j.ID, "", "")
		if fuente.ID != "" {
			return errors.New("La Fuente de Financiamiento " + j.ID + " ya existe")
		}
	}
	if err != nil {
		return err
	}
	defer session.Close()
	return c.Insert(&j)
}

// GetFuenteFinanciamientoByID Obtener un documento por el id
func GetFuenteFinanciamientoByID(id string, ue string, vigencia string) (*FuenteFinanciamiento, error) {
	var fuenteFinanciamiento FuenteFinanciamiento
	session, err := db.GetSession()
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	if err != nil {
		return nil, err
	}
	if vigencia != "0" {
		c = db.Cursor(session, FuenteFinanciamientoCollection+"_"+vigencia+"_"+ue)
	}
	defer session.Close()

	err = c.FindId(id).One(&fuenteFinanciamiento)

	return &fuenteFinanciamiento, err
}

// UpdateFuenteFinanciamiento actualiza una fuente de financiamiento
func UpdateFuenteFinanciamiento(j *FuenteFinanciamiento, id string, ue string, vigencia string) error {
	session, err := db.GetSession()
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	if err != nil {
		return err
	}
	if vigencia != "0" {
		c = db.Cursor(session, FuenteFinanciamientoCollection+"_"+vigencia+"_"+ue)
	}
	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}

// DeleteFuenteFinanciamiento elimina una fuente de financiamiento con su ID
func DeleteFuenteFinanciamiento(id string, ue string, vigencia string) error {
	session, err := db.GetSession()
	c := db.Cursor(session, FuenteFinanciamientoCollection)
	if err != nil {
		return err
	}
	if vigencia != "0" {
		c = db.Cursor(session, FuenteFinanciamientoCollection+"_"+vigencia+"_"+ue)
		fuenteObj, _ := GetFuenteFinanciamientoByID(id, ue, vigencia)
		if fuenteObj.Estado == "distribuida" {
			return errors.New("No se puede eliminar esta fuente de financiamiento, debido a que esta distribuida")
		}
	}
	defer session.Close()

	return c.RemoveId(id)
}

// GetAllFuenteFinanciamiento obtiene todos los registros de fuente de financiamiento
func GetAllFuenteFinanciamiento(query map[string]interface{}, ue, vigencia string) ([]FuenteFinanciamiento, error) {
	session, err := db.GetSession()
	var fuentesFinanciamiento []FuenteFinanciamiento
	var fuentesFinanciamientoAux []FuenteFinanciamiento
	if vigencia == "all" {
		collections, _ := session.DB("").CollectionNames()
		collections = GetCollectionsByName(collections, "fuente_financiamiento")
		for _, itemColl := range collections {
			err = session.DB("").C(itemColl).Find(nil).All(&fuentesFinanciamientoAux)
			fuentesFinanciamiento = append(fuentesFinanciamiento, fuentesFinanciamientoAux...)
		}
		if err != nil {
			return nil, err
		}
	} else {

		c := db.Cursor(session, FuenteFinanciamientoCollection)
		if err != nil {
			return nil, err
		}

		if vigencia != "0" {
			c = db.Cursor(session, FuenteFinanciamientoCollection+"_"+vigencia+"_"+ue)
		}

		err = c.Find(query).All(&fuentesFinanciamiento)

	}
	defer session.Close()
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
func GetFuentesByRubroApropiacion(idRubroApropiacion string, ue string, vigencia string) (fuentes []FuenteFinanciamiento, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}

	c := db.Cursor(session, FuenteFinanciamientoCollection+"_"+vigencia+"_"+ue)
	defer session.Close()

	c.Find(bson.M{"rubros." + idRubroApropiacion: bson.M{"$exists": "true"}}).All(&fuentes)
	return
}

// GetCollectionsByName devuelve todas las colecciones encontradas por palabra clave
func GetCollectionsByName(collections []string, word string) []string {
	subMatches := []string{}
	for _, i := range collections {
		status, _ := regexp.MatchString("\\b"+word, i)
		if status {
			subMatches = append(subMatches, i)
		}
	}
	return subMatches
}
