package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
)

// VigenciaCollectionName nombre de la colleccion para guardar las agrupaciones de vigencias.
const VigenciaCollectionName = "vigencia_233_"

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

type VigenciaNueva struct {
	Valor         int    `json:"Valor" bson:"valor"`
	AreaFuncional string `json:"AreaFuncional" bson:"areaFuncional"`
}

//UpdateVigencia ... actializa una vigencia
func UpdateVigencia(j *Vigencia, id string, areaFuncional string) error {
	session, err := db.GetSession()
	if err != nil {
		return err
	}
	c := db.Cursor(session, VigenciaCollectionName+areaFuncional)

	defer session.Close()

	return c.Update(bson.M{"_id": id}, &j)
}
