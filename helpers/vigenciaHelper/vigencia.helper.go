package vigenciahelper

import (
	"strconv"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/globalsign/mgo/bson"

	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// AddNew adds a new record to vigencia collection.
func AddNew(value int, namespace string, cg string) (err error) {
	vigenciaStruct := models.Vigencia{
		NameSapce:    namespace,
		Valor:        value,
		ID:           namespace + "_" + strconv.Itoa(value),
		CentroGestor: cg,
		Estado:       models.EstadoRegistrada,
	}
	return crudmanager.AddNew(models.VigenciaCollectionName, vigenciaStruct)
}

// GetVigenciasByNameSpaceAndCg return an array with numeric values of collection "vigencia" by namespace and cg.
func GetVigenciasByNameSpaceAndCg(namespace, cg string) (vigenciasArr []map[string]int, err error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"name_space": namespace, "centro_gestor": cg},
		}, bson.M{
			"$group": bson.M{
				"_id": "$valor",
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id": -1,
			},
		},
	}
	var unformatedVigenciaArr []interface{}
	if unformatedVigenciaArr, err = crudmanager.RunPipe(models.VigenciaCollectionName, pipeline...); err == nil {
		for _, unformatedVigencia := range unformatedVigenciaArr {
			var unformatedVigenciaMap map[string]interface{}
			formatdata.FillStructP(unformatedVigencia, &unformatedVigenciaMap)
			vigenciasArr = append(vigenciasArr, map[string]int{"vigencia": int(unformatedVigenciaMap["_id"].(float64))})
		}
	}

	return

}
