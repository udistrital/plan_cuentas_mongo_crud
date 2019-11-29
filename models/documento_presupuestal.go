package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// DocumentoPresupuestalCollection ... Nombre de la collección
const DocumentoPresupuestalCollection = "documento_presupuestal"

// DocumentoPresupuestal ... estructura para guardar información de documentos presupuestales.
type DocumentoPresupuestal struct {
	ID            string       `json:"_id" bson:"_id,omitempty"`
	Data          interface{}  `json:"Data" bson:"data" validate:"required"`
	Tipo          string       `json:"Tipo" bson:"tipo" validate:"required"`
	AfectacionIds []string     `json:"AfectacionIds" bson:"afectacion_ids"`
	Afectacion    []Movimiento `bson:"-" validate:"required"`
	FechaRegistro string       `json:"FechaRegistro" bson:"fecha_registro" validate:"required"`
	Estado        string       `json:"Estado" bson:"estado"`
	ValorActual   float64      `json:"ValorActual" bson:"valor_actual"`
	ValorInicial  float64      `json:"ValorInicial" bson:"valor_inicial"`
	Vigencia      int          `json:"Vigencia" bson:"-" validate:"required"`
	CentroGestor  string       `json:"CentroGestor" bson:"centro_gestor" validate:"required"`
	Consecutivo   int          `json:"Consecutivo" bson:"consecutivo"`
}

// GetDocumentoPresupuestalByDataID Obtiene la información de un documento presupuestal a través del atributo _id dentro de data
func GetDocumentoPresupuestalByDataID(id, vigencia, centroGestor string) (documentoPresupuestal DocumentoPresupuestal, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	defer session.Close()

	collectionFixed := DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor

	c := db.Cursor(session, collectionFixed)
	err = c.Find(bson.M{"data._id": id}).One(&documentoPresupuestal)
	return
}

// GetAllDocumentoPresupuestal ...
func GetAllDocumentoPresupuestal(vigencia, centroGestor string, query map[string]interface{}) (docPresupuestales []DocumentoPresupuestal, err error) {
	session, err := db.GetSession()
	if err != nil {
		return
	}
	defer session.Close()

	collectionFixed := DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor
	c := db.Cursor(session, collectionFixed)
	err = c.Find(query).All(&docPresupuestales)
	return
}
