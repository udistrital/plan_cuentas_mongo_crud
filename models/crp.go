package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// CRPCollection es el nombre de la colección en mongo.
const SolicitudCRPCollection = "crp"

// infoCrp asociado a una solicitud de CRP
type infoCrp struct {
	Consecutivo     int       `json:"consecutivo" bson:"consecutivo"`
	FechaExpedicion time.Time `json:"fechaExpedicion" bson:"fechaExpedicion"`
	Estado          int       `json:"estado" bson:"estado"`
}

// Compromiso asociado a una solicitud de CRP
type Compromiso struct {
	NumeroCompromiso int `json:"numeroCompromiso" bson:"numeroCompromiso"`
	TipoCompromiso   int `json:"tipoCompromiso" bson:"tipoCompromiso"`
}

// SolicitudCRP información de la solicitud de un CRP
type SolicitudCRP struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Consecutivo    int           `json:"consecutivo" bson:"consecutivo"`
	ConsecutivoCDP int           `json:"consecutivoCdp" bson:"consecutivoCdp"`
	Beneficiario   string        `json:"beneficiario" bson:"beneficiario"`
	Valor          float64       `json:"valor" bson:"valor"`
	Compromiso     *Compromiso   `json:"compromiso" bson:"compromiso"`
	InfoCRP        *infoCrp      `json:"infoCrp" bson:"infoCrp"`
}

// InsertSolicitudCRP registra una SolicitudCRP en la bd
func InsertSolicitudCRP(j *SolicitudCRP) error {
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCRPCollection)
	return c.Insert(j)
}

// GetAllSolicitudCRP Obtiene todos las SolicitudCRP registradas
func GetAllSolicitudCRP(query map[string]interface{}) []SolicitudCRP {
	var SolicitudCRP []SolicitudCRP
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCRPCollection)
	if err = c.Find(query).All(&SolicitudCRP); err != nil {
		return nil
	}
	return SolicitudCRP
}

// GetSolicitudCRPByID obtiene una SolicitudCRP por su _id
func GetSolicitudCRPByID(id string) (SolicitudCRP, error) {
	var SolicitudCRP SolicitudCRP
	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}
	defer session.Close()

	c := db.Cursor(session, SolicitudCRPCollection)
	err = c.FindId(bson.ObjectIdHex(id)).One(&SolicitudCRP)
	return SolicitudCRP, err
}

// UpdateSolicitudCRP actualiza una SolicitudCRP
func UpdateSolicitudCRP(j *SolicitudCRP, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, SolicitudCRPCollection)

	defer session.Close()

	return c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
}

// DeleteSolicitudCRP elimina una SolicitudCRP con su ID
func DeleteSolicitudCRP(id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}

	c := db.Cursor(session, SolicitudCRPCollection)
	defer session.Close()

	return c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
