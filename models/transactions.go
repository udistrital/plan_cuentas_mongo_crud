package models

import (
	"github.com/udistrital/plan_cuentas_mongo_crud/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// TrRegistroFuente Transacción para registrar todas las afectaciones de una fuente de financiamiento
// y en caso de ser necesario, también la información padre de dicha fuente de finaciamiento.
func TrRegistroFuente(session *mgo.Session, options []interface{}) {
	var ops []txn.Op
	c := db.Cursor(session, TransactionCollection)
	runner := txn.NewRunner(c)

	for _, op := range options {
		if op != nil {
			ops = append(ops, op.(txn.Op))
		}
	}

	id := bson.NewObjectId()
	err := runner.Run(ops, id, nil)
	if err != nil {
		panic(err.Error())
	}
}
