package controllers

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	vigenciahelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/vigenciaHelper"
	"github.com/udistrital/utils_oas/responseformat"
)

// VigenciaController estructura para un controlador de beego
type VigenciaController struct {
	beego.Controller
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
