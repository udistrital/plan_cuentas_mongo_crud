package movimientohelper

import (
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/txn"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
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
	movimientoParameter, err := movimientoManager.GetInitialMovimientoParameterByHijo(movimiento.Tipo)

	if err == nil && movimientoParameter.FatherCollectionName != "" {
		propagationCollectionName = movimientoParameter.FatherCollectionName
		if movimientoParameter.FatherUUIKeyName != "" {
			fatherUUIKey = movimientoParameter.FatherUUIKeyName
		}
	}

	propagationCollectionName += collectionPostFixName
	documentoPresupuestalFixedCollectionName := models.DocumentoPresupuestalCollection + collectionPostFixName

	var runFlag = true

	if err != nil {
		logs.Error("1", err)
		return
	}
	documentoPadre := make(map[string]interface{})
	movimientoPadre, err := crudmanager.GetDocumentByID(movimiento.Padre, propagationCollectionName)
	documentoPadre, errMap := commonhelper.ToMap(movimientoPadre, "bson")

	if errMap != nil {
		panic(errMap.Error())
	}

	if err != nil {
		logs.Error(err.Error(), movimiento.Padre, propagationCollectionName)
		runFlag = false
	}

	movimientoHijo := make(map[string]interface{})

	movimientoHijo, errMap = commonhelper.ToMap(movimiento, "bson")
	if errMap != nil {
		panic(errMap.Error())
	}

	var propagationName = movimientoHijo["tipo"].(string)
	var propagationValue = movimientoHijo["valor_inicial"].(float64)
	for runFlag {
		if documentoPadre["movimientos"] == nil {
			documentoPadre["movimientos"] = make(map[string]interface{})
		}

		movimientosPadreData := documentoPadre["movimientos"].(map[string]interface{})

		if movimientosPadreData[propagationName] == nil {
			movimientosPadreData[propagationName] = propagationValue * float64(movimientoParameter.Multiplicador)
		} else {
			newMovimeintoPadreValorActual := movimientosPadreData[propagationName].(float64) + (propagationValue * float64(movimientoParameter.Multiplicador))
			movimientosPadreData[propagationName] = newMovimeintoPadreValorActual
		}

		documentoPadreValorActual := documentoPadre["valor_actual"].(float64)
		documentoPadreValorActual += propagationValue * float64(movimientoParameter.Multiplicador)
		if documentoPadreValorActual == 0 {
			documentoPadre["estado"] = "total_comprometido"
		} else if documentoPadreValorActual > 0 {
			documentoPadre["estado"] = "parcial_comprometido"
		} else {
			errorMessage := ""
			if documentoPadre["documento_presupuestal_uuid"] == nil {
				errorMessage = "Cannot Perform operation, bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
			} else {
				errorMessage = "Cannot Perform operation, presupuestal document " + documentoPadre["documento_presupuestal_uuid"].(string) + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
			}
			logs.Error(errorMessage)
			panic(errorMessage)
		}

		documentoPadre["valor_actual"] = documentoPadreValorActual

		documentoPresupuestal := models.DocumentoPresupuestal{}
		if documentoPadre["documento_presupuestal_uuid"] != nil {
			if balance[documentoPadre["documento_presupuestal_uuid"].(string)].ID == "" {
				documentoPresupuestalIntfc, err := crudmanager.GetDocumentByID(documentoPadre["documento_presupuestal_uuid"].(string), documentoPresupuestalFixedCollectionName)
				if err == nil {
					formatdata.FillStructP(documentoPresupuestalIntfc, &documentoPresupuestal)
					documentoPresupuestal.ValorActual += (propagationValue * float64(movimientoParameter.Multiplicador))
					balance[documentoPadre["documento_presupuestal_uuid"].(string)] = documentoPresupuestal
				}
			} else {
				documentoPresupuestal = balance[documentoPadre["documento_presupuestal_uuid"].(string)]
				documentoPresupuestal.ValorActual += (propagationValue * float64(movimientoParameter.Multiplicador))
				balance[documentoPadre["documento_presupuestal_uuid"].(string)] = documentoPresupuestal
			}
			if balance[documentoPadre["documento_presupuestal_uuid"].(string)].ValorActual < 0 {
				errorMessage := "Cannot Perform operation, presupuestal document " + documentoPresupuestal.ID + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
				logs.Error(errorMessage)
				panic(errorMessage)
			} else {
				if documentoPresupuestal.ValorActual == 0 {
					documentoPresupuestal.Estado = "total_comprometido"
				} else {
					documentoPresupuestal.Estado = "parcial_comprometido"
				}
				trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(documentoPresupuestalFixedCollectionName, "", "estado,valor_actual", documentoPresupuestal)...)
			}
		}

		documentoPadre["movimientos"] = movimientosPadreData
		trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(propagationCollectionName, fatherUUIKey, "estado,movimientos,valor_actual", documentoPadre)...)
		movimientoHijo = documentoPadre
		nextMovimientoParameter, _ := movimientoManager.GetInitialMovimientoParameterByHijo(movimientoParameter.TipoMovimientoPadre)
		if nextMovimientoParameter.TipoMovimientoPadre != "" {
			movimientoParameter = nextMovimientoParameter
		}
		movimientoParameter, err = movimientoManager.GetOneMovimientoParameterByHijoAndPadre(propagationName, movimientoParameter.TipoMovimientoPadre)

		if err != nil {
			runFlag = false
		} else {
			if movimientoParameter.FatherCollectionName != "" {
				propagationCollectionName = movimientoParameter.FatherCollectionName
				if movimientoParameter.FatherUUIKeyName != "" {
					fatherUUIKey = movimientoParameter.FatherUUIKeyName
				}
			}

			propagationCollectionName += collectionPostFixName
			documentoPadre = make(map[string]interface{})
			movimientoPadre, err = crudmanager.GetDocumentByID(movimientoHijo["padre"].(string), propagationCollectionName)
			documentoPadre, errMap = commonhelper.ToMap(movimientoPadre, "bson")
			if errMap != nil {
				panic(errMap.Error())
			}
			if err != nil {
				runFlag = false
			}
		}

	}
	return
}
