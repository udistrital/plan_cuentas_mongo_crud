package models

import (
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// NecesidadesCollection es el nombre de la colección en mongo.
const NecesidadesCollection = "necesidades"

// Actividades asociadas a una meta
type actividad struct {
	Codigo string `json:"codigo" bson:"codigo"`
	Nombre string `json:"nombre" bson:"nombre"`
}

// Metas de una necesidad
type meta struct {
	Nombre      string       `json:"nombre" bson:"nombre"`
	Descripcion string       `json:"descripcion" bson:"descripcion"`
	Actividades []*actividad `json:"actividades" bson:"actividades"`
}

// Rubro de la necesidad (es el que va a tener las metas)
type rubro struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Metas  []*meta `json:"metas" bson:"metas"`
}

// Necesidad información de la necesidad
type Necesidad struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Descripcion string        `json:"descripcion" bson:"descripcion"`
	Rubros      *rubro        `json:"rubros" bson:"rubros"`
}

// InsertNecesidad registra una necesidad en la bd
func InsertNecesidad(j *Necesidad) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, NecesidadesCollection)
	return c.Insert(j)
}

// GetAllNecesidad Obtiene todos las necesidades registradas
func GetAllNecesidad(query map[string]interface{}) []Necesidad {
	var necesidades []Necesidad
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, NecesidadesCollection)
	if err = c.Find(query).All(&necesidades); err != nil {
		return nil
	}
	return necesidades
}
