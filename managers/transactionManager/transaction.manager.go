package transactionManager

import (
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// ConvertToTransactionItem ... This method manage the Mongo's transaction items from the current orm for registration of new elements in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToTransactionItem(collectionName, uuidKey, fieldsToIgnore string, models ...interface{}) (ops []txn.Op) {
	return buildTransactionArr("d-", collectionName, uuidKey, models, "", fieldsToIgnore)
}

// ConvertToUpdateTransactionItem ... This method manage the Mongo's transaction items from the current orm for update elements that currently exist in the DB , you should pass
// the collection name and the model structure. This method will return a transaction elemnt.
func ConvertToUpdateTransactionItem(collectionName, uuidKey, fields string, models ...interface{}) (ops []txn.Op) {
	return buildTransactionArr("d+", collectionName, uuidKey, models, fields, "")
}

// RunTransaction ... Perform a transaction over the DB with the options element.
func RunTransaction(ops []txn.Op) {

	session, _ := db.GetSession()

	defer func() {
		session.Close()
	}()
	c := db.Cursor(session, models.TransactionCollection)
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

func buildTransactionArr(assertType, collectionName, uuidKey string, models []interface{}, filedsToUpdate, fieldsToIgnore string) (ops []txn.Op) {
	for _, model := range models {
		modelMap, err := commonhelper.ToMap(model, "bson")
		if err != nil {
			panic(err.Error())
		}

		uuid := ""
		if uuidKey == "" {
			uuidKey = "_id"
		}
		if assertType == "d+" {
			if filedsToUpdate != "" {
				filedsArr := strings.Split(filedsToUpdate, ",")
				for field := range modelMap {
					var deleteField = true
					for _, fieldToUpdate := range filedsArr {
						if fieldToUpdate == field || field == uuidKey {
							deleteField = false
						}
					}
					if deleteField {
						if modelMap[field] != nil {
							delete(modelMap, field)
						}
					}
				}
			}

			uuid = modelMap[uuidKey].(string)
		} else {
			filedsArr := strings.Split(fieldsToIgnore, ",")
			for field := range modelMap {
				for _, fieldsToIgnore := range filedsArr {
					if fieldsToIgnore == field && field != uuidKey {
						delete(modelMap, field)
					}
				}
			}
		}

		ops = append(ops, buildTransactionItem(assertType, collectionName, uuid, modelMap))
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
