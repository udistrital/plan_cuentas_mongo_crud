package rubroApropiacionManager

import (
	"strconv"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego/logs"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// TrRegistrarNodoHoja transacción que registra una nueva hoja y modifica los hijos del padre
func TrRegistrarNodoHoja(nodoHoja *models.NodoRubroApropiacion, ue string, vigencia int) error {

	session, err := db.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()
	searchRubro(nodoHoja.ID, ue, vigencia)
	ops := []txn.Op{{
		C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
		Id:     nodoHoja.ID,
		Assert: "d-",
		Insert: nodoHoja,
	}}

	if propOps, err := PropagarValorApropiacion(nodoHoja, nodoHoja.ApropiacionInicial, ue, vigencia); err == nil {

		ops = append(ops, propOps...)
	}
	return runner.Run(ops, id, nil)
}

// TrActualizarValorApropiacion ... Actualiza el valor de una apropiacion y propaga el cambio en el arbol.
func TrActualizarValorApropiacion(nodo *models.NodoRubroApropiacion, objectID string, ue string, vigencia int) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()
	if nodoOldInfo, err := models.GetNodoRubroApropiacionById(nodo.ID, ue, vigencia); err == nil {
		collName := models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue

		ops := []txn.Op{{
			C:      collName,
			Id:     nodo.ID,
			Assert: bson.M{"_id": nodo.ID},
			Update: bson.D{{"$set", bson.D{{"apropiacionInicial", nodo.ApropiacionInicial}}}},
		}}
		if nodoOldInfo.ApropiacionInicial != nodo.ApropiacionInicial {
			if propOps, err := PropagarValorApropiacion(nodo, nodo.ApropiacionInicial-nodoOldInfo.ApropiacionInicial, ue, vigencia); err == nil {
				ops = append(ops, propOps...)
			}
		}

		return runner.Run(ops, id, nil)
	} else {
		return err
	}

}

func searchRubro(nodo string, ue string, vigencia int) models.NodoRubro {
	rubroPadre, err := models.GetNodoRubroById(nodo)
	if err != nil {
		message := err.Error()
		if message == "not found" {
			message = "Rubro " + nodo + " Does not exist!"
		}
		logs.Error(message)
		panic(message)
	}
	return rubroPadre
}

// PropagarValorApropiacion ... Propaga valores dentro del arbol de apropiaciones a partir de un nodo. El nodo debe tener los campos vigencia y unidad ejecutora.
func PropagarValorApropiacion(nodoHijo *models.NodoRubroApropiacion, propagationValue float64, ue string, vigencia int) (ops []txn.Op, err error) {
	var nodo models.NodoRubroApropiacion
	formatdata.FillStructP(nodoHijo, &nodo)

	for nodo.Padre != "" {
		var nodoPadre models.NodoRubroApropiacion

		nodoPadrePointer, err := models.GetNodoRubroApropiacionById(nodo.Padre, ue, vigencia)
		formatdata.FillStructP(nodoPadrePointer, &nodoPadre)

		if err != nil && err.Error() == "not found" {
			rubroPadre := searchRubro(nodo.Padre, ue, vigencia)
			nodoPadre = nodo
			nodoPadre.NodoRubro = &rubroPadre
			nodoPadre.ID = rubroPadre.ID
			nodoPadre.Padre = rubroPadre.Padre
			nodoPadre.Hijos = []string{nodo.ID}
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: "d-",
				Insert: nodoPadre,
			})
		} else {
			nodoPadre.ApropiacionInicial += propagationValue
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: bson.M{"_id": nodoPadre.ID},
				Update: bson.D{{"$set", bson.D{{"apropiacionInicial", nodoPadre.ApropiacionInicial}}}},
			})
			if childrenExist(nodo.ID, nodoPadre.Hijos) != true {

				nodoPadre.Hijos = append(nodoPadre.Hijos, nodo.ID)

				ops = append(ops, txn.Op{
					C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
					Id:     nodoPadre.ID,
					Assert: bson.M{"_id": nodoPadre.ID},
					Update: bson.D{{"$set", bson.D{{"nodorubro.hijos", nodoPadre.Hijos}}}},
				})
			}
		}
		nodo = nodoPadre
	}

	return
}

func childrenExist(childCode string, childrenArray []string) bool {
	for _, child := range childrenArray {
		if child == childCode {
			return true
		}
	}
	return false
}
