package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	// "log"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"

	// "github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// ProductoController operations for FuenteFinanciamiento
type ProductoController struct {
	beego.Controller
}

// URLMapping ...
func (j *ProductoController) URLMapping() {
	j.Mapping("Get", j.Get)
	j.Mapping("GetAll", j.GetAll)
	j.Mapping("Post", j.Post)
	j.Mapping("Delete", j.Delete)
}

// GetAll
// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 Producto models.Producto
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ProductoController) GetAll() {
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

	obs := models.GetAllProducto(query)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// Get ...
// Get obtiene un elemento por su id
// @Title Get
// @Description get Producto by nombre
// @Param	nombre		path 	string	true		"El nombre de la Producto a consultar"
// @Success 200 {object} models.Producto
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ProductoController) Get() {
	id := j.GetString(":id")
	fmt.Println(id)
	if id != "" {
		producto, err := models.GetProductoById(id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = producto
		}
	}
	j.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Producto
// @Param	body		body 	models.Producto	true		"body for Producto content"
// @Success 201 {object} models.Producto
// @Failure 403 body is empty
// @router / [post]
func (j *ProductoController) Post() {
	var producto models.Producto
	json.Unmarshal(j.Ctx.Input.RequestBody, &producto)

	if err := models.InsertProducto(producto); err != nil {
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = err.Error()
	}
}

// Delete ...
// @Title Eliminar Producto
// @Description Eliminar Producto
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ProductoController) Delete() {
	objectId := j.Ctx.Input.Param(":objectId")
	err := models.DeleteProductoById(objectId)
	if err == nil {
		j.Data["json"] = "delete success!"
	} else {
		j.Data["json"] = err.Error()
	}
}

// Put ...
// @Title Update
// @Description update the Producto
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ProductoController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var producto models.Producto
	json.Unmarshal(j.Ctx.Input.RequestBody, &producto)

	if err := models.UpdateProducto(producto, objectId); err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
}
