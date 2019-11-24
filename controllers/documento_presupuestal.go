package controllers

import (
	"github.com/astaxie/beego"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
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
