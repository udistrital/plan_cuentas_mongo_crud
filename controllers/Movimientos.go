package controllers

import (
	"encoding/json"

	"github.com/udistrital/plan_cuentas_mongo_crud/compositors/movimientoCompositor"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"

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
// @Param	body		body 	models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router RegistrarMovimiento/:tipoPago [post]
func (j *MovimientosController) RegistrarMovimiento() {

	var (
		tipo           string
		movimientoData models.Movimiento
		requestData    map[string]interface{}
	)

	tipo = j.GetString(":tipoPago")

	defer func() {
		j.Data["json"] = movimientoData
	}()

	if err := json.Unmarshal(j.Ctx.Input.RequestBody, &requestData); err != nil {
		logManager.LogError(err.Error())
		panic(err.Error())
	}

	movimientoData = movimientoHelper.FormatMovimientoRequestData(requestData, tipo)

	movimientoCompositor.AddMovimientoTransaction(movimientoData)

}
