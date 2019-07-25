package models

import (
	"github.com/globalsign/mgo"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// ProductosCollection es el nombre de la colección en mongo.
const ProductosCollection = "productos"

// Producto ...
type Producto struct {
	*General
	ID string `json:"_id" bson:"_id,omitempty"`
}

// InsertProducto registra un producto en la bd
func InsertProducto(session *mgo.Session, j Producto) error {
	c := db.Cursor(session, ProductosCollection)
	defer session.Close()
	err := c.Insert(j)
	return err
}
