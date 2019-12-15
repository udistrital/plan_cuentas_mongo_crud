package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	vigenciahelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/vigenciaHelper"
	"github.com/udistrital/utils_oas/responseformat"
)

// VigenciaController estructura para un controlador de beego
type VigenciaController struct {
	beego.Controller
	response map[string]interface{}
}

// GetVigenciasByNameSpace ...
// @Title GetVigenciasByNameSpace
// @Description Retorna las vigencias a las cuales se ha registrado un name_space en el sistema.
// @Param	naemespace		path 	string	true		"name space al que pertenece el grupo de vigencias consultado"
// @Param	centrogestor		path 	string	true		"centro gestor al que pertenece el grupo de vigencias consultado"
// @Success 200 {string} success
// @Failure 403 error
// @router /:namespace/:centrogestor [get]
func (j *VigenciaController) GetVigenciasByNameSpace() {
	nameSpace := j.GetString(":namespace")
	cg := j.GetString(":centrogestor")
	vigArr, err := vigenciahelper.GetVigenciasByNameSpaceAndCg(nameSpace, cg)
	if err != nil {
		responseformat.SetResponseFormat(&j.Controller, err.Error(), "", 500)
	}
	responseformat.SetResponseFormat(&j.Controller, vigArr, "", 200)
}

// GetVigenciasCurrentVigenciaWithOffset ...
// @Title GetVigenciasCurrentVigenciaWithOffset
// @Description Retorna la vigencia actual según la hora del servidor y añade o quita segùn el offset lo indica , por defecto 0.
// @Param	offset		query 	string	true		"offset para determinar vigencia ej: offset= 1; vigencia+1 , offset=-1; vigencia -1"
// @Success 200 {string} success
// @Failure 403 error
// @router /vigencia_actual [get]
func (j *VigenciaController) GetVigenciasCurrentVigenciaWithOffset() {
	var offset int
	var err error
	if offsetStr := j.GetString("offset"); offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			responseformat.SetResponseFormat(&j.Controller, err.Error(), "", 500)
		}
	}
	currentTime := time.Now()
	year := currentTime.Year() + offset
	responseformat.SetResponseFormat(&j.Controller, year, "", 200)
}

// GetVigenciaActual ...
// @Title GetVigenciaActual
// @Description Retorna la vigencia actual a partir del estado en el que esta se encuentre. Posibles estados: actual, cerrada, creada
// @Success 200 {string} success
// @Failure 403 error
// @router /vigencia_actual_prueba [get]
func (j *VigenciaController) GetVigenciaActual() {
	var err error
	var objVigenciaActual []interface{}
	objVigenciaActual, err = vigenciahelper.GetVigenciaActual()
	if err != nil {
		responseformat.SetResponseFormat(&j.Controller, err, "", 403)
	}
	responseformat.SetResponseFormat(&j.Controller, objVigenciaActual[0], "", 200)
}

// AgregarVigencia ...
// @Title AgregarVigencia
// @Description create vigencia
// @Param	body		body 	models.Vigencia	true		"body for Producto content"
// @Success 201 {object} models.Vigencia
// @Failure 403 body is empty
// @router /agregar_vigencia [post]
func (j *VigenciaController) AgregarVigencia() {
	var vigencia map[string]interface{}
	json.Unmarshal(j.Ctx.Input.RequestBody, &vigencia)
	if err := vigenciahelper.AddNew(int((vigencia["Valor"]).(float64)), (vigencia["NameSpace"]).(string), (vigencia["AreaFuncional"]).(string), (vigencia["CentroGestor"]).(string), vigenciahelper.VigenciaActual); err == nil {
		j.response = DefaultResponse(201, nil, &vigencia)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}
