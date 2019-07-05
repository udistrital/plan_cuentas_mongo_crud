package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// TrRegistroFuente Transacción para registrar todas las afectaciones de una fuente de financiamiento
// y en caso de ser necesario, también la información padre de dicha fuente de finaciamiento.
func TrRegistroFuente(session *mgo.Session, options []interface{}) {
	var ops []txn.Op
	c := db.Cursor(session, TransactionCollection)
	defer func() {
		session.Close()
		if r := recover(); r != nil {
			logs.Error(r)
			logManager.LogError(r)
			panic(r)
		}
	}()
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
