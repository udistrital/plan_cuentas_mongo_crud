package models

import (
	"fmt"

	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// ArbolRubroApropiacion2018Collection constante para la colección
const ArbolRubroApropiacion2018Collection = "arbolrubroapropiacion2018"

// ArbolRubroApropiacionCollection constante para la colección
const ArbolRubroApropiacionCollection = "arbolrubroapropiacion"

// ArbolRubroApropiacion es la estructura del documento que se va a registrar
type ArbolRubroApropiacion struct {
	Id                  string                        `json:"_id" bson:"_id,omitempty"`
	Idpsql              string                        `json:"idpsql"`
	Nombre              string                        `json:"nombre"`
	Descripcion         string                        `json:"descripcion"`
	Unidad_ejecutora    string                        `json:"unidad_ejecutora"`
	Padre               string                        `json:"padre"`
	Hijos               []string                      `json:"hijos"`
	Apropiacion_inicial int                           `json:"apropiacion_inicial"`
	Movimientos         map[string]map[string]float64 `json:"movimientos"`
}

func GetAllArbolRubroApropiacion(session *mgo.Session, query map[string]interface{}, ue, vigencia string) []ArbolRubroApropiacion {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	var arbolRubroApropiacion []ArbolRubroApropiacion
	err := c.Find(query).All(&arbolRubroApropiacion)
	if err != nil {
		fmt.Println(err)
	}
	return arbolRubroApropiacion
}

// UpdateArbolRubroApropiacion Update function to ArbolRubroApropiacion
func UpdateArbolRubroApropiacion(session *mgo.Session, j ArbolRubroApropiacion, id, ue, vigencia string) error {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	// Update
	fmt.Println("id update: ", id)
	err := c.Update(bson.M{"_id": id}, &j)
	/*if err != nil {
		fmt.Println("updatw error")
		panic(err)
	}*/
	return err

}

// InsertArbolRubroApropiacion Register function to ArbolRubroApropiacion
func InsertArbolRubroApropiacion(session *mgo.Session, j *ArbolRubroApropiacion, ue, vigencia string) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	c.Insert(&j)

}

// GetArbolRubroApropiacionById Obtener un documento por el id
func GetArbolRubroApropiacionById(session *mgo.Session, id, ue, vigencia string) (*ArbolRubroApropiacion, error) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	var arbolRubroApropiacion *ArbolRubroApropiacion
	err := c.Find(bson.M{"_id": id}).One(&arbolRubroApropiacion)
	return arbolRubroApropiacion, err
}

func DeleteArbolRubroApropiacion2018ById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ArbolRubroApropiacion2018Collection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}

func GetNodoApropiacion(session *mgo.Session, id, ue, vigencia string) (ArbolRubroApropiacion, error) {
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	var nodo ArbolRubroApropiacion
	err := c.Find(bson.M{"_id": id}).One(&nodo)
	return nodo, err
}

func GetRaicesApropiacion(session *mgo.Session, ue, vigencia string) ([]ArbolRubroApropiacion, error) {
	var roots []ArbolRubroApropiacion
	c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	defer session.Close()
	err := c.Find(bson.M{
		"$or": []bson.M{bson.M{"padre": nil},
			bson.M{"padre": ""}},
		"idpsql":           bson.M{"$ne": nil},
		"unidad_ejecutora": bson.M{"$in": []string{"0", ue}},
	}).All(&roots)
	// fmt.Println("roots: ", roots)
	return roots, err
}

func EstrctTransaccionArbolApropiacion(session *mgo.Session, estructuras []*ArbolRubroApropiacion, ue, vigencia string) (ops []txn.Op, err error) {
	//var ops []txn.Op
	// c := db.Cursor(session, ArbolRubroApropiacionCollection+"_"+vigencia+"_"+ue)
	// runner := txn.NewRunner(c)
	for _, estructura := range estructuras {
		op := txn.Op{
			C:      ArbolRubroApropiacionCollection + "_" + vigencia + "_" + ue,
			Id:     estructura.Id,
			Assert: "d+",
			Update: bson.D{{"$set", bson.D{{"movimientos", estructura.Movimientos}}}},
		}
		ops = append(ops, op)
	}
	// id := bson.NewObjectId()
	// err = runner.Run(ops, id, nil)

	return ops, err
}
