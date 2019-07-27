package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"

	"github.com/udistrital/plan_cuentas_mongo_crud/compositors/movimientoCompositor"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/transactionManager"

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
	collectionName := models.MovimientosCollection

	var (
		movimientoRequestData []models.Movimiento
	)

	defer func() {
		j.Data["json"] = body
	}()

	if err := json.Unmarshal(j.Ctx.Input.RequestBody, &movimientoRequestData); err != nil {
		logManager.LogError(err.Error())
		body = err
		return
	}

	// Check for required fields in struct
	for _, movimientoElmnt := range movimientoRequestData {
		var movimientoData []interface{}
		var movimientoIntfc interface{}
		movimientoIntfc = movimientoElmnt
		if errStrc := formatdata.StructValidation(movimientoElmnt); len(errStrc) > 0 {
			responseformat.SetResponseFormat(&j.Controller, errStrc, "", 422)
		}

		insertMovimientoData := transactionManager.ConvertToTransactionItem(collectionName, movimientoIntfc)
		movimientoData = append(movimientoData, insertMovimientoData...)
		propagacionData := movimientoCompositor.BuildPropagacionValoresTr(movimientoElmnt)
		if len(propagacionData) > 0 {
			movimientoData = append(movimientoData, propagacionData...)
			logs.Debug(movimientoData)

		}
		transactionManager.RunTransaction(collectionName, movimientoData)
	}
	// Perform Mongo's Transaction.

	body = movimientoRequestData
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
		j.Data["json"] = body
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
