package compositors

import (
	"github.com/astaxie/beego/logs"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	documentopresupuestalhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/documentoPresupuestalHelper"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

type DocumentoPresupuestalCompositor struct {
}

// GetAllByType This method will return documento presupuestal data with its afectation objects by its type , year and CG.
func (c *DocumentoPresupuestalCompositor) GetAllByType(vigencia, centroGestor, tipo string) []models.DocumentoPresupuestal {
	documentoPresupuestalData := documentopresupuestalmanager.GetByType(vigencia, centroGestor, tipo)
	movimientosData, err := movimientoManager.GetAllMovimiento(vigencia, centroGestor)

	if err != nil {
		logs.Error(err.Error())
		return nil
	}

	movInterfaceArr := commonhelper.ConvertToInterfaceArr(movimientosData)
	movimientoCollector := commonhelper.ArrToMapByKey("_id", movInterfaceArr...)
	for i := 0; i < len(documentoPresupuestalData); i++ {
		documentopresupuestalhelper.JoinDocumentoPresupuestalMovs(&documentoPresupuestalData[i], movimientoCollector)
	}

	return documentoPresupuestalData
}
