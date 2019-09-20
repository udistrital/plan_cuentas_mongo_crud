package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
	response map[string]interface{}
}

// URLMapping ...
func (j *FuenteFinanciamientoController) URLMapping() {
	j.Mapping("Post", j.Post)
	j.Mapping("Put", j.Put)
	j.Mapping("VincularFuente", j.VincularFuente)
	j.Mapping("Delete", j.Delete)
	j.Mapping("GetAll", j.GetAll)
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 FuenteFunanciamiento models.FuenteFinanciamiento
// @Failure 403 :objectId is empty
// @router / [get]
func (j *FuenteFinanciamientoController) GetAll() {
	var query = make(map[string]interface{})

	if v := j.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				j.Data["json"] = errors.New("Consulta invalida")
				j.ServeJSON()
				return
			}

			if i, err := strconv.Atoi(kv[1]); err == nil {
				k, v := kv[0], i
				query[k] = v
			} else {
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}
	}

	if obs, err := models.GetAllFuenteFinanciamiento(query); len(obs) == 0 {
		j.response = DefaultResponse(403, err, nil)
	} else {
		j.response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Get obtiene un elemento por su _id
// @Title Get
// @Description get FuenteFinancimiento by nombre
// @Param	nombre		path 	string	true		"El nombre de la FuenteFinancimiento a consultar"
// @Success 200 {object} models.FuenteFinancimiento
// @Failure 403 :uid is empty
// @router /:objectId [get]
func (j *FuenteFinanciamientoController) Get() {
	objectID := j.GetString(":objectId")
	if fuente, err := models.GetFuenteFinanciamientoByID(objectID); err == nil {
		j.response = DefaultResponse(200, nil, &fuente)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (j *FuenteFinanciamientoController) Post() {
	var fuente models.FuenteFinanciamiento
	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)

	if err := models.InsertFuenteFinanciamiento(&fuente); err == nil {
		j.response = DefaultResponse(200, nil, "insert success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// VincularFuente ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /VincularFuente [post]
func (j *FuenteFinanciamientoController) VincularFuente() {
	var fuente models.FuenteFinanciamiento
	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)
}

// Put de HTTP
// @Title Update
// @Description update the FuenteFinanciamiento
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *FuenteFinanciamientoController) Put() {
	objectID := j.Ctx.Input.Param(":objectId")
	var fuente models.FuenteFinanciamiento

	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)

	if err := models.UpdateFuenteFinanciamiento(&fuente, objectID); err == nil {
		j.response = DefaultResponse(200, nil, "update success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Delete ...
// @Title Borrar FuenteFinanciamiento
// @Description Borrar FuenteFinanciamiento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *FuenteFinanciamientoController) Delete() {
	objectID := j.Ctx.Input.Param(":objectId")

	if err := models.DeleteFuenteFinanciamiento(objectID); err == nil {
		j.response = DefaultResponse(200, nil, "delete success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// GetWithRubro ...
// @Title GetWithRubro
// @Description Borrar FuenteFinanciamiento
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /fuente_financiamiento_apropiacion/:rubro_apropiacion_id [get]
func (j *FuenteFinanciamientoController) GetWithRubro() {
	rubroApropiacionID := j.Ctx.Input.Param(":rubro_apropiacion_id")

	if fuentes, err := models.GetFuentesByRubroApropiacion(rubroApropiacionID); err == nil {
		j.response = DefaultResponse(200, nil, &fuentes)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}
