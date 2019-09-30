package movimientohelper

import (
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/txn"
	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// BuildPropagacionValoresTr ... Build a mgo transaction item as Array of interfaces .
// This method search in "movimientos_parametros" collection for the afectation's config recursively.
func BuildPropagacionValoresTr(movimiento models.Movimiento, balance map[string]models.DocumentoPresupuestal, collectionPostFixName string) (trData []txn.Op) {
	propagationCollectionName := models.MovimientosCollection
	fatherUUIKey := "_id"
	movimientoParameter, err := movimientoManager.GetOneMovimientoParameterByHijo(movimiento.Tipo)

	if err == nil && movimientoParameter.FatherCollectionName != "" {
		propagationCollectionName = movimientoParameter.FatherCollectionName
		if movimientoParameter.FatherUUIKeyName != "" {
			fatherUUIKey = movimientoParameter.FatherUUIKeyName
		}
	}

	propagationCollectionName += collectionPostFixName
	documentoPresupuestalFixedCollectionName := models.DocumentoPresupuestalCollection + collectionPostFixName

	var (
		treeActualLevel int
		runFlag         = true
	)

	if err != nil {
		logs.Error("1", err)
		return
	}
	documentoPadre := make(map[string]interface{})
	if propagationCollectionName == models.MovimientosCollection+collectionPostFixName {

		movimientoPadre, err := movimientoManager.GetOneMovimientoByTipo(movimiento.Padre, movimientoParameter.TipoMovimientoPadre, collectionPostFixName)
		formatdata.FillStructP(movimientoPadre, &documentoPadre)
		if err != nil {
			runFlag = false
		}
	} else {
		movimientoPadre, err := crudmanager.GetDocumentByID(movimiento.Padre, propagationCollectionName)
		formatdata.FillStructP(movimientoPadre, &documentoPadre)
		if err != nil {
			logs.Error(err.Error(), movimiento.Padre, propagationCollectionName)
			runFlag = false
		}
	}

	movimientoHijo := make(map[string]interface{})
	formatdata.FillStructP(movimiento, &movimientoHijo)
	var propagationName = movimientoHijo["Tipo"].(string)
	for runFlag {
		treeActualLevel++
		if documentoPadre["Movimientos"] == nil {
			documentoPadre["Movimientos"] = make(map[string]interface{})
		}

		movimientosPadreData := documentoPadre["Movimientos"].(map[string]interface{})

		if movimientosPadreData[propagationName] == nil {
			movimientosPadreData[propagationName] = movimientoHijo["ValorInicial"].(float64) * float64(movimientoParameter.Multiplicador)
		} else {
			newMovimeintoPadreValorActual := movimientosPadreData[propagationName].(float64) + (movimientoHijo["ValorInicial"].(float64) * float64(movimientoParameter.Multiplicador))
			movimientosPadreData[propagationName] = newMovimeintoPadreValorActual
		}

		if treeActualLevel == 1 {
			documentoPadreValorActual := documentoPadre["ValorActual"].(float64)
			documentoPadreValorActual += movimientoHijo["ValorInicial"].(float64) * float64(movimientoParameter.Multiplicador)
			if documentoPadreValorActual == 0 {
				documentoPadre["Estado"] = "total_comprometido"
			} else if documentoPadreValorActual > 0 {
				documentoPadre["Estado"] = "parcial_comprometido"
			} else {
				errorMessage := ""
				if documentoPadre["DocumentoPresupuestalUUID"] == nil {
					errorMessage = "Cannot Perform operation, bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
				} else {
					formatdata.JsonPrint(documentoPadre)
					errorMessage = "Cannot Perform operation, presupuestal document " + documentoPadre["DocumentoPresupuestalUUID"].(string) + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
				}
				logs.Error(errorMessage)
				panic(errorMessage)
			}

			documentoPadre["ValorActual"] = documentoPadreValorActual

			documentoPresupuestal := models.DocumentoPresupuestal{}
			if documentoPadre["DocumentoPresupuestalUUID"] != nil {
				if balance[documentoPadre["DocumentoPresupuestalUUID"].(string)].ID == "" {
					documentoPresupuestalIntfc, err := crudmanager.GetDocumentByID(documentoPadre["DocumentoPresupuestalUUID"].(string), documentoPresupuestalFixedCollectionName)
					if err == nil {
						formatdata.FillStructP(documentoPresupuestalIntfc, &documentoPresupuestal)
						documentoPresupuestal.ValorActual += (movimientoHijo["ValorInicial"].(float64) * float64(movimientoParameter.Multiplicador))
						balance[documentoPadre["DocumentoPresupuestalUUID"].(string)] = documentoPresupuestal
					}
				} else {
					documentoPresupuestal = balance[documentoPadre["DocumentoPresupuestalUUID"].(string)]
					documentoPresupuestal.ValorActual += (movimientoHijo["ValorInicial"].(float64) * float64(movimientoParameter.Multiplicador))
					balance[documentoPadre["DocumentoPresupuestalUUID"].(string)] = documentoPresupuestal
				}
				if balance[documentoPadre["DocumentoPresupuestalUUID"].(string)].ValorActual < 0 {
					errorMessage := "Cannot Perform operation, presupuestal document " + documentoPresupuestal.ID + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
					logs.Error(errorMessage)
					panic(errorMessage)
				} else {
					if documentoPresupuestal.ValorActual == 0 {
						documentoPresupuestal.Estado = "total_comprometido"
					} else {
						documentoPresupuestal.Estado = "parcial_comprometido"
					}
					trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(documentoPresupuestalFixedCollectionName, "", "Estado,ValorActual", documentoPresupuestal)...)
				}
			}
		}

		documentoPadre["Movimientos"] = movimientosPadreData
		trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(propagationCollectionName, fatherUUIKey, "Estado,Movimientos,ValorActual", documentoPadre)...)
		formatdata.JsonPrint(trData)
		movimientoHijo = documentoPadre
		if propagationCollectionName == models.MovimientosCollection+collectionPostFixName {

			movimientoParameter, err = movimientoManager.GetOneMovimientoParameterByHijo(movimientoHijo["Tipo"].(string))
			if err != nil {
				if err.Error() == "not found" {
					runFlag = false
				} else {
					logs.Error("2", err)
					panic(err)
				}
			} else {
				documentoPadre = make(map[string]interface{})

				movimientoPadre, err := movimientoManager.GetOneMovimientoByTipo(movimientoHijo["Padre"].(string), movimientoParameter.TipoMovimientoPadre, collectionPostFixName)

				formatdata.FillStructP(movimientoPadre, &documentoPadre)

				if err != nil {
					runFlag = false
				}

			}
		} else {
			movimientoParameter, err = movimientoManager.GetOneMovimientoParameterByHijo(movimientoParameter.TipoMovimientoPadre)
			if err != nil {
				runFlag = false
			} else {
				documentoPadre = make(map[string]interface{})
				movimientoPadre, err := crudmanager.GetDocumentByID(movimientoHijo["Padre"].(string), propagationCollectionName)
				formatdata.FillStructP(movimientoPadre, &documentoPadre)
				if err != nil {
					runFlag = false
				}
			}
		}

	}
	return
}
