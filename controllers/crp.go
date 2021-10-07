package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/globalsign/mgo" // Inicializa mgo para poder usar sus métodos
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// SolicitudesCRPController estructura para un controlador de beego
type SolicitudesCRPController struct {
	beego.Controller
	response map[string]interface{}
}

// URLMapping ...
func (j *SolicitudesCRPController) URLMapping() {
	j.Mapping("GetAll", j.GetAll)
	j.Mapping("Get", j.Get)
	j.Mapping("Post", j.Post)
	j.Mapping("Put", j.Put)
	j.Mapping("Delete", j.Delete)
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 {object} []models.SolicitudCRP
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..., if the filter value includes !$ at the begining, the value won't be converted to int"
// @Failure 403 :objectId is empty
// @router / [get]
func (j *SolicitudesCRPController) GetAll() {
	var query = make(map[string]interface{})

	if v := j.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				j.Data["json"] = errors.New("Consulta invalida")
				j.ServeJSON()
				return
			}

			if i, err := strconv.Atoi(kv[1]); err == nil && !strings.Contains(kv[1], "!$") {
				k, v := kv[0], i
				query[k] = v
			} else {
				kv[1] = strings.Split(kv[1], "!$")[1]
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}
	}

	err := errors.New("No data")

	response := DefaultResponse(204, err, nil)
	if obs := models.GetAllSolicitudCRP(query); len(obs) > 0 {
		response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = response
	j.ServeJSON()
}

// Get ...
// Get obtiene un elemento por su id
// @Title Get
// @Description get SolicitudCRP by nombre
// @Param	id		path 	string	true		"El nombre de la SolicitudCRP a consultar"
// @Success 200 {object} models.SolicitudCRP
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *SolicitudesCRPController) Get() {
	objectId := j.GetString(":id")

	if objectId != "" {
		SolicitudCRP, err := models.GetSolicitudCRPByID(objectId)
		if err == nil {
			j.response = DefaultResponse(200, nil, &SolicitudCRP)
		} else {
			j.response = DefaultResponse(403, err, nil)
		}
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// @Title Post
// @Description Post
// @Param	body		body 	models.SolicitudCRP	true		"Body para la creacion de SolicitudesCRP"
// @Success 200 {object} string
// @Failure 403 body is empty
// @router / [post]
func (j *SolicitudesCRPController) Post() {
	var SolicitudCRP models.SolicitudCRP
	json.Unmarshal(j.Ctx.Input.RequestBody, &SolicitudCRP)
	fmt.Println(j.Ctx.Input.RequestBody, &SolicitudCRP)
	if err := models.InsertSolicitudCRP(&SolicitudCRP); err == nil {
		j.response = DefaultResponse(200, nil, "insert success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update the SolicitudCRP
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SolicitudCRP	true		"The body"
// @Success 200 {object} string
// @Failure 403 :id is empty
// @router /:id [put]
func (j *SolicitudesCRPController) Put() {
	objectID := j.Ctx.Input.Param(":id")
	var SolicitudCRP models.SolicitudCRP

	json.Unmarshal(j.Ctx.Input.RequestBody, &SolicitudCRP)

	if err := models.UpdateSolicitudCRP(&SolicitudCRP, objectID); err == nil {
		j.response = DefaultResponse(200, nil, "update success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Delete ...
// @Title Borrar SolicitudCRP
// @Description Borrar SolicitudCRP
// @Param	id		path 	string	true		"El id del objeto que se quiere borrar"
// @Success 200 {object} string
// @Failure 403 id is empty
// @router /:id [delete]
func (j *SolicitudesCRPController) Delete() {
	objectID := j.Ctx.Input.Param(":id")

	if err := models.DeleteSolicitudCRP(objectID); err == nil {
		j.response = DefaultResponse(200, nil, "delete success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}
