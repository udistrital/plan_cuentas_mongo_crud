package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/compositors"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

type DocumentoPresupuestalController struct {
	beego.Controller
	response map[string]interface{}
}

// GetAllQuery función para obtener todos los objetos con la opción de hacer queries en la BD
// @Title GetAllQuery
// @Description get all objects with data bases query
// @Param vigencia        path  int    true  "Vigencia"
// @Param CG              path  string true  "Centro Gestor (Unidad Ejecutora?)"
// @Param query           query string false "Filter. e.g. col1:v1,col2:v2 ..., if the filter value includes !$ at the begining, the value won't be converted to int"
// @Success 200 {object} []models.DocumentoPresupuestal
// @Failure 403 :objectId is empty
// @router /:vigencia/:CG/ [get]
func (j *DocumentoPresupuestalController) GetAllQuery() {
	var query = make(map[string]interface{})

	vigencia := j.GetString(":vigencia")
	centroGestor := j.GetString(":CG")

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

	obs, err := models.GetAllDocumentoPresupuestal(vigencia, centroGestor, query)
	j.Data["json"] = DefaultResponse(0, err, &obs)
	j.ServeJSON()
}

// Get ...
// Get obtiene un elemento por su id
// @Title Get
// @Description get documento presupuestal by id
// @Param	id		path 	string	true		"El id de la DocumentoPresupuestal a consultar"
// @Param vigencia      path  int    true  "Vigencia"
// @Param areaFuncional path  int    true  "Area Funcional"
// @Success 200 {object} models.DocumentoPresupuestal
// @Failure 403 :objectId is empty
// @router /documento/:vigencia/:areaFuncional/:id [get]
func (j *DocumentoPresupuestalController) Get() {
	objectId := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	areaFuncional := j.GetString(":areaFuncional")

	docPresupuestal, err := models.GetDocumentoPresupuestalById(objectId, vigencia, areaFuncional)
	if err == nil {
		j.response = DefaultResponse(200, nil, &docPresupuestal)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update a documento presupuestal document
// @Param	id			path 	string							true		"The id you want to update"
// @Param vigencia      path  int    true  "Vigencia"
// @Param areaFuncional path  int    true  "Area Funcional"
// @Param	body		body 	models.DocumentoPresupuestal	true		"The body"
// @Success 200 {object} string
// @Failure 403 :id is empty
// @Failure 403 :vigencia is empty
// @Failure 403 :areaFuncional is empty
// @router /:vigencia/:areaFuncional/:id [put]
func (j *DocumentoPresupuestalController) Put() {
	objectID := j.Ctx.Input.Param(":id")
	vigencia := j.Ctx.Input.Param(":vigencia")
	areaFuncional := j.Ctx.Input.Param(":areaFuncional")

	var docPresupuestal models.DocumentoPresupuestal

	json.Unmarshal(j.Ctx.Input.RequestBody, &docPresupuestal)

	if err := models.UpdateDocumentoPresupuestal(&docPresupuestal, objectID, vigencia, areaFuncional); err == nil {
		j.response = DefaultResponse(200, nil, "update success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Param vigencia path  int    true  "Vigencia"
// @Param CG       path  string true  "Centro Gestor (Unidad Ejecutora?)"
// @Param tipo     path  string true  "Tipo"
// @Success 200 {object} []models.DocumentoPresupuestal
// @Failure 403 :objectId is empty
// @router /:vigencia/:CG/:tipo [get]
func (j *DocumentoPresupuestalController) GetAll() {
	vigencia := j.GetString(":vigencia")
	centroGestor := j.GetString(":CG")
	tipo := j.GetString(":tipo")

	rows := documentopresupuestalmanager.GetByType(vigencia, centroGestor, tipo)

	response := commonhelper.DefaultResponse(200, nil, &rows)

	j.Data["json"] = response
	j.ServeJSON()
}

// GetAllCdp función para obtener todos los movimientos de CDP, de una vigencia, sin importar el centro gestor
// @Title GetAllCdp
// @Description get all cdp objects
// @Param vigencia        path  int    true  "Vigencia"
// @Success 200 {object} []models.DocumentoPresupuestal
// @Failure 403 :vigencia is empty
// @router /get_all_cdp/:vigencia [get]
func (j *DocumentoPresupuestalController) GetAllCdp() {
	var rows []models.DocumentoPresupuestal
	vigencia := j.GetString(":vigencia")
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "1", "cdp")...)
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "2", "cdp")...)

	response := commonhelper.DefaultResponse(200, nil, &rows)

	j.Data["json"] = response
	j.ServeJSON()
}

// GetInfoCdp Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Title GetInfoCdp
// @Description Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Param id path string true "_id de Solicitud CDP"
// @Success 200 {object} models.DocumentoPresupuestal
// @Failure 403 :id is empty
// @router /get_info_cdp/:id [get]
func (j *DocumentoPresupuestalController) GetInfoCdp() {
	id := j.GetString(":id")
	data := documentopresupuestalmanager.GetCDPByID(id)

	response := commonhelper.DefaultResponse(200, nil, &data)
	j.Data["json"] = response
	j.ServeJSON()
}

// GetDocMovByParent Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Title GetDocMovByParent
// @Description Obtiene un documento presupuestal de tipo cdp con su id de solicitud
// @Param id       path  string true  "The parentUUID you want to get"
// @Param vigencia path  int    true  "Vigencia"
// @Param CG       path  string true  "Centro Gestor (Unidad Ejecutora?)"
// @Success 200 {object} []models.DocumentoPresupuestal
// @Failure 403 :id is empty
// @router /get_doc_mov_by_parent/:vigencia/:CG/:id [get]
func (j *DocumentoPresupuestalController) GetDocMovByParent() {
	id := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	centroGestor := j.GetString(":CG")
	docPresComp := compositors.DocumentoPresupuestalCompositor{}
	data := docPresComp.GetMovDocumentPresByParent(vigencia, centroGestor, id)

	response := commonhelper.DefaultResponse(200, nil, &data)
	j.Data["json"] = response
	j.ServeJSON()
}

// GetDocMovByRubro ...
// @Title GetDocMovByRubro
// @Description Obtiene todos los CDPs expedidos con movimientos en un rubro padre dado
// @Param	vigencia	path	int	true	"Vigencia del CDP"
// @Param	rubro	path	string	true	"Rubro padre relacionado al CDP"
// @Param areaFuncional path  int    true  "Area Funcional"
// @Success 200 {object} []models.DocumentoPresupuestal Listado de documentos relacionados
// @Failure 500 Internal server error
// @router /get_doc_mov_rubro/:vigencia/:areaFuncional/:rubro [get]
func (j *DocumentoPresupuestalController) GetAllDocMovByRubro() {

	vigencia := j.Ctx.Input.Param(":vigencia")
	rubro := j.Ctx.Input.Param(":rubro")
	centroGestor := j.Ctx.Input.Param(":areaFuncional")

	docPresComp := compositors.DocumentoPresupuestalCompositor{}
	docs, err := docPresComp.GetAllDocumentoPresupuestalMovimientosByRubro(vigencia, centroGestor, rubro)
	j.Data["json"] = DefaultResponse(0, err, &docs)
	j.ServeJSON()
}
