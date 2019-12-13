package vigenciahelper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/udistrital/utils_oas/formatdata"

	"github.com/globalsign/mgo/bson"

	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

const VigenciaActual, VigenciaSiguiente, VigenciaCerrada = "actual", "siguiente", "cerrada"

// AddNew adds a new record to vigencia collection. Cada que se agregue una nueva vigencia
func AddNew(value int, namespace string, areafuncional string, cg string, estado string) error {

	vigenciaStruct := models.Vigencia{
		ID:            strconv.Itoa(value) + "_" + cg,
		NameSpace:     namespace,
		AreaFuncional: areafuncional,
		CentroGestor:  cg,
		Valor:         value,
		Estado:        estado,
	}
	if strings.ToLower(estado) == VigenciaActual {
		if consultarVigencia(vigenciaStruct) {
			if err := models.UpdateVigencia(&vigenciaStruct, vigenciaStruct.ID); err == nil {
				return err
			}
		}
	}
	if err := AgregarVigenciaSiguiente(value+1, namespace, areafuncional, cg); err != nil {
		return err
	}
	return crudmanager.AddNew(models.VigenciaCollectionName, vigenciaStruct)
}

//Agrega una nueva vigencia con el estado de siguiente, cuando se registra una viencia nueva
func AgregarVigenciaSiguiente(value int, namespace string, areafuncional, cg string) error {
	nuevaVigencia := models.Vigencia{
		ID:            strconv.Itoa(value) + "_" + cg,
		NameSpace:     namespace,
		AreaFuncional: areafuncional,
		CentroGestor:  cg,
		Valor:         value,
		Estado:        VigenciaSiguiente,
	}
	return crudmanager.AddNew(models.VigenciaCollectionName, nuevaVigencia)
}

// Consulta si ya existe una vigencia en la base de datos, devuelve true si existe, false en caso contrario
func consultarVigencia(vig models.Vigencia) (existe bool) {
	existe = false
	if vigencia, err := GetVigenciaById(vig.ID); err == nil {
		if len(vigencia) != 0 {
			existe = true
		}
	}
	fmt.Println(existe)
	return
}

// GetVigenciaByID obtiene y devuelve la vigenci con el id que se pasó por parametros.
func GetVigenciaById(id string) (vigencia []interface{}, err error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"_id": id},
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
	vigencia, err = crudmanager.RunPipe(models.VigenciaCollectionName, pipeline...)
	fmt.Println("Vigencia: ", vigencia, err)
	return
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
