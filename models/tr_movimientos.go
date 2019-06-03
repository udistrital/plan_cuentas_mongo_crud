package models

import (
	"fmt"

	"github.com/manucorporat/try"

	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
)

const TransactionCollection = "transacciones"

func RegistrarMovimiento(session *mgo.Session, options []interface{}) (err error) {
	try.This(func() {
		var ops []txn.Op
		c := db.Cursor(session, TransactionCollection)
		runner := txn.NewRunner(c)

		for _, op := range options {
			ops = append(ops, op.(txn.Op))
		}

		id := bson.NewObjectId()
		if err = runner.Run(ops, id, nil); err != nil {
			fmt.Errorf("%s \n", err.Error())
			panic(err.Error())
		}
	}).Catch(func(e try.E) {
		fmt.Println("Error en RegistrarMovimiento: ", e)
		panic(e)
	})

	return err
}
