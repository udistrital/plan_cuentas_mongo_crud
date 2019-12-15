package models

import (
	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// VigenciaCollectionName nombre de la colleccion para guardar las agrupaciones de vigencias.
const VigenciaCollectionName = "vigencia"

// Vigencia estructura para acceder de forma mas rápida a la información de las vigencias registradas.
type Vigencia struct {
	ID            string `json:"Id" bson:"_id,omitempty"`
	NameSpace     string `json:"NameSpace" bson:"name_space"`
	CentroGestor  string `json:"CentroGestor" bson:"centro_gestor"`
	AreaFuncional string `json:"AreaFuncional" bson:"area_funcional"`
	Valor         int    `json:"Valor" bson:"valor"`
	Estado        string `json:"Estado" bson:"estado"`
}

//UpdateVigencia ... actializa una vigencia
func UpdateVigencia(j *Vigencia, id string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, VigenciaCollectionName)

	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}
