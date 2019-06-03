package models

import (
	"fmt"

	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

const ArbolRubrosCollection = "arbol_rubro"

type ArbolRubros struct {
	Id               string   `json:"_id" bson:"_id,omitempty"`
	Idpsql           string   `json:"idpsql"`
	Nombre           string   `json:"nombre"`
	Descripcion      string   `json:"descripcion"`
	Hijos            []string `json:"hijos"`
	Padre            string   `json:"padre"`
	Unidad_Ejecutora string   `json:"unidad_ejecutora"`
}

func UpdateArbolRubros(session *mgo.Session, j ArbolRubros, id string) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err
}

func InsertArbolRubros(session *mgo.Session, j ArbolRubros) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	err := c.Insert(j)
	return err
}

func GetAllArbolRubross(session *mgo.Session, query map[string]interface{}) []ArbolRubros {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	// fmt.Println("Getting all arbolrubross")
	var arbolrubross []ArbolRubros
	err := c.Find(query).All(&arbolrubross)
	if err != nil {
		fmt.Println(err)
	}
	return arbolrubross
}

func GetArbolRubrosById(session *mgo.Session, id string) (ArbolRubros, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var arbolrubross ArbolRubros
	err := c.Find(bson.M{"_id": id}).One(&arbolrubross)
	return arbolrubross, err
}

func DeleteArbolRubrosById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	err := c.Remove(bson.M{"_id": id})
	return "ok", err
}

func GetArbolRubrosByIdPsql(session *mgo.Session, idPsql string) (ArbolRubros, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var rubro ArbolRubros
	err := c.Find(bson.M{"idpsql": idPsql}).One(&rubro)
	return rubro, err
}

/*
 Obtiene un nodo del arbol a partir de su id y su unidad ejecutora
*/
func GetNodo(session *mgo.Session, id, ue string) (ArbolRubros, error) {
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	var nodo ArbolRubros
	err := c.Find(bson.M{"_id": id, "unidad_ejecutora": bson.M{"$in": []string{"0", ue}}}).One(&nodo)
	return nodo, err
}

func RegistrarRubroTransacton(rubroPadre, rubroHijo ArbolRubros, session *mgo.Session) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	runner := txn.NewRunner(c)
	ops := []txn.Op{{
		C:      ArbolRubrosCollection,
		Id:     rubroHijo.Id,
		Assert: "d-",
		Insert: rubroHijo,
	}, {
		C:      ArbolRubrosCollection,
		Id:     rubroPadre.Id,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	}}
	id := bson.NewObjectId() // Optional
	err := runner.Run(ops, id, nil)
	return err
}

func EliminarRubroTransaccion(rubroPadre, rubroHijo ArbolRubros, session *mgo.Session) error {
	c := db.Cursor(session, ArbolRubrosCollection)
	runner := txn.NewRunner(c)
	ops := []txn.Op{{
		C:      ArbolRubrosCollection,
		Id:     rubroHijo.Id,
		Assert: "d+",
		Remove: true,
	}, {
		C:      ArbolRubrosCollection,
		Id:     rubroPadre.Id,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	}}
	id := bson.NewObjectId()
	err := runner.Run(ops, id, nil)
	return err
}

func GetRaices(session *mgo.Session, ue string) ([]ArbolRubros, error) {
	var (
		roots []ArbolRubros
	)
	c := db.Cursor(session, ArbolRubrosCollection)
	defer session.Close()
	// bson.M{ "$or": []bson.M{ bson.M{"padre": nil}, bson.M{"padre": } }, "idpsql": bson.M{"$ne": nil} }
	// err := c.Find(bson.M{"padre": nil, "idpsql": bson.M{"$ne": nil}}).All(&roots)
	err := c.Find(bson.M{
		"$or": []bson.M{bson.M{"padre": nil},
			bson.M{"padre": ""}},
		"idpsql":           bson.M{"$ne": nil},
		"unidad_ejecutora": bson.M{"$in": []string{"0", ue}},
	}).All(&roots)
	fmt.Println("roots: ", roots)
	return roots, err
}
