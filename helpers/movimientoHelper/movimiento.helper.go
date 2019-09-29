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
func BuildPropagacionValoresTr(movimiento models.Movimiento, balance map[string]models.DocumentoPresupuestal) (trData []txn.Op) {
	movimientoParameter, err := movimientoManager.GetOneMovimientoParameterByHijo(movimiento.Tipo)

	var (
		treeActualLevel      int
		arrMovimientosUpdted []interface{}
		runFlag              = true
	)

	if err != nil {
		logs.Error("1", err)
		return
	}
	movimientoPadre, err := movimientoManager.GetOneMovimientoByTipo(movimiento.Padre, movimientoParameter.TipoMovimientoPadre)

	movimientoHijo := movimiento
	var propagationName = movimientoHijo.Tipo

	if err != nil {
		runFlag = false
	}

	for runFlag {
		treeActualLevel++
		if len(movimientoPadre.Movimientos) == 0 {
			movimientoPadre.Movimientos = make(map[string]float64)
		}

		if movimientoPadre.Movimientos[movimientoHijo.Tipo] == 0 {
			movimientoPadre.Movimientos[propagationName] = movimientoHijo.ValorInicial * float64(movimientoParameter.Multiplicador)
		} else {
			movimientoPadre.Movimientos[propagationName] += (movimientoHijo.ValorInicial * float64(movimientoParameter.Multiplicador))
		}

		if treeActualLevel == 1 {

			movimientoPadre.ValorActual += movimientoHijo.ValorInicial * float64(movimientoParameter.Multiplicador)
			if movimientoPadre.ValorActual == 0 {
				movimientoPadre.Estado = "total_comprometido"
			} else if movimientoPadre.ValorActual > 0 {
				movimientoPadre.Estado = "parcial_comprometido"
			} else {
				errorMessage := "Cannot Perform operation, presupuestal document " + movimientoPadre.DocumentoPresupuestalUUID + " for bag " + movimientoPadre.ID + " has no balance left!"
				logs.Error(errorMessage)
				panic(errorMessage)
			}

			documentoPresupuestal := models.DocumentoPresupuestal{}

			if balance[movimientoPadre.DocumentoPresupuestalUUID].ID == "" {
				documentoPresupuestalIntfc, err := crudmanager.GetDocumentByID(movimientoPadre.DocumentoPresupuestalUUID, models.DocumentoPresupuestalCollection)
				formatdata.FillStructP(documentoPresupuestalIntfc, &documentoPresupuestal)
				if err == nil {
					documentoPresupuestal.ValorActual += (movimientoHijo.ValorInicial * float64(movimientoParameter.Multiplicador))
					balance[movimientoPadre.DocumentoPresupuestalUUID] = documentoPresupuestal
				}
			} else {
				documentoPresupuestal = balance[movimientoPadre.DocumentoPresupuestalUUID]
				documentoPresupuestal.ValorActual += (movimientoHijo.ValorInicial * float64(movimientoParameter.Multiplicador))
				balance[movimientoPadre.DocumentoPresupuestalUUID] = documentoPresupuestal
			}

			if balance[movimientoPadre.DocumentoPresupuestalUUID].ValorActual < 0 {
				errorMessage := "Cannot Perform operation, presupuestal document " + documentoPresupuestal.ID + " for bag " + movimientoPadre.ID + " has no balance left!"
				logs.Error(errorMessage)
				panic(errorMessage)
			} else {
				if documentoPresupuestal.ValorActual == 0 {
					documentoPresupuestal.Estado = "total_comprometido"
				} else {
					documentoPresupuestal.Estado = "parcial_comprometido"
				}
				trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(models.DocumentoPresupuestalCollection, "", "Estado,ValorActual", documentoPresupuestal)...)
			}
		}

		arrMovimientosUpdted = append(arrMovimientosUpdted, movimientoPadre)

		movimientoHijo = movimientoPadre
		movimientoParameter, err := movimientoManager.GetOneMovimientoParameterByHijo(movimientoHijo.Tipo)

		if err != nil {
			if err.Error() == "not found" {
				runFlag = false
			} else {
				logs.Error("2", err)
				panic(err)
			}
		} else {
			movimientoPadre.Movimientos = make(map[string]float64)
			movimientoPadre, err = movimientoManager.GetOneMovimientoByTipo(movimientoHijo.Padre, movimientoParameter.TipoMovimientoPadre)
			if err != nil {
				runFlag = false
			}
		}

	}

	trData = append(trData, transactionManager.ConvertToUpdateTransactionItem(models.MovimientosCollection, "", "Estado,Movimientos,ValorActual", arrMovimientosUpdted...)...)
	return
}
