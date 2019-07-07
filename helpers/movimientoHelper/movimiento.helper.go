package movimientoHelper

import (
	"github.com/udistrital/utils_oas/formatdata"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func FormatMovimientoRequestData(requestData map[string]interface{}, tipo string) (movimientoData models.Movimiento) {
	switch tipo {
	case "modificacion":
		return convertToModificacionStruct(requestData)
	default:
		panic("Tipo De Movimiento No Encontrado")
	}
}

func convertToModificacionStruct(data map[string]interface{}) (movimientoData models.Movimiento) {

	if err := formatdata.FillStruct(data, &movimientoData); err != nil {
		movimientoData.IDPsql = int(data["Id"].(float64))
	}

	return
}
