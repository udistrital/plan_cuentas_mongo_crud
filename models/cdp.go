package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// SolicitudCDPCollection es el nombre de la colección en mongo.
const SolicitudCDPCollection = "cdp"

// SolicitudCDP información de la solicitud de un CDP
type SolicitudCDP struct {
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Consecutivo          int           `json:"consecutivo" bson:"consecutivo"`
	Entidad              int           `json:"entidad" bson:"entidad"`
	CentroGestor         int           `json:"centroGestor" bson:"centroGestor"`
	Necesidad            int           `json:"necesidad" bson:"necesidad"`
	Vigencia             string        `json:"vigencia" bson:"vigencia"`
	FechaRegistro        time.Time     `json:"fechaRegistro" bson:"fechaRegistro"`
	Estado               interface{}   `json:"estado" bson:"estado"`
	JustificacionRechazo string        `json:"justificacionRechazo" bson:"justificacionRechazo"`
	Activo               bool          `json:"activo" bson:"activo"`
	FechaCreacion        time.Time     `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaModificacion    time.Time     `json:"fechaModificacion" bson:"fechaModificacion"`
}

// InsertSolicitudCDP registra una SolicitudCDP en la bd
func InsertSolicitudCDP(j *SolicitudCDP) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCDPCollection)
	return c.Insert(j)
}

// GetAllSolicitudCDP Obtiene todos las SolicitudCDP registradas
func GetAllSolicitudCDP(query map[string]interface{}) []SolicitudCDP {
	var SolicitudCDP []SolicitudCDP
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCDPCollection)
	if err = c.Find(query).All(&SolicitudCDP); err != nil {
		return nil
	}
	return SolicitudCDP
}

// GetSolicitudCDPByID obtiene una SolicitudCDP por su _id
func GetSolicitudCDPByID(id string) (SolicitudCDP, error) {
	var SolicitudCDP SolicitudCDP
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCDPCollection)
	err = c.FindId(bson.ObjectIdHex(id)).One(&SolicitudCDP)
	return SolicitudCDP, err
}

// UpdateSolicitudCDP actualiza una SolicitudCDP
func UpdateSolicitudCDP(j *SolicitudCDP, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, SolicitudCDPCollection)

	defer session.Close()

	return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
}

// DeleteSolicitudCDP elimina una SolicitudCDP con su ID
func DeleteSolicitudCDP(id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, SolicitudCDPCollection)
	defer session.Close()

	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
