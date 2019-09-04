package movimientoCompositor

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/txn"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// AddMovimientoTransaction ... Add Movimiento's document to mongo db and it's afectation
// over the apropiation's tree.
func AddMovimientoTransaction(movimientoData ...models.Movimiento) []interface{} {
	var (
		ops []interface{}
	)

	for _, element := range movimientoData {
		opMov := transactionManager.ConvertToTransactionItem(models.MovimientosCollection, "", element)
		ops = append(ops, opMov)
	}

	return ops
}

// BuildPropagacionValoresTr ... Build a mgo transaction item as Array of interfaces .
// This method search in "movimientos_parametros" collection for the afectation's config recursively.
func BuildPropagacionValoresTr(movimiento models.Movimiento) (trData []txn.Op) {
	movimientoParameter, err := movimientoManager.GetOneMovimientoParameterByHijo(movimiento.Tipo)
	var arrMovimientosUpdted []interface{}
	var runFlag = true

	if err != nil {
		logs.Error("1", err)
		return
	}

	movimientoPadre, err := movimientoManager.GetOneMovimientoByTipo(movimiento.DocumentoPadre, movimientoParameter.TipoMovimientoPadre)
	movimientoHijo := movimiento

	if err != nil {
		runFlag = false
	}
	fmt.Println("Documento padre: ", movimientoPadre)
	fmt.Println("Movimiento parameter: ", movimientoParameter)
	for runFlag {

		if len(movimientoPadre.Movimientos) == 0 {
			movimientoPadre.Movimientos = make(map[string]float64)
		}

		if movimientoPadre.Movimientos[movimientoHijo.Tipo] == 0 {
			movimientoPadre.Movimientos[movimientoHijo.Tipo] = movimientoHijo.Valor * float64(movimientoParameter.Multiplicador)
			fmt.Println("debio irse por aquí....")
		} else {
			movimientoPadre.Movimientos[movimientoHijo.Tipo] += (movimientoHijo.Valor * float64(movimientoParameter.Multiplicador))
		}

		arrMovimientosUpdted = append(arrMovimientosUpdted, movimientoPadre)

		movimientoHijo = movimientoPadre
		movimientoParameter, err := movimientoManager.GetOneMovimientoParameterByHijo(movimientoHijo.Tipo)

		if err != nil {
			logs.Error("2", err)
			panic(err)
		}
		movimientoPadre.Movimientos = make(map[string]float64)
		movimientoPadre, err = movimientoManager.GetOneMovimientoByTipo(movimientoHijo.DocumentoPadre, movimientoParameter.TipoMovimientoPadre)
		if err != nil {
			runFlag = false
		}
	}

	trData = transactionManager.ConvertToUpdateTransactionItem(models.MovimientosCollection, "", arrMovimientosUpdted...)

	return
}
