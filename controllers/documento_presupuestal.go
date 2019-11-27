package controllers

import (
	"github.com/astaxie/beego"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

type DocumentoPresupuestalController struct {
	beego.Controller
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 DocumentoPresupuestal models.DocumentoPresupuestal
// @Failure 403 :objectId is empty
// @router /:vigencia/:CG/:tipo [get]
func (j *DocumentoPresupuestalController) GetAll() {
	vigencia := j.GetString(":vigencia")
	centroGestor := j.GetString(":CG")
	tipo := j.GetString(":tipo")

	rows := documentopresupuestalmanager.GetByType(vigencia, centroGestor, tipo)

	response := commonhelper.DefaultResponse(200, nil, &rows)

	j.Data["json"] = response
	j.ServeJSON()
}

// GetAllCdp función para obtener todos los movimientos de CDP, de una vigencia, sin importar el centro gestor
// @Title GetAllCdp
// @Description get all cdp objects
// @Success 200 rows []models.DocumentoPresupuestal
// @Failure 403 :vigencia is empty
// @router get_all_cdp/:vigencia [get]
func (j *DocumentoPresupuestalController) GetAllCdp() {
	var rows []models.DocumentoPresupuestal
	vigencia := j.GetString(":vigencia")
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "1", "cdp")...)
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "2", "cdp")...)

	response := commonhelper.DefaultResponse(200, nil, &rows)

	j.Data["json"] = response
	j.ServeJSON()
}

// GetInfoCdp Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Title GetInfoCdp
// @Description Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Success 200 documentoPresupuestal models.DocumentoPresupuestal
// @Failure 403 :id is empty
// @router /get_info_cdp/:id [get]
func (j *DocumentoPresupuestalController) GetInfoCdp() {
	id := j.GetString(":id")
	data := documentopresupuestalmanager.GetCDPByID(id)

	response := commonhelper.DefaultResponse(200, nil, &data)
	j.Data["json"] = response
	j.ServeJSON()
}
