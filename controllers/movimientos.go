package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/udistrital/plan_cuentas_mongo_crud/compositors/movimientoCompositor"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	movimientohelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
)

// MovimientosController estructura para un controlador de beego
type MovimientosController struct {
	beego.Controller
}

// RegistrarMovimiento ...
// @Title RegistrarMovimiento
// @Description Registra los movimientos (como cdp, rp, ver variable tipoMovimiento) y los propaga tanto en la colección
// arbolrubrosapropiacion_[vigencia]_[unidad_ejecutura], como en la colección movimientos. Utiliza la función registrarValores para registrar los valores,
// y se le envian como párametro el nombre de los movimientos que se van a guardar en el atributo movimiento de la colección arbolrubrosapropiacion,
// al igual que se envia la variable dataValor, que son los valores del movimiento enviados desde el api_mid_financiera
// @Param	body		body 	[]models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router /RegistrarMovimientos [post]
func (j *MovimientosController) RegistrarMovimiento() {
	var body interface{}

	var (
		documentoPresupuestalRequestData models.DocumentoPresupuestal
	)

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			responseformat.SetResponseFormat(&j.Controller, r, "", 500)
		}
		responseformat.SetResponseFormat(&j.Controller, body, "", 200)

	}()

	if err := json.Unmarshal(j.Ctx.Input.RequestBody, &documentoPresupuestalRequestData); err != nil {
		logManager.LogError(err.Error())
		panic(err.Error())
	}
	if errStrc := formatdata.StructValidation(&documentoPresupuestalRequestData); len(errStrc) > 0 {
		logs.Error(errStrc)
		responseformat.SetResponseFormat(&j.Controller, errStrc, "", 422)
		return
	}

	for _, movimientoElmnt := range documentoPresupuestalRequestData.Afectacion {

		if errStrc := formatdata.StructValidation(movimientoElmnt); len(errStrc) > 0 {
			logs.Error(errStrc)
			responseformat.SetResponseFormat(&j.Controller, errStrc, "", 422)
			return
		}
	}

	resulData := movimientoCompositor.DocumentoPresupuestalRegister(&documentoPresupuestalRequestData)
	body = resulData
}

// RegistrarMovimientoParameter ...
// @Title RegistrarMovimientoParameter
// @Description Registra los movimientos (como cdp, rp, ver variable tipoMovimiento) y los propaga tanto en la colección
// arbolrubrosapropiacion_[vigencia]_[unidad_ejecutura], como en la colección movimientos. Utiliza la función registrarValores para registrar los valores,
// y se le envian como párametro el nombre de los movimientos que se van a guardar en el atributo movimiento de la colección arbolrubrosapropiacion,
// al igual que se envia la variable dataValor, que son los valores del movimiento enviados desde el api_mid_financiera
// @Param	body		body 	[]models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router /RegistrarMovimientoParameter [post]
func (j *MovimientosController) RegistrarMovimientoParameter() {
	var body interface{}
	var (
		movimientoParamRequestData models.MovimientoParameter
	)

	defer func() {
		if r := recover(); r != nil {
			responseformat.SetResponseFormat(&j.Controller, r, "", 500)
		}
		responseformat.SetResponseFormat(&j.Controller, body, "", 200)
	}()

	if err := json.Unmarshal(j.Ctx.Input.RequestBody, &movimientoParamRequestData); err != nil {
		logManager.LogError(err.Error())
		body = err
		return
	}

	if errStrc := formatdata.StructValidation(movimientoParamRequestData); len(errStrc) > 0 {
		responseformat.SetResponseFormat(&j.Controller, errStrc, "", 422)
	}

	if err := movimientoManager.SaveMovimientoParameter(&movimientoParamRequestData); err == nil {
		body = movimientoParamRequestData
	} else {
		body = err
	}

}

// GetMovimientosByDocumentoPresupuestalUUID función para obtener todos los objetos por parentUUID
// @Title GetMovimientosByDocumentoPresupuestalUUID
// @Description get all objects
// @Success 200 Movimiento models.Movimiento
// @Failure 403 :objectId is empty
// @router /:vigencia/:CG/:parentUUID [get]
func (j *MovimientosController) GetMovimientosByDocumentoPresupuestalUUID() {
	vigencia := j.GetString(":vigencia")
	centroGestor := j.GetString(":CG")
	parentUUID := j.GetString(":parentUUID")

	rows, err := movimientoManager.GetMovimientoByDocumentoPresupuestalUUID(vigencia, centroGestor, parentUUID)
	rowsJoined, err := movimientohelper.JoinGeneratedDocPresWithMov(rows, vigencia, centroGestor)
	response := commonhelper.DefaultResponse(200, err, &rowsJoined)

	j.Data["json"] = response
	j.ServeJSON()
}
