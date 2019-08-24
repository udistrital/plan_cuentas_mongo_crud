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

	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()

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

// PropagarValorApropiacion ... Propaga valores dentro del arbol de apropiaciones a partir de un nodo. El nodo debe tener los campos vigencia y unidad ejecutora.
func PropagarValorApropiacion(nodoHijo *models.NodoRubroApropiacion, propagationValue float64, ue string, vigencia int) (ops []txn.Op, err error) {
	var nodo models.NodoRubroApropiacion
	var nodoPadre models.NodoRubroApropiacion
	formatdata.FillStructP(nodoHijo, &nodo)

	for nodo.Padre != "" {
		nodoPadrePointer, err := models.GetNodoRubroApropiacionById(nodo.Padre, ue, vigencia)
		formatdata.FillStructP(nodoPadrePointer, &nodoPadre)

		if err != nil && err.Error() == "not found" {
			rubroPadre, err := models.GetNodoRubroById(nodo.Padre)
			if err != nil {
				message := err.Error()
				if message == "not found" {
					message = "Rubro " + nodo.Padre + " Does not exist!"
				}
				logs.Error(message)
				panic(message)
			}
			nodoPadre = nodo
			nodoPadre.NodoRubro = &rubroPadre
			nodoPadre.ID = rubroPadre.ID
			nodoPadre.Padre = rubroPadre.Padre
			nodoPadre.NodoRubro.Hijos = []string{nodo.ID}
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: "d-",
				Insert: nodoPadre,
			})
		} else {
			nodoPadre.ApropiacionInicial += propagationValue
			logs.Debug("padre", nodoPadre.ID, "hijos", nodoPadre.Hijos, "agregar", nodo.ID)
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: "d+",
				Update: bson.D{{"$set", bson.D{{"apropiacionInicial", nodoPadre.ApropiacionInicial}}}},
			})
			if !childrenExist(nodo.ID, nodoPadre.Hijos) {
				nodoPadre.Hijos = append(nodoPadre.Hijos, nodo.ID)
				ops = append(ops, txn.Op{
					C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
					Id:     nodoPadre.ID,
					Assert: "d+",
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
