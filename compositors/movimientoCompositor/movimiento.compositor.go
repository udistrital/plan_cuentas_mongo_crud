package movimientoCompositor

import (
	"github.com/globalsign/mgo/txn"
	movimientohelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// DocumentoPresupuestalRegister ... Add a new DocumentoPresuuestal document and it's propagation.
func DocumentoPresupuestalRegister(documentoPresupuestalRequestData *models.DocumentoPresupuestal) {
	var (
		movimientoData           []txn.Op
		movimientoDataInserted   []txn.Op
		valorActualDocumentoPres float64
	)
	initialState := "expedido"
	balanceMap := make(map[string]models.DocumentoPresupuestal)

	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		valorActualDocumentoPres += movimientoElmnt.ValorInicial
	}

	documentoPresupuestalRequestData.ValorActual = valorActualDocumentoPres
	documentoPresupuestalRequestData.ValorInicial = valorActualDocumentoPres
	documentoPresupuestalRequestData.Estado = initialState
	documentoPresupuestalOpStruct := transactionManager.ConvertToTransactionItem(models.DocumentoPresupuestalCollection, "", documentoPresupuestalRequestData)
	documentoPresupuestalRequestData.ID = documentoPresupuestalOpStruct[0].Id.(string)
	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		movimientoElmnt.DocumentoPresupuestalUUID = documentoPresupuestalOpStruct[0].Id.(string)
		movimientoElmnt.ValorActual = movimientoElmnt.ValorInicial
		movimientoElmnt.Estado = initialState
		insertMovimientoData := transactionManager.ConvertToTransactionItem(models.MovimientosCollection, "", movimientoElmnt)
		movimientoDataInserted = append(movimientoDataInserted, insertMovimientoData...)
		movimientoData = append(movimientoData, insertMovimientoData...)
		propagacionData := movimientohelper.BuildPropagacionValoresTr(movimientoElmnt, balanceMap)

		if len(propagacionData) > 0 {
			movimientoData = append(movimientoData, propagacionData...)
		}

		valorActualDocumentoPres += movimientoElmnt.ValorInicial

	}

	documentoPresupuestalRequestData.AfectacionIds = transactionManager.GetTrStructIds(movimientoDataInserted)
	movimientoData = append(movimientoData, documentoPresupuestalOpStruct...)
	// Perform Mongo's Transaction.
	transactionManager.RunTransaction(models.MovimientosCollection, movimientoData)
	updateAfectationData := transactionManager.ConvertToUpdateTransactionItem(models.DocumentoPresupuestalCollection, "", *documentoPresupuestalRequestData)
	transactionManager.RunTransaction(models.DocumentoPresupuestalCollection, updateAfectationData)
}

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
