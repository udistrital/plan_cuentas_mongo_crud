package controllers 

import (
	// "encoding/json"
	// "log"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// "github.com/udistrital/plan_cuentas_mongo_crud/db"
	// "github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	// "github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// ProductoController operations for FuenteFinanciamiento
type ProductoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductoController) URLMapping() {
	c.Mapping("Post", c.Post)
}


// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (c *ProductoController) Post() {
	
}