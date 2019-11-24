package documentopresupuestalmanager

import (
	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func GetByType(vigencia, centroGestor, tipo string) []models.DocumentoPresupuestal {
	query := make(map[string]interface{})
	collectionFixedName := models.DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor
	var documentoPresupuestalRows []models.DocumentoPresupuestal
	query["tipo"] = map[string]interface{}{
		"$regex": ".*" + tipo + ".*",
	}
	crudmanager.GetAllFromDB(query, collectionFixedName, &documentoPresupuestalRows)
	return documentoPresupuestalRows
}
