package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// NodoRubroApropiacionController struct del controlador, utiliza los atributos y funciones de un controlador de beego
type NodoRubroApropiacionController struct {
	beego.Controller
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 NodoRubroApropiacion models.NodoRubroApropiacion
// @Failure 403 :objectId is empty
// @router /:vigencia/:unidadEjecutora [get]
func (j *NodoRubroApropiacionController) GetAll() {
	session, _ := db.GetSession()
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	var query = make(map[string]interface{})
	fmt.Println("get all funciton: ", vigencia, unidadEjecutora)
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

	obs := models.GetAllNodoRubroApropiacion(session, query, unidadEjecutora, vigencia)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

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
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = arbolrubroapropiacion
		}
	}
	j.ServeJSON()
}

// Delete elimina
// @Title Delete NodoRubroApropiacion2018
// @Description Borrar NodoRubroApropiacion2018
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *NodoRubroApropiacionController) Delete() {
	session, _ := db.GetSession()
	objectID := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteNodoRubroApropiacionById(session, objectID)
	j.Data["json"] = result
	j.ServeJSON()
}

// Post Método Post de HTTP
// @Title Crear NodoRubroApropiacion2018
// @Description Crear NodoRubroApropiacion2018
// @Param	body		body 	models.NodoRubroApropiacion2018	true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {int} NodoRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router / [post]
func (j *NodoRubroApropiacionController) Post() {
	var arbolrubroapropiacion *models.NodoRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)

	if err := models.InsertNodoRubroApropiacion(arbolrubroapropiacion); err == nil {
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = "error!"
	}

	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update the NodoRubroApropiacion2018
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId/:vigencia/:unidadEjecutora [put]
func (j *NodoRubroApropiacionController) Put() {
	objectID := j.Ctx.Input.Param(":objectId")
	vigencia := j.Ctx.Input.Param(":vigencia")
	unidadEjecutora := j.Ctx.Input.Param(":unidadEjecutora")
	var arbolrubroapropiacion models.NodoRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
	session, _ := db.GetSession()
	vigenciaInt, _ := strconv.Atoi(vigencia)
	err := models.UpdateNodoRubroApropiacion(session, arbolrubroapropiacion, objectID, unidadEjecutora, vigenciaInt)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// Options options
// @Title Preflight options
// @Description Crear NodoRubroApropiacion2018
// @Param	body		body 	models.NodoRubroApropiacion2018	true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {int} NodoRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router / [options]
func (j *NodoRubroApropiacionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// NodoRubroApropiacion2018DeleteOptions NodoRubroApropiacion2018DeleteOptions
// @Title Preflight options
// @Description Crear NodoRubroApropiacion2018
// @Param	body		body 	models.NodoRubroApropiacion2018 true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {int} NodoRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *NodoRubroApropiacionController) NodoRubroApropiacion2018DeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// ArbolApropiacion devuelve un árbol desde la raiz indicada
// @Title Preflight ArbolApropiacion
// @Description Devuelve un nivel del árbol de apropiaciones
// @Param	body		body 	models.NodoRubroApropiacion2018 true		"Body para la creacion de NodoRubroApropiacion2018"
// @Success 200 {object} models.Object
// @Failure 403 body is empty
// @router /ArbolApropiacion/:raiz/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) ArbolApropiacion() {
	nodoRaiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	vigenciastr := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var arbolApropacionessGrande []map[string]interface{}
	vigencia, _ := strconv.Atoi(vigenciastr)
	raiz, err := models.GetNodoApropiacion(session, nodoRaiz, ueStr, vigencia)

	if err == nil {
		arbolApropiaciones := make(map[string]interface{})
		arbolApropiaciones["Id"] = raiz.General.IDPsql
		arbolApropiaciones["Codigo"] = raiz.ID
		arbolApropiaciones["Nombre"] = raiz.General.Nombre
		arbolApropiaciones["IsLeaf"] = true
		arbolApropiaciones["UnidadEjecutora"] = raiz.NodoRubro.UnidadEjecutora
		arbolApropiaciones["ApropiacionInicial"] = raiz.ApropiacionInicial

		var hijos []interface{}
		for j := 0; j < len(raiz.Hijos); j++ {
			hijo := getHijoApropiacion(raiz.Hijos[j], ueStr, vigencia)
			if len(hijo) > 0 {
				arbolApropiaciones["IsLeaf"] = false
				hijos = append(hijos, hijo)
			}
		}
		arbolApropiaciones["Hijos"] = hijos
		arbolApropacionessGrande = append(arbolApropacionessGrande, arbolApropiaciones)

		j.Data["json"] = arbolApropacionessGrande
	} else {
		j.Data["json"] = err
	}

	j.ServeJSON()
}

// RaicesArbolApropiacion
// @Title RaicesArbolApropiacion
// @Description RaicesArbolApropiacion
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /RaicesArbolApropiacion/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) RaicesArbolApropiacion() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var roots []map[string]interface{}
	vigencia, _ := strconv.Atoi(vigenciaStr)
	raices, err := models.GetRaicesApropiacion(session, ueStr, vigencia)
	for i := 0; i < len(raices); i++ {
		idPsql := raices[i].General.IDPsql
		root := map[string]interface{}{
			"Id":                 idPsql,
			"Codigo":             raices[i].ID,
			"Nombre":             raices[i].General.Nombre,
			"Hijos":              raices[i].NodoRubro.Hijos,
			"IsLeaf":             true,
			"UnidadEjecutora":    raices[i].NodoRubro.UnidadEjecutora,
			"ApropiacionInicial": raices[i].ApropiacionInicial,
		}
		if len(raices[i].Hijos) > 0 {
			var hijos []map[string]interface{}
			root["IsLeaf"] = false
			for j := 0; j < len(root["Hijos"].([]string)); j++ {
				hijo := getHijoApropiacion(root["Hijos"].([]string)[j], ueStr, vigencia)
				if len(hijo) > 0 {
					hijos = append(hijos, hijo)
				}
			}
			root["Hijos"] = hijos
		}
		roots = append(roots, root)
	}

	if err != nil {
		j.Data["json"] = err
	} else {
		j.Data["json"] = roots
	}

	j.ServeJSON()
}

// Obtiene y devuelve el nodo hijo de la apropiación, devolviendolo en un objeto tipo json (map[string]interface{})
// Se devuelve un objeto de este tipo y no de models con el fin de utilizar la estructura de json utilizada ya en el cliente
// y no tener que hacer grandes modificaciones en el
func getHijoApropiacion(id, ue string, vigencia int) map[string]interface{} {
	rubroHijo, _ := models.GetNodoRubroApropiacionById(id, ue, vigencia)
	hijo := make(map[string]interface{})
	if rubroHijo != nil {
		if rubroHijo.ID != "" {
			hijo["Id"] = rubroHijo.General.IDPsql
			hijo["Codigo"] = rubroHijo.ID
			hijo["Nombre"] = rubroHijo.General.Nombre
			hijo["IsLeaf"] = false
			hijo["UnidadEjecutora"] = rubroHijo.NodoRubro.UnidadEjecutora
			hijo["ApropiacionInicial"] = rubroHijo.ApropiacionInicial
			if len(rubroHijo.Hijos) == 0 {
				hijo["IsLeaf"] = true
				hijo["Hijos"] = nil
				return hijo
			}
		}
	}

	return hijo
}

// FullArbolRubroApropiaciones ...
// @Title FullArbolRubroApropiaciones
// @Description Construye el árbol a un nivel dependiendo de la raíz
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
		j.Data["json"] = err
		panic(err)
	}

	raizApropiacion, err := models.GetNodoRubroApropiacionById(raiz, ueStr, vigencia)

	tree := rubroApropiacionHelper.BuildTree(raizApropiacion)

	j.Data["json"] = tree
}
