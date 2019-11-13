package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"


	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroHelper"
	vigenciahelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/vigenciaHelper"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroApropiacionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// NodoRubroApropiacionController struct del controlador, utiliza los atributos y funciones de un controlador de beego
type NodoRubroApropiacionController struct {
	beego.Controller
	response map[string]interface{}
}

// URLMapping ...
func (j *NodoRubroApropiacionController) URLMapping() {
	j.Mapping("Post", j.Post)
	j.Mapping("Put", j.Put)
	j.Mapping("Delete", j.Delete)
	j.Mapping("Get", j.Get)
	j.Mapping("GetAll", j.GetAll)
	j.Mapping("ArbolApropiacionPadreHijo", j.ArbolApropiacionPadreHijo)
	j.Mapping("RaicesArbolApropiacion", j.RaicesArbolApropiacion)
	j.Mapping("FullArbolRubroApropiaciones", j.FullArbolRubroApropiaciones)
	j.Mapping("FullArbolApropiaciones", j.FullArbolApropiaciones)
	j.Mapping("GetAllVigencia", j.GetAllVigencia)
	j.Mapping("GetHojas", j.GetHojas)
	j.Mapping("AprobacionMasiva", j.AprobacionMasiva)
	j.Mapping("TreeByState", j.TreeByState)
}

// GetAllVigencia función para obtener todos los objetos con una vigencia y una unidad ejecutora
// @Title GetAllVigencia
// @Description get all objects
// @Success 200 NodoRubroApropiacion models.NodoRubroApropiacion
// @Failure 403 :vigencia is empty
// @Failure 403 :unidadEjecutora is empty
// @router /:vigencia/:unidadEjecutora [get]
func (j *NodoRubroApropiacionController) GetAllVigencia() {
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
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

	if obs := models.GetAllNodoRubroApropiacion(query, unidadEjecutora, vigencia); len(obs) > 0 {
		response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = response
	j.ServeJSON()
}

func DefaultResponse(code int, err error, info interface{}) map[string]interface{} {
	response := make(map[string]interface{})

	response["Code"] = code
	response["Message"] = nil
	response["Body"] = info

	if err != nil {
		response["Message"] = err.Error()
		response["Type"] = "error"
	}

	return response
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 NodoRubroApropiacion models.NodoRubroApropiacion
// @Failure 403 :objectId is empty
// @router / [get]
func (j *NodoRubroApropiacionController) GetAll() {
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
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
	response := DefaultResponse(404, err, nil)

	if obs := models.GetAllNodoRubroApropiacion(query, unidadEjecutora, vigencia); len(obs) > 0 {
		response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = response
	j.ServeJSON()
}

// Get Método Get de HTTP
// @Title Get
// @Description get NodoRubroApropiacion2018 by nombre
// @Param	nombre		path 	string	true		"El nombre de la NodoRubroApropiacion2018 a consultar"
// @Success 200 {object} models.NodoRubroApropiacion2018
// @Failure 403 :uid is empty
// @router /:id/:vigencia/:unidadEjecutora [get]
func (j *NodoRubroApropiacionController) Get() {
	id := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	if id != "" {
		vigenciaInt, _ := strconv.Atoi(vigencia)
		arbolrubroapropiacion, err := models.GetNodoRubroApropiacionById(id, unidadEjecutora, vigenciaInt)
		if err == nil {
			j.response = DefaultResponse(200, nil, &arbolrubroapropiacion)
		} else {
			j.response = DefaultResponse(403, err, nil)
		}
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// Delete elimina
// @Title Delete NodoRubroApropiacion2018
// @Description Borrar NodoRubroApropiacion2018
// @Param	id		path 	string	true		"El id del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 id is empty
// @router /:id [delete]
func (j *NodoRubroApropiacionController) Delete() {
	session, _ := db.GetSession()
	objectID := j.Ctx.Input.Param(":id")
	if result, err := models.DeleteNodoRubroApropiacionById(session, objectID); err == nil {
		j.response = DefaultResponse(200, nil, result)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// Post Método Post de HTTP
// @Title Post NodoRubroApropiacion2018
// @Description Post NodoRubroApropiacion2018
// @Param	body		body 	models.NodoRubroApropiacion2018	true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {int} NodoRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router / [post]
func (j *NodoRubroApropiacionController) Post() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			j.Ctx.ResponseWriter.WriteHeader(500)
			j.Data["json"] = r
		}
		j.ServeJSON()
	}()

	var nodoRubroApropiacion *models.NodoRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &nodoRubroApropiacion)

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(nodoRubroApropiacion, nodoRubroApropiacion.UnidadEjecutora, nodoRubroApropiacion.Vigencia); err == nil {
		go vigenciahelper.AddNew(nodoRubroApropiacion.Vigencia, models.ApropiacionVigenciaNameSpace, nodoRubroApropiacion.UnidadEjecutora)
		j.response = DefaultResponse(200, nil, "insert success")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response

}

// Put de HTTP
// @Title Update
// @Description update the NodoRubroApropiacion2018
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :id is empty
// @router /:id/:vigencia/:unidadEjecutora [put]
func (j *NodoRubroApropiacionController) Put() {
	objectID := j.Ctx.Input.Param(":id")
	vigencia := j.Ctx.Input.Param(":vigencia")
	unidadEjecutora := j.Ctx.Input.Param(":unidadEjecutora")
	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			j.Ctx.ResponseWriter.WriteHeader(500)
			j.Data["json"] = r
		}
		j.ServeJSON()
	}()
	var arbolrubroapropiacion models.NodoRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
	vigenciaInt, _ := strconv.Atoi(vigencia)
	err := rubroApropiacionManager.TrActualizarValorApropiacion(&arbolrubroapropiacion, objectID, unidadEjecutora, vigenciaInt)
	if err == nil {
		j.response = DefaultResponse(200, nil, "update success")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// ArbolApropiacionPadreHijo devuelve un árbol desde la raiz indicada
// @Title Preflight ArbolApropiacionPadreHijo
// @Description Devuelve un nivel del árbol de apropiaciones
// @Param	body		body 	models.NodoRubroApropiacion2018 true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {object} models.Object
// @Failure 403 body is empty
// @router /ArbolApropiacionPadreHijo/:raiz/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) ArbolApropiacionPadreHijo() {
	nodoRaiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	vigenciastr := j.GetString(":vigencia")
	var arbolApropacionessGrande []map[string]interface{}

	vigencia, _ := strconv.Atoi(vigenciastr)
	raiz, err := models.GetNodoApropiacion(nodoRaiz, ueStr, vigencia)

	if err == nil {
		arbolApropiaciones := make(map[string]interface{})
		arbolApropiaciones["Codigo"] = raiz.ID
		arbolApropiaciones["Nombre"] = raiz.General.Nombre
		arbolApropiaciones["IsLeaf"] = true
		arbolApropiaciones["UnidadEjecutora"] = raiz.NodoRubro.UnidadEjecutora
		arbolApropiaciones["ValorInicial"] = raiz.ValorInicial

		var hijos []interface{}
		for j := 0; j < len(raiz.Hijos); j++ {
			hijo := rubroApropiacionHelper.GetHijoApropiacion(raiz.Hijos[j], ueStr, vigencia)
			if len(hijo) > 0 {
				arbolApropiaciones["IsLeaf"] = false
				hijos = append(hijos, hijo)
			}
		}
		arbolApropiaciones["Hijos"] = hijos
		arbolApropacionessGrande = append(arbolApropacionessGrande, arbolApropiaciones)

		j.response = DefaultResponse(200, nil, arbolApropacionessGrande)
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// RaicesArbolApropiacion ...
// @Title RaicesArbolApropiacion
// @Description RaicesArbolApropiacion
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /RaicesArbolApropiacion/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) RaicesArbolApropiacion() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")

	var roots []map[string]interface{}
	vigencia, _ := strconv.Atoi(vigenciaStr)
	raices, err := models.GetRaicesApropiacion(ueStr, vigencia)
	for i := 0; i < len(raices); i++ {
		root := map[string]interface{}{
			"Codigo":          raices[i].ID,
			"Nombre":          raices[i].General.Nombre,
			"Hijos":           raices[i].NodoRubro.Hijos,
			"IsLeaf":          true,
			"UnidadEjecutora": raices[i].NodoRubro.UnidadEjecutora,
			"ValorInicial":    raices[i].ValorInicial,
		}
		if len(raices[i].Hijos) > 0 {
			var hijos []map[string]interface{}
			root["IsLeaf"] = false
			for j := 0; j < len(root["Hijos"].([]string)); j++ {
				hijo := rubroApropiacionHelper.GetHijoApropiacion(root["Hijos"].([]string)[j], ueStr, vigencia)
				if len(hijo) > 0 {
					hijos = append(hijos, hijo)
				}
			}
			root["Hijos"] = hijos
		}
		roots = append(roots, root)
	}

	if err == nil {
		j.response = DefaultResponse(200, nil, &roots)
	} else {
		j.response = DefaultResponse(404, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// FullArbolRubroApropiaciones ...
// @Title FullArbolRubroApropiaciones
// @Description Construye el árbol dependiendo de la raíz
// @Param body body stringtrue "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /arbol_apropiacion/:raiz/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) FullArbolRubroApropiaciones() {
	raiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")

	vigencia, err := strconv.Atoi(vigenciaStr)
	if err != nil {
		j.response = DefaultResponse(404, err, nil)
		panic(err)
	} else {
		raizApropiacion, _ := models.GetNodoRubroApropiacionById(raiz, ueStr, vigencia)
		tree := rubroApropiacionHelper.BuildTree(raizApropiacion)
		j.response = DefaultResponse(200, nil, &tree)

	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// FullArbolApropiaciones ...
// @Title FullArbolApropiaciones
// @Description Construye el árbol completo con valores
// @Param body body stringtrue "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /arbol_apropiacion_valores/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) FullArbolApropiaciones() {
	unidadEjecutora := j.GetString(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")

	vigencia, err := strconv.Atoi(vigenciaStr)

	if err != nil {
		j.response = DefaultResponse(404, err, nil)
	} else {
		tree := rubroApropiacionHelper.ValuesTree(unidadEjecutora, vigencia, "")
		j.response = DefaultResponse(200, nil, &tree)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// GetHojas ...
// @Title GetHojas
// @Description Devuelve un arreglo con todos los nodos hoja correspondientes a la vigencia y ue
// @Param body body string	true "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /get_hojas/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) GetHojas() {
	unidadEjecutora := j.GetString(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")

	leafs, err := models.GetHojasApropiacion(unidadEjecutora, vigencia)

	if err != nil {
		j.response = DefaultResponse(404, err, nil)
	} else {
		j.response = DefaultResponse(200, nil, &leafs)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// ComprobarBalanceArbolApropiaciones ...
// @Title ComprobarBalanceArbolApropiaciones
// @Description ComprobarBalanceArbolApropiaciones
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /comprobar_balance/:unidadEjecutora/:vigencia [post]
func (j *NodoRubroApropiacionController) ComprobarBalanceArbolApropiaciones() {

	response := make(map[string]interface{})

	var (
		movimientos []models.Movimiento
	)

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			responseformat.SetResponseFormat(&j.Controller, r, "", 500)
		}
		responseformat.SetResponseFormat(&j.Controller, response, "", 200)

	}()

	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")

	vigencia, _ := strconv.Atoi(vigenciaStr)
	raices, err := models.GetRaicesApropiacion(ueStr, vigencia)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(j.Ctx.Input.RequestBody, &movimientos)

	balance := make(map[string]map[string]interface{})
	if len(movimientos) > 0 {
		balance = rubroApropiacionHelper.SimulatePropagationValues(movimientos, vigenciaStr, ueStr)
	}

	var rootCompValue float64
	values := make(map[string]models.NodoRubroApropiacion)
	balanceado := true
	approved := false
	rootsParamsIndexed := rubroHelper.GetRubroParamsIndexedByKey(ueStr, "Valor")

	for _, raiz := range raices {
		if rootsParamsIndexed[raiz.ID] != nil {
			values[raiz.ID] = raiz
			if raiz.Estado == models.EstadoAprobada {
				rootsAprpovedTotal++
			}
		}
		// perform this operation only if there are some simulation to perform ...
		if rootsParamsIndexed[raiz.ID] != nil && balance[raiz.ID] != nil {
			actualRaiz := raiz
			actualRaiz.ValorActual = balance[raiz.ID]["valor_actual"].(float64)
			values[raiz.ID] = actualRaiz
		}
	}
	var indexValue int
	for _, rootValue := range values {
		if indexValue == 0 {
			rootCompValue = rootValue.ValorActual
		}
		if rootCompValue != rootValue.ValorActual || rootValue.ValorActual == 0 {
			balanceado = false
		}
		if rubroInfo, e := rubroManager.SearchRubro(rootValue.ID, ueStr); e {
			if rubroInfo.ID == "2" {
				response["totalIngresos"] = rootValue.ValorActual
			} else if rubroInfo.ID == "3" {
				response["totalGastos"] = rootValue.ValorActual
			}

		}
		indexValue++
	}

	if rootCompValue == 0 {
		balanceado = false
	}

	if response["totalGastos"] == nil || response["totalIngresos"] == nil {
		balanceado = false
	}

	// if the tree's roots are all approved then the whole tree is approved..
	if rootsAprpovedTotal == len(rootsParamsIndexed) {
		approved = true
	}

	response["balanceado"] = balanceado
	response["approved"] = approved


}

// AprobacionMasiva ...
// @Title AprobacionMasiva
// @Description Cambia el estado de los documentos arbol_rubro_apropiacion de una vigencia y unidad ejecutora
// @Param unidadEjecutora unidadEjecutora string	true "Unidad Ejecutora de la apropiación"
// @Param vigencia vigencia string	true "Vigencia de la apropiación"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /aprobacion_masiva/:unidadEjecutora/:vigencia [post]
func (j *NodoRubroApropiacionController) AprobacionMasiva() {
	unidadEjecutora := j.GetString(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")

	if err := rubroApropiacionManager.TrAprobarApropiaciones(unidadEjecutora, vigencia); err == nil {
		j.response = DefaultResponse(200, nil, "Ok")
	} else {
		j.response = DefaultResponse(404, err, nil)
	}

	j.Data["json"] = j.response
	j.ServeJSON()
}

// TreeByState ...
// @Title TreeByState
// @Description Genera el árbol dependiendo del estado de las apropiaciones
// @Param unidadEjecutora unidadEjecutora string	true "Unidad Ejecutora de la apropiación"
// @Param vigencia vigencia string	true "Vigencia de la apropiación"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /arbol_por_estado/:unidadEjecutora/:vigencia/:estado [get]
func (j *NodoRubroApropiacionController) TreeByState() {
	unidadEjecutora := j.GetString(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")
	estado := j.GetString(":estado")
	vigencia, err := strconv.Atoi(vigenciaStr)
	if err != nil {
		j.response = DefaultResponse(404, err, nil)
	} else {
		tree := rubroApropiacionHelper.ValuesTree(unidadEjecutora, vigencia, estado)
		j.response = DefaultResponse(200, nil, &tree)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}
