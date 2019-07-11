package transactionManager

import (
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
)

const TransactionCollection = "transaction"

// ConvertToTransactionItem ... This method manage the Mongo's transaction items from the current orm for registration of new elements in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToTransactionItem(collectionName string, model interface{}) (ops txn.Op) {
	return buildTransactionItem("d-", collectionName, model)
}

// ConvertToUpdateTransactionItem ... This method manage the Mongo's transaction items from the current orm for update elements that currently exist in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToUpdateTransactionItem(collectionName string, model interface{}) (ops txn.Op) {
	return buildTransactionItem("d+", collectionName, model)
}

// RunTransaction ... Perform a transaction over the DB with the options element.
func RunTransaction(collectionName string, options []interface{}) {

	session, _ := db.GetSession()

	defer func() {
		session.Close()
	}()
	var ops []txn.Op
	c := db.Cursor(session, collectionName+"_"+TransactionCollection)
	runner := txn.NewRunner(c)

	for _, op := range options {
		ops = append(ops, op.(txn.Op))
	}

	id := bson.NewObjectId()
	if err := runner.Run(ops, id, nil); err != nil {
		logManager.LogError(err.Error())
		panic(err.Error())
	}

}

func buildTransactionItem(assertType, collectionName string, model interface{}) (ops txn.Op) {
	ID := bson.NewObjectId().Hex()
	op := txn.Op{
		C:      collectionName,
		Id:     ID,
		Assert: assertType,
		Insert: model,
	}
	return op
}
