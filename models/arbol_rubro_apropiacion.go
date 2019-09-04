package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// NodoRubroApropiacionCollection constante para la colección
const NodoRubroApropiacionCollection = "arbol_rubro_apropiacion"

/* 	EstadoAprobada el estado aprobado de una apropiación, significa que la apropiación ya ha sido aprobada
EstadoRegistrada el estado registrada de un apropiación, significa que sólo se ha registrado
EstadoRechazada el estado rechazada de una apropiación, significa que se la apropiación ha sido rechazada */
const EstadoAprobada, EstadoRegistrada, EstadoRechazada = "aprobada", "registrada", "rechazada"

// NodoRubroApropiacion es la estructura de un nodo rubro pero sumandole la apropiación
type NodoRubroApropiacion struct {
	*NodoRubro
	ID                   string                            `json:"Codigo" bson:"_id,omitempty"`
	ApropiacionInicial   float64                           `json:"ApropiacionInicial" bson:"apropiacionInicial"`
	ApropiacionUtilizada float64                           `json:"ApropiacionUtilizada" bson:"apropiacionUtilizada"`
	Movimientos          []string                          `json:"Movimientos" bson:"movimientos"`
	Productos            map[string]map[string]interface{} `json:"Productos" bson:"productos"`
	Estado               string                            `json:"Estado" bson:"estado"`
}

func GetAllNodoRubroApropiacion(query map[string]interface{}, ue, vigencia string) []NodoRubroApropiacion {
	var NodoRubroApropiacion []NodoRubroApropiacion

	session, err := db.GetSession()
	if err != nil {
		log.Println(err.Error())
		return NodoRubroApropiacion
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()

	err = c.Find(query).All(&NodoRubroApropiacion)
	if err != nil {
		log.Println(err.Error())
	}
	return NodoRubroApropiacion
}

// UpdateNodoRubroApropiacion Update function to NodoRubroApropiacion
func UpdateNodoRubroApropiacion(session *mgo.Session, j NodoRubroApropiacion, id, ue string, vigencia int) error {
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+strconv.Itoa(vigencia)+"_"+ue)
	defer session.Close()
	return c.Update(bson.M{"_id": id}, &j)
}

// InsertNodoRubroApropiacion Register function to NodoRubroApropiacion
func InsertNodoRubroApropiacion(j *NodoRubroApropiacion) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+strconv.Itoa(j.Vigencia)+"_"+j.UnidadEjecutora)
	defer session.Close()
	return c.Insert(&j)
}

// GetNodoRubroApropiacionById Obtener un documento por el id
func GetNodoRubroApropiacionById(id, ue string, vigencia int) (*NodoRubroApropiacion, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, err
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+strconv.Itoa(vigencia)+"_"+ue)
	var nodoRubroApropiacion *NodoRubroApropiacion
	err = c.FindId(id).One(&nodoRubroApropiacion)

	defer session.Close()
	return nodoRubroApropiacion, err
}

func GetNodoRubroApropiacionByState(id, ue, vigencia, estado string) (*NodoRubroApropiacion, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+vigencia+"_"+ue)

	var nodoRubroApropiacion *NodoRubroApropiacion
	fmt.Println("id: ", id)
	err = c.Find(bson.M{"estado": estado, "_id": id}).One(&nodoRubroApropiacion)
	return nodoRubroApropiacion, err
}

func DeleteNodoRubroApropiacionById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, NodoRubroApropiacionCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}

func GetNodoApropiacion(id, ue string, vigencia int) (nodo NodoRubroApropiacion, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+strconv.Itoa(vigencia)+"_"+ue)
	defer session.Close()
	err = c.Find(bson.M{"_id": id}).One(&nodo)
	return
}

func GetRaicesApropiacion(ue string, vigencia int) (roots []NodoRubroApropiacion, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+strconv.Itoa(vigencia)+"_"+ue)
	defer session.Close()
	err = c.Find(bson.M{
		"$or": []bson.M{bson.M{"nodorubro.padre": nil}, bson.M{"nodorubro.padre": ""}},
	}).All(&roots)
	return
}

func EstrctTransaccionArbolApropiacion(session *mgo.Session, estructuras []*NodoRubroApropiacion, ue string, vigencia int) (ops []txn.Op, err error) {
	//var ops []txn.Op
	// c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	// runner := txn.NewRunner(c)
	for _, estructura := range estructuras {
		op := txn.Op{
			C:      NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
			Id:     estructura.ID,
			Assert: "d+",
			Update: bson.D{{"$set", bson.D{{"movimientos", estructura.Movimientos}}}},
		}
		ops = append(ops, op)
	}
	// id := bson.NewObjectId()
	// err = runner.Run(ops, id, nil)

	return ops, err
}

// GetHojasApropiacion devuelve todos los nodos cuyos hijos sean un arreglo vació
func GetHojasApropiacion(ue, vigencia string) (leafs []NodoRubroApropiacion, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	c := db.Cursor(session, NodoRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()

	err = c.Find(bson.M{"nodorubro.hijos": []string{}}).All(&leafs)
	return
}
