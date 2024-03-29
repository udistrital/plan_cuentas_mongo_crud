package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/globalsign/mgo" // Inicializa mgo para poder usar sus métodos
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// NecesidadesController estructura para un controlador de beego
type NecesidadesController struct {
	beego.Controller
	response map[string]interface{}
}

// URLMapping ...
func (j *NecesidadesController) URLMapping() {
	j.Mapping("GetAll", j.GetAll)
	// j.Mapping("Get", j.Get)
	j.Mapping("Post", j.Post)
	// j.Mapping("Put", j.Put)
	// j.Mapping("Delete", j.Delete)
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Param query query string true "Consulta"
// @Success 200 {object} []models.Necesidad
// @Failure 403 :objectId is empty
// @router / [get]
func (j *NecesidadesController) GetAll() {
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

	err := errors.New("Bad info response")

	response := DefaultResponse(403, err, nil)

	if obs := models.GetAllNecesidad(query); len(obs) > 0 {
		response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = response
	j.ServeJSON()
}

// Get obtiene un elemento por su id
// @Title Get
// @Description get Necesidad by nombre
// @Param	id		path 	string	true		"El id de la Necesidad a consultar"
// @Success 200 {object} models.Necesidad
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *NecesidadesController) Get() {
	objectId := j.GetString(":id")
	if objectId != "" {
		necesidad, err := models.GetNecesidadByID(objectId)
		if err == nil {
			j.response = DefaultResponse(200, nil, &necesidad)
		} else {
			j.response = DefaultResponse(403, err, nil)
		}
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// @Title Post
// @Description Post
// @Param	body		body 	models.Necesidad	true		"Body para la creacion de Necesidades"
// @Success 200 {object} string
// @Failure 403 body is empty
// @router / [post]
func (j *NecesidadesController) Post() {
	var necesidad models.Necesidad
	json.Unmarshal(j.Ctx.Input.RequestBody, &necesidad)

	if err := models.InsertNecesidad(&necesidad); err == nil {
		j.response = DefaultResponse(200, nil, "insert success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update the Necesidad
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Necesidad	true		"The body"
// @Success 200 {object} string
// @Failure 403 :id is empty
// @router /:id [put]
func (j *NecesidadesController) Put() {
	objectID := j.Ctx.Input.Param(":id")
	var necesidad models.Necesidad

	json.Unmarshal(j.Ctx.Input.RequestBody, &necesidad)

	if err := models.UpdateNecesidad(&necesidad, objectID); err == nil {
		j.response = DefaultResponse(200, nil, "update success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Delete ...
// @Title Borrar Necesidad
// @Description Borrar Necesidad
// @Param	id		path 	string	true		"El id del objeto que se quiere borrar"
// @Success 200 {object} string
// @Failure 403 id is empty
// @router /:id [delete]
func (j *NecesidadesController) Delete() {
	objectID := j.Ctx.Input.Param(":id")

	if err := models.DeleteNecesidad(objectID); err == nil {
		j.response = DefaultResponse(200, nil, "delete success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}
