package rubroApropiacionManager

import (
	"fmt"
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// TrRegistrarNodoHoja transacción que registra una nueva hoja y modifica los hijos del padre
func TrRegistrarNodoHoja(nodoHoja *models.NodoRubroApropiacion, ue string, vigencia int) error {
	var ops []txn.Op

	fmt.Println(nodoHoja)

	session, err := db.GetSession()

	if err != nil {
		return err
	}

	c := db.Cursor(session, models.NodoRubroApropiacionCollection)
	runner := txn.NewRunner(c)

	nodoPadre, err := models.GetNodoRubroApropiacionById(nodoHoja.Padre, nodoHoja.UnidadEjecutora, nodoHoja.Vigencia)

	fmt.Println(nodoPadre.ID)

	if err != nil {
		return err
	}

	nodoPadre.Hijos = append(nodoPadre.Hijos, nodoHoja.ID)

	id := bson.NewObjectId()

	// opRegister := txn.Op{
	// 	C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
	// 	Id:     nodoHoja.ID,
	// 	Assert: "d-",
	// 	Insert: nodoHoja,
	// }

	opUpdate := txn.Op{
		C:      models.NodoRubroApropiacionCollection + "_" + strconv.Itoa(vigencia) + "_" + ue,
		Id:     nodoPadre.ID,
		Assert: "d-",
		Update: bson.D{{"$set", bson.D{{"nodorubro.hijos", nodoPadre.Hijos}}}},
	}

	// ops = append(ops, opRegister, opUpdate)
	ops = append(ops, opUpdate)
	// ops = append(ops, opRegister)

	return runner.Run(ops, id, nil)
}
