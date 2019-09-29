package crudmanager

import (
	"log"

	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// GetAllFromDB ... get an array data from db by a collection and query.
func GetAllFromDB(query map[string]interface{}, collectionName string) (collectionData interface{}) {
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

	return collectionData
}

// GetDocumentByID ... get a document values by it's id.
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
