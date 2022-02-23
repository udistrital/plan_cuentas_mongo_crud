package compositors

import (
	"strings"

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

func (c *DocumentoPresupuestalCompositor) GetMovDocumentPresByParent(vigencia, centroGestor, parentUUID string) []models.DocumentoPresupuestal {
	collectionFixedName := "_" + vigencia + "_" + centroGestor
	var docPresArr []models.DocumentoPresupuestal
	movs, err := movimientoManager.GetAllMovimientoByParentUUID(parentUUID, collectionFixedName)
	if err != nil {
		logs.Error("error at GetMovDocumentPresByParent ", err.Error())
	}

	for _, mov := range movs {
		doc, err := documentopresupuestalmanager.GetOneByType(mov.DocumentoPresupuestalUUID, vigencia, centroGestor, "")
		if err == nil {
			docPresArr = append(docPresArr, doc)
		}
	}

	return docPresArr
}

func (c *DocumentoPresupuestalCompositor) GetAllDocumentoPresupuestalMovimientosByRubro(vigencia, centroGestor, rubro string) (d []models.DocumentoPresupuestal, err error) {
	var docs []models.DocumentoPresupuestal
	predocs := c.GetAllByType(vigencia, centroGestor, "cdp")
	for i := range predocs {
		if predocs[i].Estado == "expedido" {
			for j := range predocs[i].Afectacion {
				r := predocs[i].Afectacion[j].Padre
				if strings.Contains(r, rubro) {
					Rubro, err := models.GetNodoRubroReducidoById(r)
					if err != nil {
						return d, err
					}
					predocs[i].Afectacion[j].RubroDetalle = &Rubro
					docs = append(docs, predocs[i])
				}
			}
		}
	}
	return docs, nil
}
