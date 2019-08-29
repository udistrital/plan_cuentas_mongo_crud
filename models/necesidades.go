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
	Codigo string  `json:"codigo" bson:"codigo"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Metas de una necesidad
type meta struct {
	Codigo 		string 		 `json:"codigo" bson:"codigo"`
	Actividades []*actividad `json:"actividades" bson:"actividades"`
	Valor       float64      `json:"valor" bson:"valor"`
}

// Rubro de la necesidad (es el que va a tener las metas)
type rubro struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Metas  []*meta `json:"metas" bson:"metas"`
}

// Fuentes de la necesidad 
type fuente struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Productos de la necesidad 
type producto struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Productos de la necesidad 
type detalleServicio struct {
	Codigo 			string  `json:"codigo" bson:"codigo"`
	Valor  			float64 `json:"valor" bson:"valor"`
	Descripcion 	string  `json:"descripcion" bson:"descripcion"`
}

// Necesidad información de la necesidad
type Necesidad struct {
	ID               bson.ObjectId 		`json:"_id" bson:"_id,omitempty"`
	IDAdministrativa int           		`json:"idAdministrativa" bson:"idAdministrativa"`
	Valor            float64       		`json:"valor" bson:"valor"`
	Rubros           *rubro        		`json:"rubros" bson:"rubros"`
	Fuentes			 []*fuente	   		`json:"fuentes" bson:"fuentes"`
	Productos		 []*producto   		`json:"productos" bson:"productos"`
	DetalleServicio	 *detalleServicio	`json:"detalleServicio" bson:"detalleServicio"`
	TipoContrato	 int		   		`json:"tipoContrato" bson:"tipoContrato"`
}

// NecesidadCollection constante para la colección
const NecesidadCollection = "necesidades"

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

// UpdateNecesidad actualiza una necesidad
func UpdateNecesidad(j *Necesidad, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, NecesidadesCollection)

	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}

// DeleteNecesidad elimina una necesidad con su ID
func DeleteNecesidad(id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, NecesidadCollection)
	defer session.Close()

	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}