package movimientoCompositor

import (
	"strconv"

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
	balanceMap := make(map[string]map[string]interface{})
	collectionPostFixName := "_" + strconv.Itoa(documentoPresupuestalRequestData.Vigencia) + "_" + documentoPresupuestalRequestData.CentroGestor

	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		valorActualDocumentoPres += movimientoElmnt.ValorInicial
	}

	documentoPresupuestalRequestData.ValorInicial = valorActualDocumentoPres
	documentoPresupuestalRequestData.ValorActual = valorActualDocumentoPres
	documentoPresupuestalRequestData.Estado = initialState
	documentoPresupuestalOpStruct := transactionManager.ConvertToTransactionItem(models.DocumentoPresupuestalCollection+collectionPostFixName, "", "Afectacion", documentoPresupuestalRequestData)
	documentoPresupuestalRequestData.ID = documentoPresupuestalOpStruct[0].Id.(string)
	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		movimientoElmnt.DocumentoPresupuestalUUID = documentoPresupuestalOpStruct[0].Id.(string)
		movimientoElmnt.ValorActual = movimientoElmnt.ValorInicial
		movimientoElmnt.Estado = initialState
		insertMovimientoData := transactionManager.ConvertToTransactionItem(models.MovimientosCollection+collectionPostFixName, "", "", movimientoElmnt)
		movimientoDataInserted = append(movimientoDataInserted, insertMovimientoData...)
		movimientoData = append(movimientoData, insertMovimientoData...)
		propagacionData := movimientohelper.BuildPropagacionValoresTr(movimientoElmnt, balanceMap, collectionPostFixName)

		if len(propagacionData) > 0 {
			movimientoData = append(movimientoData, propagacionData...)
		}

		valorActualDocumentoPres += movimientoElmnt.ValorInicial

	}

	documentoPresupuestalRequestData.AfectacionIds = transactionManager.GetTrStructIds(movimientoDataInserted)
	movimientoData = append(movimientoData, documentoPresupuestalOpStruct...)
	// Perform Mongo's Transaction.
	transactionManager.RunTransaction(movimientoData)
	updateAfectationData := transactionManager.ConvertToUpdateTransactionItem(models.DocumentoPresupuestalCollection+collectionPostFixName, "", "afectacion_ids", *documentoPresupuestalRequestData)
	transactionManager.RunTransaction(updateAfectationData)
}

// AddMovimientoTransaction ... Add Movimiento's document to mongo db and it's afectation
// over the apropiation's tree.
func AddMovimientoTransaction(movimientoData ...models.Movimiento) []interface{} {
	var (
		ops []interface{}
	)

	for _, element := range movimientoData {
		opMov := transactionManager.ConvertToTransactionItem(models.MovimientosCollection, "", "", element)
		ops = append(ops, opMov)
	}

	return ops
}
