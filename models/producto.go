package models

import (
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// ProductosCollection es el nombre de la colección en mongo.
const ProductosCollection = "productos"

// Producto ...
type Producto struct {
	*General
	ID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
}

// InsertProducto registra un producto en la bd
func InsertProducto(j Producto) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, ProductosCollection)
	return c.Insert(j)
}

// GetAllProducto Obtiene todos los producstos registrados
func GetAllProducto(query map[string]interface{}) []Producto {
	var productos []Producto
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, ProductosCollection)
	if err = c.Find(query).All(&productos); err != nil {
		return nil
	}
	return productos
}

// GetProductoById obtiene un producto por su _id
func GetProductoById(id string) (Producto, error) {
	var producto Producto
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, ProductosCollection)
	err = c.FindId(bson.ObjectIdHex(id)).One(&producto)
	return producto, err
}

// DeleteProductoById elimina un producto por su _id
func DeleteProductoById(id string) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, ProductosCollection)
	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// UpdateProducto actualiza un prodcuto
func UpdateProducto(j Producto, id string) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, ProductosCollection)
	return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
}
