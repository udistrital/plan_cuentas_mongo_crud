package rubroApropiacionManager

import (
	"fmt"
	// "log"
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

	c := db.Cursor(session, models.NodoRubroApropiacionCollection)
	runner := txn.NewRunner(c)

	id := bson.NewObjectId()

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
			C:  models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
			Id: nodoPadre.ID,
			Assert: bson.M{"_id": nodoPadre.ID},
			Update: bson.D{{"$set", bson.D{{"nodorubro.hijos", nodoPadre.Hijos}}}},
		})
	} 

	return runner.Run(ops, id, nil)
}
