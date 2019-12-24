package documentopresupuestalmanager

import (
	"fmt"
	"log"

	"strconv"

	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func GetByType(vigencia, centroGestor, tipo string) []models.DocumentoPresupuestal {
	query := make(map[string]interface{})
	collectionFixedName := models.DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor
	var documentoPresupuestalRows []models.DocumentoPresupuestal
	query["tipo"] = map[string]interface{}{
		"$regex": "^" + tipo + ".*?$",
	}

	fmt.Println(query)

	crudmanager.GetAllFromDB(query, collectionFixedName, &documentoPresupuestalRows, "-fecha_registro")
	return documentoPresupuestalRows
}

func GetByTypeLike(vigencia, centroGestor, tipo string) []models.DocumentoPresupuestal {
	query := make(map[string]interface{})
	collectionFixedName := models.DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor
	var documentoPresupuestalRows []models.DocumentoPresupuestal
	query["tipo"] = map[string]interface{}{
		"$regex": ".*" + tipo + ".*",
	}

	fmt.Println(query)

	crudmanager.GetAllFromDB(query, collectionFixedName, &documentoPresupuestalRows, "-fecha_registro")
	return documentoPresupuestalRows
}

func GetOneByType(UUID, vigencia, centroGestor, tipo string) (models.DocumentoPresupuestal, error) {
	query := make(map[string]interface{})
	collectionFixedName := models.DocumentoPresupuestalCollection + "_" + vigencia + "_" + centroGestor
	var documentoPresupuestalRows []models.DocumentoPresupuestal
	query["_id"] = UUID

	err := crudmanager.GetAllFromDB(query, collectionFixedName, &documentoPresupuestalRows)
	if err != nil {
		return models.DocumentoPresupuestal{}, err
	}
	return documentoPresupuestalRows[0], nil
}

// GetCDPByID obtiene una un CDP expedido con su _id de solicitud
func GetCDPByID(id string) (documentoPresupuestal models.DocumentoPresupuestal) {
	solicitudCdp, err := models.GetSolicitudCDPByID(id)

	if err != nil {
		log.Panicln(err.Error())
		return
	}

	documentoPresupuestal, err = models.GetDocumentoPresupuestalByDataID(id, solicitudCdp.Vigencia, strconv.Itoa(solicitudCdp.CentroGestor))
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	return
}
