package transactionManager

import (
	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/utils_oas/formatdata"
)

const TransactionCollection = "transactions"

// ConvertToTransactionItem ... This method manage the Mongo's transaction items from the current orm for registration of new elements in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToTransactionItem(collectionName string, models ...interface{}) (ops []interface{}) {
	return buildTransactionArr("d-", collectionName, models)
}

// ConvertToUpdateTransactionItem ... This method manage the Mongo's transaction items from the current orm for update elements that currently exist in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToUpdateTransactionItem(collectionName string, models ...interface{}) (ops []interface{}) {
	return buildTransactionArr("d+", collectionName, models)
}

// RunTransaction ... Perform a transaction over the DB with the options element.
func RunTransaction(collectionName string, options []interface{}) {

	session, _ := db.GetSession()

	defer func() {
		session.Close()
	}()
	var ops []txn.Op
	c := db.Cursor(session, TransactionCollection)
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

func buildTransactionItem(assertType, collectionName string, uuid string, model interface{}) (ops txn.Op) {
	ID := ""
	if uuid == "" {
		ID = bson.NewObjectId().Hex()
	} else {
		ID = uuid
	}

	if assertType == "d+" {
		var data bson.M
		formatdata.FillStructP(model, &data)

		op := txn.Op{
			C:      collectionName,
			Id:     ID,
			Assert: assertType,
			Update: bson.M{"$set": data},
		}
		return op

	}
	op := txn.Op{
		C:      collectionName,
		Id:     ID,
		Assert: assertType,
		Insert: model,
	}
	return op

}

func buildTransactionArr(assertType, collectionName string, models []interface{}) (ops []interface{}) {
	for _, model := range models {
		uuid := ""
		if assertType == "d+" {
			var modelMap map[string]interface{}
			formatdata.FillStructP(model, &modelMap)

			uuid = modelMap["_id"].(string)
		}

		ops = append(ops, buildTransactionItem(assertType, collectionName, uuid, model))
	}
	return
}
