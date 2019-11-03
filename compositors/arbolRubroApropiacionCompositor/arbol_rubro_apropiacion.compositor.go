package arbolrubroapropiacioncompositor

import (
	"strconv"

	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// CheckForAprTreeBalanceWithSimulation this function will return a balance and approved boolean flags
func CheckForAprTreeBalanceWithSimulation(movimientos []models.Movimiento, vigenciaStr, ueStr string) (balanceado bool, approved bool, response map[string]interface{}, err error) {

	vigencia, _ := strconv.Atoi(vigenciaStr)
	raices, err := models.GetRaicesApropiacion(ueStr, vigencia)

	if err != nil {
		return
	}

	balance := make(map[string]map[string]interface{})
	rootsAprpovedTotal := 0
	if len(movimientos) > 0 {
		balance = rubroApropiacionHelper.SimulatePropagationValues(movimientos, vigenciaStr, ueStr)
	}

	var rootCompValue float64
	values := make(map[string]models.NodoRubroApropiacion)
	balanceado = true
	approved = false
	rootsParamsIndexed := rubroHelper.GetRubroParamsIndexedByKey(ueStr, "Valor")

	for _, raiz := range raices {
		if rootsParamsIndexed[raiz.ID] != nil {
			values[raiz.ID] = raiz
			if raiz.Estado == models.EstadoAprobada {
				rootsAprpovedTotal++
			}
		}
		// perform this operation only if there are some simulation to perform ...
		if rootsParamsIndexed[raiz.ID] != nil && balance[raiz.ID] != nil {
			actualRaiz := raiz
			actualRaiz.ValorActual = balance[raiz.ID]["valor_actual"].(float64)
			values[raiz.ID] = actualRaiz
		}
	}
	var indexValue int
	response = make(map[string]interface{})
	for _, rootValue := range values {
		if indexValue == 0 {
			rootCompValue = rootValue.ValorActual
		}
		if rootCompValue != rootValue.ValorActual || rootValue.ValorActual == 0 {
			balanceado = false
		}
		if rubroInfo, e := rubroManager.SearchRubro(rootValue.ID, ueStr); e {
			if rubroInfo.ID == "2" {
				response["totalIngresos"] = rootValue.ValorActual
			} else if rubroInfo.ID == "3" {
				response["totalGastos"] = rootValue.ValorActual
			}

		}
		indexValue++
	}

	if rootCompValue == 0 {
		balanceado = false
	}

	if response["totalGastos"] == nil || response["totalIngresos"] == nil {
		balanceado = false
	}

	// if the tree's roots are all approved then the whole tree is approved..
	if rootsAprpovedTotal == len(rootsParamsIndexed) {
		approved = true
	}

	response["balanceado"] = balanceado
	response["approved"] = approved

	return
}
