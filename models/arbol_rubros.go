package models

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

const NodoRubroCollection = "arbol_rubro"

// NodoRubro es la estructura de un Rubro, es un nodo puesto que forma parte del árbol
type NodoRubro struct {
	*General
	ID              string   `json:"Codigo" bson:"_id,omitempty"`
	Hijos           []string `json:"Hijos" bson:"hijos"`
	Padre           string   `json:"Padre" bson:"padre"`
	UnidadEjecutora string   `json:"UnidadEjecutora" bson:"unidad_ejecutora"`
	Bloqueado       bool     `json:"Bloqueado" bson:"bloqueado"`
}

func UpdateNodoRubro(j NodoRubro, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	return c.UpdateId(id, &j)
}

func InsertNodoRubro(session *mgo.Session, j NodoRubro) error {
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	err := c.Insert(j)
	return err
}

func GetAllNodoRubro(session *mgo.Session, query map[string]interface{}) []NodoRubro {
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	// fmt.Println("Getting all NodoRubros")
	var NodoRubros []NodoRubro
	err := c.Find(query).All(&NodoRubros)
	if err != nil {
		fmt.Println(err)
	}
	return NodoRubros
}

func GetNodoRubroById(id string) (NodoRubro, error) {
	var nodoRubro NodoRubro

	session, err := db.GetSession()
	if err != nil {
		return nodoRubro, err
	}

	c := db.Cursor(session, NodoRubroCollection)
	err = c.FindId(id).One(&nodoRubro)

	defer session.Close()
	return nodoRubro, err
}

/*
 Obtiene un nodo del arbol a partir de su id y su unidad ejecutora
*/
func GetNodoRubroByIdAndUE(id, ue string) (NodoRubro, error) {
	var nodoRubro NodoRubro
	session, err := db.GetSession()
	if err != nil {
		return nodoRubro, err
	}
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	err = c.Find(bson.M{"_id": id, "unidad_ejecutora": bson.M{"$in": []string{"0", ue}}}).One(&nodoRubro)
	return nodoRubro, err
}

func DeleteNodoRubroById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	err := c.Remove(bson.M{"_id": id})
	return "ok", err
}

func GetNodoRubroByIdPsql(session *mgo.Session, idPsql string) (NodoRubro, error) {
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	var rubro NodoRubro
	err := c.Find(bson.M{"idpsql": idPsql}).One(&rubro)
	return rubro, err
}

/*
 Obtiene un nodo del arbol a partir de su id y su unidad ejecutora
*/
func GetNodo(session *mgo.Session, id, ue string) (NodoRubro, error) {
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()
	var nodo NodoRubro
	err := c.Find(bson.M{"_id": id, "unidad_ejecutora": bson.M{"$in": []string{"0", ue}}}).One(&nodo)
	return nodo, err
}

func RegistrarRubroTransacton(rubroPadre, rubroHijo NodoRubro, session *mgo.Session) error {
	c := db.Cursor(session, NodoRubroCollection)
	runner := txn.NewRunner(c)
	ops := []txn.Op{{
		C:      NodoRubroCollection,
		Id:     rubroHijo.ID,
		Assert: "d-",
		Insert: rubroHijo,
	}, {
		C:      NodoRubroCollection,
		Id:     rubroPadre.ID,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	}}
	id := bson.NewObjectId() // Optional
	err := runner.Run(ops, id, nil)
	return err
}

func EliminarRubroTransaccion(rubroPadre, rubroHijo NodoRubro, session *mgo.Session) error {
	c := db.Cursor(session, NodoRubroCollection)
	runner := txn.NewRunner(c)
	ops := []txn.Op{{
		C:      NodoRubroCollection,
		Id:     rubroHijo.ID,
		Assert: "d+",
		Remove: true,
	}, {
		C:      NodoRubroCollection,
		Id:     rubroPadre.ID,
		Assert: "d+",
		Update: bson.D{{"$set", bson.D{{"hijos", rubroPadre.Hijos}}}},
	}}
	id := bson.NewObjectId()
	err := runner.Run(ops, id, nil)
	return err
}

func GetRaices(session *mgo.Session, ue string) ([]NodoRubro, error) {
	var (
		roots []NodoRubro
	)
	c := db.Cursor(session, NodoRubroCollection)
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

// GetHojasRubro devuelve todos los nodos cuyos hijos sean un arreglo vació
func GetHojasRubro() (leafs []NodoRubroApropiacion, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	c := db.Cursor(session, NodoRubroCollection)
	defer session.Close()

	err = c.Find(bson.M{"hijos": []string{}}).All(&leafs)
	return
}

// GetRubroCode devuelve el codigo normalizado del Rubro
func GetRubroCode(rubro string) string {
	if len(rubro) > 1 {
		var b bytes.Buffer
		numcaracters := getNumHyphen(rubro)
		b.WriteString(concatRubro(numcaracters, rubro))
		return b.String()
	}
	return rubro
}

func getNumHyphen(rubro string) (num int) {
	re := regexp.MustCompile(`\-`)
	num = len(re.FindAllString(rubro, -1))
	return
}

func concatRubro(numcaracters int, rubro string) (newrubro string) {
	// Pretend to return a rubro builded
	i := strings.LastIndex(rubro, "-")
	digits := len(strings.Replace(rubro[i:], "-", "", 1))
	switch numcaracters {
	case 2:
		if digits < 2 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-00", 1)
		} else if digits == 2 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-0", 1)
		} else {
			newrubro = rubro
		}
		break
	case 6:
		if digits < 2 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-000", 1)
		} else if digits == 2 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-00", 1)
		} else if digits == 3 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-0", 1)
		} else {
			newrubro = rubro
		}
		break
	default:
		if digits < 2 {
			newrubro = rubro[:i] + strings.Replace(rubro[i:], "-", "-0", 1)
		} else {
			newrubro = rubro
		}
	}
	return
}
