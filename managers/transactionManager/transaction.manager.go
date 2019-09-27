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
func ConvertToTransactionItem(collectionName, uuidKey string, models ...interface{}) (ops []txn.Op) {
	return buildTransactionArr("d-", collectionName, uuidKey, models)
}

// ConvertToUpdateTransactionItem ... This method manage the Mongo's transaction items from the current orm for update elements that currently exist in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToUpdateTransactionItem(collectionName, uuidKey string, models ...interface{}) (ops []txn.Op) {
	return buildTransactionArr("d+", collectionName, uuidKey, models)
}

// RunTransaction ... Perform a transaction over the DB with the options element.
func RunTransaction(collectionName string, ops []txn.Op) {

	session, _ := db.GetSession()

	defer func() {
		session.Close()
	}()
	c := db.Cursor(session, TransactionCollection)
	runner := txn.NewRunner(c)

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

func buildTransactionArr(assertType, collectionName, uuidKey string, models []interface{}) (ops []txn.Op) {
	for _, model := range models {
		uuid := ""
		if assertType == "d+" {
			var modelMap map[string]interface{}
			formatdata.FillStructP(model, &modelMap)

			if uuidKey == "" {
				uuidKey = "_id"
			}
			uuid = modelMap[uuidKey].(string)
		}

		ops = append(ops, buildTransactionItem(assertType, collectionName, uuid, model))
	}
	return
}

// GetTrStructIds ... Returns the IDs of a txn.Op struct.
func GetTrStructIds(trValues []txn.Op) (idsArray []string) {
	for _, tr := range trValues {
		idsArray = append(idsArray, tr.Id.(string))
	}
	return
}
