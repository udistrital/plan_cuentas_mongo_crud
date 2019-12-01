package movimientoCompositor

import (
	"strconv"

	"github.com/globalsign/mgo/txn"
	movimientohelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/movimientoHelper"
	docMananger "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	movimientoManager "github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
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
	balance := make(map[string]map[string]interface{})
	afectationIndex := make(map[string]map[string]interface{})
	consecutivoIndex := make(map[string]int)
	vigencia := strconv.Itoa(documentoPresupuestalRequestData.Vigencia)
	centroGestor := documentoPresupuestalRequestData.CentroGestor
	collectionPostFixName := "_" + vigencia + "_" + centroGestor
	consecutivoDocumento := len(docMananger.GetByType(vigencia, centroGestor, documentoPresupuestalRequestData.Tipo)) + 1

	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		valorActualDocumentoPres += movimientoElmnt.ValorInicial
	}

	documentoPresupuestalRequestData.ValorInicial = valorActualDocumentoPres
	documentoPresupuestalRequestData.ValorActual = valorActualDocumentoPres
	documentoPresupuestalRequestData.Estado = initialState
	documentoPresupuestalRequestData.Consecutivo = consecutivoDocumento
	documentoPresupuestalOpStruct := transactionManager.ConvertToTransactionItem(models.DocumentoPresupuestalCollection+collectionPostFixName, "", "Afectacion", documentoPresupuestalRequestData)
	documentoPresupuestalRequestData.ID = documentoPresupuestalOpStruct[0].Id.(string)

	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		movimientoElmnt.DocumentoPresupuestalUUID = documentoPresupuestalOpStruct[0].Id.(string)
		movimientoElmnt.ValorActual = movimientoElmnt.ValorInicial
		movimientoElmnt.Estado = initialState
		movimientoParameter, err := movimientoManager.GetInitialMovimientoParameterByHijo(movimientoElmnt.Tipo)

		if err == nil && movimientoParameter.TipoDocumentoGenerado != nil && *movimientoParameter.TipoDocumentoGenerado != documentoPresupuestalRequestData.Tipo {
			// TODO: put this code on separate helper or manager.
			if consecutivoIndex[*movimientoParameter.TipoDocumentoGenerado] == 0 {
				consecutivoIndex[*movimientoParameter.TipoDocumentoGenerado] = len(docMananger.GetByType(strconv.Itoa(documentoPresupuestalRequestData.Vigencia), documentoPresupuestalRequestData.CentroGestor, *movimientoParameter.TipoDocumentoGenerado)) + 1
			} else {
				consecutivoIndex[*movimientoParameter.TipoDocumentoGenerado]++
			}
			generatedDocument := models.DocumentoPresupuestal{
				Tipo:          *movimientoParameter.TipoDocumentoGenerado,
				Data:          documentoPresupuestalRequestData.Data,
				FechaRegistro: documentoPresupuestalRequestData.FechaRegistro,
				ValorActual:   movimientoElmnt.ValorActual,
				ValorInicial:  movimientoElmnt.ValorInicial,
				Vigencia:      documentoPresupuestalRequestData.Vigencia,
				CentroGestor:  documentoPresupuestalRequestData.CentroGestor,
				Estado:        models.EstadoRegistrada,
				Consecutivo:   consecutivoIndex[*movimientoParameter.TipoDocumentoGenerado],
			}
			documentoPresupuestalTr := transactionManager.ConvertToTransactionItem(models.DocumentoPresupuestalCollection+collectionPostFixName, "", "", generatedDocument)
			afectationGenerated := movimientoElmnt
			afectationGenerated.Tipo = *movimientoParameter.TipoDocumentoGenerado
			afectationGenerated.DocumentoPresupuestalUUID = documentoPresupuestalTr[0].Id.(string)
			afectationGeneratedTr := transactionManager.ConvertToTransactionItem(models.MovimientosCollection+collectionPostFixName, "", "", afectationGenerated)
			movimientoData = append(movimientoData, documentoPresupuestalTr...)
			movimientoData = append(movimientoData, afectationGeneratedTr...)
		} else if err == nil && movimientoParameter.TipoDocumentoGenerado != nil && *movimientoParameter.TipoDocumentoGenerado == documentoPresupuestalRequestData.Tipo {
			panic("Document Generation of same type not allowed")
		}

		insertMovimientoData := transactionManager.ConvertToTransactionItem(models.MovimientosCollection+collectionPostFixName, "", "", movimientoElmnt)
		movimientoDataInserted = append(movimientoDataInserted, insertMovimientoData...)
		movimientoData = append(movimientoData, insertMovimientoData...)
		propagacionData := movimientohelper.BuildPropagacionValoresTr(movimientoElmnt, balance, afectationIndex, collectionPostFixName)

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
