package rubroManager

import (
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// GetRaices Returns the rubro's tree roots.
func GetRaices(ue string) []map[string]interface{} {
	var (
		roots []models.NodoRubro
	)
	session, _ := db.GetSession()
	c := db.Cursor(session, models.NodoRubroCollection)
	defer func() {
		session.Close()
		if r := recover(); r != nil {
			logs.Error(r)
			logManager.LogError(r)
			panic(r)
		}
	}()

	err := c.Find(bson.M{"padre": "", "unidad_ejecutora": "1"}).All(&roots)
	if err != nil {
		panic(err.Error())
	}
	var resul []map[string]interface{}
	formatdata.FillStructP(roots, &resul)
	return resul
}

// GetHijoRubro this function should return branches from a tree's root.
func GetHijoRubro(id, ue string) map[string]interface{} {
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetNodo(session, id, ue)
	hijo := make(map[string]interface{})

	if rubroHijo.ID != "" {
		hijo["Codigo"] = rubroHijo.ID
		hijo["Nombre"] = rubroHijo.Nombre
		hijo["IsLeaf"] = false
		hijo["UnidadEjecutora"] = rubroHijo.UnidadEjecutora
		if len(rubroHijo.Hijos) == 0 {
			hijo["IsLeaf"] = true
			hijo["Hijos"] = nil
			return hijo
		}
	}
	return hijo
}

/*
 Obtiene un nodo del arbol a partir de su id y su unidad ejecutora
*/
func GetNodo(id, ue string) map[string]interface{} {
	session, _ := db.GetSession()
	c := db.Cursor(session, models.NodoRubroCollection)
	defer func() {
		session.Close()
		if r := recover(); r != nil {
			logManager.LogError(r)
			panic(r)
		}
	}()

	var nodo models.NodoRubro
	err := c.Find(bson.M{"_id": id, "unidad_ejecutora": ue}).One(&nodo)

	if err != nil {
		panic("Cannot Find Node " + id)
	}
	var resul map[string]interface{}
	formatdata.FillStructP(nodo, &resul)
	return resul
}

// TrRegistrarNodoHoja transacción que registra una nueva hoja y modifica los hijos del padre
func TrRegistrarNodoHoja(nodoHoja *models.NodoRubro, collection string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()

	ops := []txn.Op{{
		C:      collection,
		Id:     nodoHoja.ID,
		Assert: "d-",
		Insert: nodoHoja,
	}}

	nodoPadre, err := models.GetNodoRubroById(nodoHoja.Padre)

	if err == nil {
		nodoPadre.Hijos = append(nodoPadre.Hijos, nodoHoja.ID)

		ops = append(ops, txn.Op{
			C:      collection,
			Id:     nodoPadre.ID,
			Assert: bson.M{"_id": nodoPadre.ID},
			Update: bson.D{{"$set", bson.D{{"hijos", nodoPadre.Hijos}}}},
		})
	}

	return runner.Run(ops, id, nil)
}
