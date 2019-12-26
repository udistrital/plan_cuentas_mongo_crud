package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// VigenciaCollectionName nombre de la colleccion para guardar las agrupaciones de vigencias.
const VigenciaCollection = "vigencia_233_"

// Vigencia estructura para acceder de forma mas rápida a la información de las vigencias registradas.
type Vigencia struct {
	ID                            string    `json:"Id" bson:"_id,omitempty"`
	Valor                         int       `json:"Valor" bson:"valor"`
	VigenciaEjecucionProgramacion string    `json:"Vigencia_ejecucion_programacion" bson:"vigenciaEjecucionProgramacion"`
	Activo                        bool      `json:"Activo" bson:"activo"`
	Estado                        string    `json:"Estado" bson:"estado"`
	FechaCreacion                 time.Time `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaModificacion             time.Time `json:"fechaModificacion" bson:"fechaModificacion"`
	FechaCierre                   time.Time `json:"fechaCierre" bson:"fechaCierre"`
}

type VigenciaNueva struct {
	Valor         int    `json:"Valor" bson:"valor"`
	AreaFuncional string `json:"AreaFuncional" bson:"areaFuncional"`
}

// InsertVigencia registra una Vigencia en la bd
func InsertVigencia(j *Vigencia, area string) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, VigenciaCollection+area)
	return c.Insert(j)
}

// GetAllVigencia Obtiene todos las Vigencia registradas
func GetAllVigencia(query map[string]interface{}, area string) []Vigencia {
	var Vigencia []Vigencia
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, VigenciaCollection+area)
	if err = c.Find(query).All(&Vigencia); err != nil {
		return nil
	}
	return Vigencia
}

// GetVigenciaByID obtiene una Vigencia por su _id
func GetVigenciaByID(id string, area string) (Vigencia, error) {
	var Vigencia Vigencia
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, VigenciaCollection+area)
	err = c.FindId(id).One(&Vigencia)
	return Vigencia, err
}

// UpdateVigencia actualiza una Vigencia
func UpdateVigencia(j *Vigencia, id string, area string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, VigenciaCollection+area)

	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}

// DeleteVigencia elimina una Vigencia con su ID
func DeleteVigencia(id string, area string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, VigenciaCollection+area)
	defer session.Close()

	return c.Remove(bson.M{"_id": id})
}
