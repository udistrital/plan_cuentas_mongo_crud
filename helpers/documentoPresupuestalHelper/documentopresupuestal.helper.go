package documentopresupuestalhelper

import (
	"fmt"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

func JoinDocumentoPresupuestalMovs(documentoPresupuestalData *models.DocumentoPresupuestal, movimientoCollector map[string]interface{}) {
	var movimientosArr []models.Movimiento

	for _, id := range documentoPresupuestalData.AfectacionIds {
		if movimientoCollector[id] != nil {
			var movimiento models.Movimiento
			if err := formatdata.FillStruct(movimientoCollector[id], &movimiento); err == nil {
				movimientosArr = append(movimientosArr, movimiento)
			} else {
				fmt.Println("cannot convert", err.Error())
			}
		}
	}

	if len(movimientosArr) > 0 {
		documentoPresupuestalData.Afectacion = movimientosArr
	}
}
