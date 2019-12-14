package movimientohelper

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/txn"
	crudmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/crudManager"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// BuildPropagacionValoresTr ... Build a mgo transaction item as Array of interfaces .
// This method search in "movimientos_parametros" collection for the afectation's config recursively.
func BuildPropagacionValoresTr(movimiento models.Movimiento, balance, afectationIndex map[string]map[string]interface{}, collectionPostFixName string) (trData []txn.Op) {
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
	var movimientoPadre interface{}
	var errMap error
	if afectationIndex[movimiento.Padre] != nil {
		documentoPadre = afectationIndex[movimiento.Padre]
	} else {
		movimientoPadre, err = crudmanager.GetDocumentByID(movimiento.Padre, propagationCollectionName)
		if err != nil {
			logs.Error(err.Error(), movimiento.Padre, propagationCollectionName)
			runFlag = false
		}
		documentoPadre, errMap = formatdata.ToMap(movimientoPadre, "bson")
		if errMap != nil {
			panic(errMap.Error())
		}
	}

	movimientoHijo := make(map[string]interface{})

	movimientoHijo, errMap = formatdata.ToMap(movimiento, "bson")
	if errMap != nil {
		panic(errMap.Error())
	}

	var propagationName = movimientoHijo["tipo"].(string)
	var propagationValue = movimientoHijo["valor_inicial"].(float64)
	var documentoPadreValorActual float64
	var documentoPadreNewState string
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

		if documentoPadre["valor_actual"] != nil {
			documentoPadreValorActual = documentoPadre["valor_actual"].(float64)
		}
		documentoPadreValorActual += propagationValue * float64(movimientoParameter.Multiplicador)
		if documentoPadre["estado"] != nil {
			documentoPadreNewState = documentoPadre["estado"].(string)
		}

		if documentoPadreValorActual == 0 {
			documentoPadreNewState = "total_comprometido"
		} else if documentoPadreValorActual > 0 {
			documentoPadreNewState = "parcial_comprometido"
		} else {
			errorMessage := ""
			if documentoPadre["documento_presupuestal_uuid"] == nil {
				errorMessage = "Saldo Insuficiente en " + documentoPadre[fatherUUIKey].(string)
			} else {
				errorMessage = "Saldo Insuficiente en" + documentoPadre[fatherUUIKey].(string) // "presupuestal document " + documentoPadre["documento_presupuestal_uuid"].(string) + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
			}
			logs.Error(errorMessage)
			panic(errorMessage)
		}
		if !movimientoParameter.WithOutChangeState {
			documentoPadre["estado"] = documentoPadreNewState
		}

		documentoPadre["valor_actual"] = documentoPadreValorActual

		documentoPresupuestal := make(map[string]interface{})
		balanceKey := ""
		if documentoPadre["documento_presupuestal_uuid"] != nil {
			balanceKey = "documento_presupuestal_uuid"
		} else if documentoPadre["_id"] != nil {
			balanceKey = "_id"
		}

		if balanceKey != "" {
			balanceCollectionName := documentoPresupuestalFixedCollectionName
			if movimientoParameter.FatherCollectionName != "" {
				balanceCollectionName = movimientoParameter.FatherCollectionName + collectionPostFixName
			}
			if balance[documentoPadre[balanceKey].(string)] == nil {
				balance[documentoPadre[balanceKey].(string)] = make(map[string]interface{})

				formatdata.JsonPrint(movimientoParameter.Multiplicador)
				formatdata.JsonPrint(documentoPadre)
				documentoPresupuestalIntfc, err := crudmanager.GetDocumentByID(documentoPadre[balanceKey].(string), balanceCollectionName)
				if err == nil {
					formatdata.FillStructP(documentoPresupuestalIntfc, &documentoPresupuestal)
					documentoPresupuestal["valor_actual"] = documentoPresupuestal["valor_actual"].(float64) + (propagationValue * float64(movimientoParameter.Multiplicador))
					balance[documentoPadre[balanceKey].(string)] = documentoPresupuestal
				}
			} else {
				documentoPresupuestal = balance[documentoPadre[balanceKey].(string)]
				documentoPresupuestal["valor_actual"] = documentoPresupuestal["valor_actual"].(float64) + (propagationValue * float64(movimientoParameter.Multiplicador))
				balance[documentoPadre[balanceKey].(string)] = documentoPresupuestal
			}

			if balance[documentoPadre[balanceKey].(string)]["valor_actual"].(float64) < 0 {
				errorMessage := "Cannot Perform operation, presupuestal document " + documentoPresupuestal["_id"].(string) + " for bag " + documentoPadre[fatherUUIKey].(string) + " has no balance left!"
				logs.Error(errorMessage)
				panic(errorMessage)
			} else {
				if !movimientoParameter.WithOutChangeState {
					if documentoPresupuestal["valor_actual"].(float64) == 0 {
						documentoPresupuestal["estado"] = "total_comprometido"
					} else {
						documentoPresupuestal["estado"] = "parcial_comprometido"
					}
					trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(balanceCollectionName, "", "estado,valor_actual", documentoPresupuestal)...)

				}

			}
		}
		if documentoPadre["_id"] != nil {
			documentoPadre["movimientos"] = movimientosPadreData
			afectationIndex[documentoPadre["_id"].(string)] = documentoPadre
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
				if movimientoHijo["padre"] != nil {

					if afectationIndex[movimientoHijo["padre"].(string)] != nil {
						documentoPadre = afectationIndex[movimientoHijo["padre"].(string)]
					} else {
						movimientoPadre, err = crudmanager.GetDocumentByID(movimientoHijo["padre"].(string), propagationCollectionName)
						if err != nil {
							logs.Error(err.Error(), movimientoHijo["padre"].(string), propagationCollectionName)
							runFlag = false
						}
						documentoPadre, errMap = formatdata.ToMap(movimientoPadre, "bson")
						if errMap != nil {
							panic(errMap.Error())
						}
						if err != nil {
							runFlag = false
						}

					}
				}

			}
		} else {
			runFlag = false
		}

	}
	return
}

func JoinGeneratedDocPresWithMov(movimientos []models.Movimiento, vigencia, cg string) (result []map[string]interface{}, err error) {
	var movimientosJoined []map[string]interface{}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
			return
		}
	}()
	for _, mov := range movimientos {
		movMap := make(map[string]interface{})
		if err := formatdata.FillStruct(mov, &movMap); err != nil {
			return nil, errors.New("Cannont get generated document info")
		}
		if mov.DocumentosPresGenerados != nil {
			var documentsGenerated []models.DocumentoPresupuestal
			for _, doc := range *mov.DocumentosPresGenerados {
				docGenerated, err := documentopresupuestalmanager.GetOneByType(doc, vigencia, cg, mov.Tipo)
				if err == nil {
					documentsGenerated = append(documentsGenerated, docGenerated)
				}
			}
			movMap["DocumentsGenerated"] = documentsGenerated
		}
		movimientosJoined = append(movimientosJoined, movMap)
	}
	return movimientosJoined, nil
}
