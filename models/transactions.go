package models

import (
	"fmt"

	"github.com/manucorporat/try"
	"github.com/udistrital/financiera_mongo_crud/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

// TrRegistroFuente Transacción para registrar todas las afectaciones de una fuente de financiamiento
// y en caso de ser necesario, también la información padre de dicha fuente de finaciamiento.
func TrRegistroFuente(session *mgo.Session, options []interface{}) (err error) {
	try.This(func() {
		var ops []txn.Op
		c := db.Cursor(session, TransactionCollection)
		runner := txn.NewRunner(c)

		for _, op := range options {
			if op != nil {
				ops = append(ops, op.(txn.Op))
			}
		}

		id := bson.NewObjectId()
		if err = runner.Run(ops, id, nil); err != nil {
			errorString := fmt.Errorf("%s \n", err.Error())
			fmt.Println(errorString)
			panic(err.Error())
		}
	}).Catch(func(e try.E) {
		fmt.Println("Error en TrRegistroFuente: ", e)
		panic(e)
	})

	return err
}
