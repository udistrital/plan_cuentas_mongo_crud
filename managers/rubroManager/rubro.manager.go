package rubroManager

import (
	"errors"

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
	_, exist := SearchRubro(nodoHoja.ID, nodoHoja.UnidadEjecutora)
	if exist {
		panic("Rubro Code Already exist for this cg")
	}
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
			Update: bson.D{{"$set", bson.D{{"hijos", nodoPadre.Hijos}, {"bloqueado", true}}}},
		})
	}

	return runner.Run(ops, id, nil)
}

// TrEliminarNodoHoja transacción que elimina una hoja y actualiza el arreglo de hijos del padre
func TrEliminarNodoHoja(idNodoHoja, collection string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()
	nodo, _ := SearchRubro(idNodoHoja, "1")
	if nodo.Bloqueado {
		return errors.New("No se Puede ELiminar Este Rubro, Puede Que Tenga Rubros Hijo o Que Posea Apropiaciones Desiganadas")
	}
	ops := []txn.Op{{
		C:      collection,
		Id:     idNodoHoja,
		Assert: "d+",
		Remove: true,
	}}

	nodoHoja, err := models.GetNodoRubroById(idNodoHoja)

	if err != nil {
		return err
	}

	nodoPadre, err := models.GetNodoRubroById(nodoHoja.Padre)

	if err == nil {
		// nodoPadre.Hijos = append(nodoPadre.Hijos, nodoHoja.ID)
		nodoPadre.Hijos = remove(nodoPadre.Hijos, idNodoHoja)
		ops = append(ops, txn.Op{
			C:      collection,
			Id:     nodoPadre.ID,
			Assert: bson.M{"_id": nodoPadre.ID},
			Update: bson.D{{"$set", bson.D{{"hijos", nodoPadre.Hijos}}}},
		})
	}

	return runner.Run(ops, id, nil)
}

func remove(slice []string, element string) []string {
	newSlice := &slice
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			slice[i] = slice[len(slice)-1] // Copy last element to index i.
			slice[len(slice)-1] = ""       // Erase last element (write zero value).
			slice = slice[:len(slice)-1]
		}
	}
	return *newSlice
}

// SearchRubro ...  find rubro by code
func SearchRubro(nodo string, ue string) (models.NodoRubro, bool) {
	rubro, err := models.GetNodoRubroByIdAndUE(nodo, ue)
	if err != nil {
		return rubro, false
	}
	return rubro, true
}
