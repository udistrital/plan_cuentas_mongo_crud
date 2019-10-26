package crudmanager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
)

// GetAllFromDB ... get an array data from db by a collection and query.
func GetAllFromDB(query map[string]interface{}, collectionName string, outStruct interface{}) {
	var collectionData []interface{}
	var resulData []interface{}
	session, c, err := GetDBCursorByCollection(collectionName)
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
	session, c, err := GetDBCursorByCollection(collectionName)

	if err != nil {
		return nil, err
	}
	var documentData interface{}
	err = c.FindId(uuid).One(&documentData)
	defer session.Close()
	return documentData, err
}

// RunPipe runs pipe of mongo's aggregation func.
func RunPipe(collectionName string, queries ...bson.M) ([]interface{}, error) {
	session, c, err := GetDBCursorByCollection(collectionName)

	if err != nil {
		return nil, err
	}

	defer session.Close()
	var result []interface{}
	pipeline := c.Pipe(queries)
	pipeline.All(&result)

	return result, err
}

// AddNew this function will add new data to a specific collection.
func AddNew(collectionName string, data ...interface{}) error {
	session, c, err := GetDBCursorByCollection(collectionName)
	if err != nil {
		return err
	}
	defer session.Close()
	return c.Insert(data...)
}

// GetDBCursorByCollection Return a mgo session and cursor by it's name.
func GetDBCursorByCollection(collectionName string) (*mgo.Session, *mgo.Collection, error) {
	session, err := db.GetSession()
	if err != nil {
		return nil, nil, err
	}
	c := db.Cursor(session, collectionName)
	return session, c, err
}
