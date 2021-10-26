package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	vigenciahelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/vigenciaHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
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
// @Param	namespace		path 	string	true		"name space al que pertenece el grupo de vigencias consultado"
// @Param	centrogestor		path 	string	true		"centro gestor al que pertenece el grupo de vigencias consultado"
// @Success 200 {object} []map[string]int
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
// @Success 200 {object} uint
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
// @Description Retorna la vigencia del área funcional cuyo estado sea actual.
// @Param area_funcional 	path 	string	true	"Área funcional a la que pertenece la vigencia que se quiere consultar"
// @Success 200 {object} []interface{}
// @Failure 403 error
// @router /vigencia_actual_area/:area_funcional [get]
func (j *VigenciaController) GetVigenciaActual() {
	var err error
	var objVigenciaActual []interface{}
	objVigenciaActual, err = vigenciahelper.GetVigenciaActual(j.GetString(":area_funcional"))
	if err != nil || len(objVigenciaActual) == 0 {
		responseformat.SetResponseFormat(&j.Controller, err, "", 403)
	}

	responseformat.SetResponseFormat(&j.Controller, objVigenciaActual, "", 200)
}

// GetTodasVigencias ...
// @Title GetTodasVigencias
// @Description Retorna las vigencias guardadas en las diferentes colecciones de la base de datos.
// @Success 200 {object} []interface{}
// @Failure 403 error
// @router /vigencias_total [get]
func (j *VigenciaController) GetTodasVigencias() {
	var err error
	var vigenciasArr []interface{}
	vigenciasArr, err = vigenciahelper.GetTodasVigencias()
	if err != nil || len(vigenciasArr) == 0 {
		responseformat.SetResponseFormat(&j.Controller, err, "", 403)
	}
	responseformat.SetResponseFormat(&j.Controller, vigenciasArr, "", 200)
}

// CerrarVigencia ...
// @Title CerrarVigencia
// @Description Se cierra la vigencia que se encuentre con estado actual en la colección, dependiendo del área funcional que le llegue.
// @Param area_funcional 	path 	string	true	"Área funcional a la que pertenece la vigencia que se quiere cerrar."
// @Success 200 {object} map[string]interface{}
// @Failure 403 error
// @router /cerrar_vigencia_actual/:area_funcional [get]
func (j *VigenciaController) CerrarVigencia() {
	err := vigenciahelper.CerrarVigencia(j.GetString(":area_funcional"))
	j.response = DefaultResponse(200, err, "vigencia cerrada")

	j.Data["json"] = j.response
	j.ServeJSON()
}

// AgregarVigencia ...
// @Title AgregarVigencia
// @Description create vigencia
// @Param	body		body 	models.VigenciaNueva	true		"body for Producto content"
// @Success 201 {object} models.VigenciaNueva
// @Failure 403 body is empty
// @router /agregar_vigencia [post]
func (j *VigenciaController) AgregarVigencia() {
	var vigencia models.VigenciaNueva
	json.Unmarshal(j.Ctx.Input.RequestBody, &vigencia)
	if err := vigenciahelper.AddNew(vigencia.Valor, vigenciahelper.VigenciaActual, vigencia.AreaFuncional); err == nil {
		j.response = DefaultResponse(201, nil, &vigencia)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}
