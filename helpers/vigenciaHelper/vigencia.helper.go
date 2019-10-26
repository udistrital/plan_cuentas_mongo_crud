package vigenciahelper

import (
	"strconv"

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
	}
	return crudmanager.AddNew(models.VigenciaCollectionName, vigenciaStruct)
}
