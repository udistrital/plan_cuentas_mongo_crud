package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// MovimientosCollection es el nombre de la colección en mongo.
const MovimientosCollection = "movimientos"

// Movimiento es una estructura generica para los tipos de movimiento registados.
type Movimiento struct {
	ID              string                   `json:"_id" bson:"_id,omitempty"`
	IDPsql          int                      `json:"idpsql"`
	RubrosAfecta    []map[string]interface{} `json:"rubros_afecta"`
	Tipo            string                   `json:"tipo"`
	DocumentoPadre  string                   `json:"documento_padre"`
	FechaRegistro   string                   `json:"fecha_registro"`
	UnidadEjecutora int                      `json:"unidad_ejecutora"`
}

// GetMovimientoByPsqlId Obtener un documento por el idpsql
func GetMovimientoByPsqlId(session *mgo.Session, id, tipo string) (*Movimiento, error) {
	c := db.Cursor(session, MovimientosCollection)
	defer session.Close()
	var Movimiento *Movimiento
	err := c.Find(bson.M{"idpsql": id, "tipo": tipo}).One(&Movimiento)
	return Movimiento, err
}
