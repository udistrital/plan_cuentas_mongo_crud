package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
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

	obs, err := models.GetAllFuenteFinanciamiento(query)

	if len(obs) == 0 {
		j.Data["json"] = err
	} else {
		j.Data["json"] = &obs
	}

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
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = err.Error()
	}

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
	fmt.Println("fuente:", fuente)
}

// Put de HTTP
// @Title Update
// @Description update the FuenteFinanciamiento
// @Param	codigo		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :codigo is empty
// @router /:codigo [put]
func (j *FuenteFinanciamientoController) Put() {
	codigo := j.Ctx.Input.Param(":codigo")
	var fuente models.FuenteFinanciamiento

	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)

	if err := models.UpdateFuenteFinanciamiento(&fuente, codigo); err == nil {
		j.Data["json"] = "update success!"
	} else {
		j.Data["json"] = err.Error()
	}

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

	if err := models.DeleteFuenteFinanciamiento(objectID); err != nil {
		j.Data["json"] = "delete success!"
	} else {
		j.Data["json"] = err.Error()
	}

	j.ServeJSON()
}
