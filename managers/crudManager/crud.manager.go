package crudmanager

import (
	"log"

	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
)

// GetAllFromDB ... get an array data from db by a collection and query.
func GetAllFromDB(query map[string]interface{}, collectionName string, outStruct interface{}) {
	var collectionData []interface{}
	var resulData []interface{}
	session, err := db.GetSession()
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := db.Cursor(session, collectionName)
	defer session.Close()
	err = c.Find(query).All(&collectionData)
	if err != nil {
		log.Println(err.Error())
	}
	for _, partialData := range collectionData {
		var data interface{}
		commonhelper.FillStructBson(partialData, &data)
		resulData = append(resulData, data)
	}
	commonhelper.FillArrBson(resulData, outStruct)
}

// GetDocumentByID ... get a document values by it's id. Returns a map with bson tags basis struct.
func GetDocumentByID(uuid, collectionName string) (interface{}, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, err
	}
	c := db.Cursor(session, collectionName)
	var documentData interface{}
	err = c.FindId(uuid).One(&documentData)
	defer session.Close()
	return documentData, err
}
