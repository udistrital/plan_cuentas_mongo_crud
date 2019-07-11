package movimientoCompositor

import (
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// AddMovimientoTransaction ... Add Movimiento's document to mongo db and it's afectation
// over the apropiation's tree.
func AddMovimientoTransaction(movimientoData ...models.Movimiento) {
	var (
		ops []interface{}
		// month string
		// year  string
	)

	// layout := "2006-01-02T15:04:05.000Z"
	// t, err := time.Parse(layout, movimientoData.FechaRegistro)

	// if err != nil {
	// 	logManager.LogError(err.Error())
	// 	panic(err.Error())
	// }

	// month = strconv.Itoa(int(t.Month()))
	// year = strconv.Itoa(int(t.Year()))

	for _, element := range movimientoData {
		opMov := transactionManager.ConvertToTransactionItem(models.MovimientosCollection, element)
		ops = append(ops, opMov)
	}

	// afectacionMap := make(map[string]map[string]float64)

	// for _, afect := range movimientoData.Afectacion {

	// 	if afectacionMap[afect["Codigo"].(string)] == nil {
	// 		afectacionMap[afect["Codigo"].(string)] = make(map[string]float64)
	// 	}

	// 	afectacionMap[afect["Codigo"].(string)][afect["Tipo"].(string)] += afect["Valor"].(float64)
	// }

	// for k, v := range afectacionMap {
	// 	op, err := apropiacionHelper.PropagacionValores(k, month, year, strconv.Itoa(movimientoData.UnidadEjecutora), v)
	// 	if err != nil {
	// 		logManager.LogError(err.Error())
	// 		panic(err.Error())
	// 	}
	// 	ops = append(ops, op...)
	// }

	transactionManager.RunTransaction(models.MovimientosCollection, ops)

}
