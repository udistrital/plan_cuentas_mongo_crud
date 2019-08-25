package rubroApropiacionManager

import (
	"strconv"

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

	nodoHoja.Estado = models.EstadoRegistrada

	ops := []txn.Op{{
		C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
		Id:     nodoHoja.ID,
		Assert: "d-",
		Insert: nodoHoja,
	}}

	nodoPadre, err := models.GetNodoRubroApropiacionById(nodoHoja.Padre, nodoHoja.UnidadEjecutora, nodoHoja.Vigencia)

	if err == nil {
		nodoPadre.Hijos = append(nodoPadre.Hijos, nodoHoja.ID)
		ops = append(ops, txn.Op{
			C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
			Id:     nodoPadre.ID,
			Assert: bson.M{"_id": nodoPadre.ID},
			Update: bson.D{{"$set", bson.D{{"nodorubro.hijos", nodoPadre.Hijos}}}},
		})
	}

	return runner.Run(ops, id, nil)
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
