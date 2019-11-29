package controllers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	commonhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/commonHelper"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mongo_crud/managers/documentoPresupuestalManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

type DocumentoPresupuestalController struct {
	beego.Controller
}

// GetAllQuery función para obtener todos los objetos con la opción de hacer queries en la BD
// @Title GetAllQuery
// @Description get all objects with data bases query
// @Success 200 DocumentoPresupuestal models.DocumentoPresupuestal
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

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 DocumentoPresupuestal models.DocumentoPresupuestal
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
// @Success 200 rows []models.DocumentoPresupuestal
// @Failure 403 :vigencia is empty
// @router get_all_cdp/:vigencia [get]
func (j *DocumentoPresupuestalController) GetAllCdp() {
	var rows []models.DocumentoPresupuestal
	vigencia := j.GetString(":vigencia")
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "1", "cdp")...)
	rows = append(rows, documentopresupuestalmanager.GetByType(vigencia, "2", "cdp")...)

	response := commonhelper.DefaultResponse(200, nil, &rows)

	j.Data["json"] = response
	j.ServeJSON()
}

// GetInfoCdp Obtiene un documento presupuestal de tipo cdp con su id
// @Title GetInfoCdp
// @Description Obtiene un documento presupuestal de tipo cdp con su id
// @Success 200 documentoPresupuestal models.DocumentoPresupuestal
// @Failure 403 :id is empty
// @router /get_info_cdp/:id [get]
func (j *DocumentoPresupuestalController) GetInfoCdp() {
	id := j.GetString(":id")
	data := documentopresupuestalmanager.GetCDPByID(id)

	response := commonhelper.DefaultResponse(200, nil, &data)
	j.Data["json"] = response
	j.ServeJSON()
}
