package rubroApropiacionManager

import (
	"fmt"
	"strconv"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego/logs"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// TrRegistrarNodoHoja transacción que registra una nueva hoja y modifica los hijos del padre
func TrRegistrarNodoHoja(nodoHoja *models.NodoRubroApropiacion, ue string, vigencia int) error {
	if nodoHoja.ValorInicial <= 0 {
		err := fmt.Errorf("Valor de la apropiación debe ser mayor a 0")
		return err
	}

	session, err := db.GetSession()
	if err != nil {
		return err
	}

	defer session.Close()
	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)

	searchRubro(nodoHoja.ID, ue, vigencia)
	nodoHoja.ValorActual = nodoHoja.ValorInicial
	ops := []txn.Op{{
		C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
		Id:     nodoHoja.ID,
		Assert: "d-",
		Insert: nodoHoja,
	}, {
		C:      models.NodoRubroCollection,
		Id:     nodoHoja.ID,
		Assert: bson.M{"_id": nodoHoja.ID},
		Update: bson.D{{"$set", bson.D{{"bloqueado", true}, {"apropiaciones", true}}}},
	}}

	if propOps, err := PropagarValorApropiacion(nodoHoja, nodoHoja.ValorInicial, ue, vigencia); err == nil {
		ops = append(ops, propOps...)
	}
	id := bson.NewObjectId()
	return runner.Run(ops, id, nil)
}

// TrActualizarValorApropiacion ... Actualiza el valor de una apropiacion y propaga el cambio en el arbol.
func TrActualizarValorApropiacion(nodo *models.NodoRubroApropiacion, objectID string, ue string, vigencia int) error {
	if nodo.ValorInicial <= 0 {
		err := fmt.Errorf("Valor de la apropiación debe ser mayor a 0")
		return err
	}

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
			Update: bson.D{{"$set", bson.D{{"valor_inicial", nodo.ValorInicial}, {"valor_actual", nodo.ValorInicial}}}},
		}}
		if nodoOldInfo.ValorInicial != nodo.ValorInicial {
			if propOps, err := PropagarValorApropiacion(nodo, nodo.ValorInicial-nodoOldInfo.ValorInicial, ue, vigencia); err == nil {
				ops = append(ops, propOps...)
			}
		}
		if nodo.Productos != nil {
			err := rubroApropiacionHelper.IsAprApproved(nodo)
			if err != nil {
				return err
			}
			err = rubroApropiacionHelper.IsMaxPercentProduct(nodo.Productos)
			if err != nil {
				return err
			}
			producOps := []txn.Op{{
				C:      collName,
				Id:     nodo.ID,
				Assert: bson.M{"_id": nodo.ID},
				Update: bson.D{{"$set", bson.D{{"productos", nodo.Productos}}}},
			}}
			ops = append(ops, producOps...)
		}

		return runner.Run(ops, id, nil)
	}
	return err

}

func searchRubro(nodo string, ue string, vigencia int) models.NodoRubro {
	rubroPadre, err := models.GetNodoRubroByIdAndUE(nodo, ue)
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
			nodoPadre.NodoRubro.General = &models.General{Vigencia: vigencia}
			nodoPadre.ValorActual = nodoPadre.ValorInicial
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: "d-",
				Insert: nodoPadre,
			})
		} else {
			nodoPadre.ValorInicial += propagationValue
			nodoPadre.ValorActual = nodoPadre.ValorInicial
			ops = append(ops, txn.Op{
				C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
				Id:     nodoPadre.ID,
				Assert: bson.M{"_id": nodoPadre.ID},
				Update: bson.D{{"$set", bson.D{{"valor_inicial", nodoPadre.ValorInicial}, {"valor_actual", nodoPadre.ValorInicial}}}},
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

// TrAprobarApropiaciones actualiza todas las apropiaciones que tengan el estado registrada por el estado aprobada
func TrAprobarApropiaciones(ue, vigencia string) error {
	var ops []txn.Op
	var query = map[string]interface{}{"estado": "registrada"}

	apropiacionesRegistradas := models.GetAllNodoRubroApropiacion(query, ue, vigencia)
	for _, apropiacion := range apropiacionesRegistradas {
		ops = append(ops, txn.Op{
			C:      models.NodoRubroApropiacionCollection + "_" + vigencia + "_" + ue,
			Id:     apropiacion.ID,
			Assert: bson.M{"_id": apropiacion.ID},
			Update: bson.D{{"$set", bson.D{{"estado", models.EstadoAprobada}}}},
		})
	}

	session, err := db.GetSession()
	if err != nil {
		return err
	}
	defer session.Close()
	c := db.Cursor(session, models.TransactionCollection)
	runner := txn.NewRunner(c)
	id := bson.NewObjectId()

	return runner.Run(ops, id, nil)
}
