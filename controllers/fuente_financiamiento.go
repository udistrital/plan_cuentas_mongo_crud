package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (c *FuenteFinanciamientoController) Post() {
	var (
		fuente  models.FuenteFinanciamiento
		options []interface{}
	)

	session, err := db.GetSession()
	if err != nil {
		log.Panicln(err.Error())
	}

	defer func() {
		if r := recover(); r != nil {
			session.Close()
			logs.Error(r)
			logManager.LogError(r)
			panic(r)
		}
	}()

	json.Unmarshal(c.Ctx.Input.RequestBody, &fuente)
	log.Println(fuente)

	op, err := models.PostFuentePadreTransaccion(session, &fuente)

	options = append(options, op)
	models.TrRegistroFuente(session, options)
}
