package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/globalsign/mgo" // Inicializa mgo para poder usar sus métodos
	"github.com/manucorporat/try"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// ArbolRubrosController estructura para un controlador de beego
type ArbolRubrosController struct {
	beego.Controller
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 ArbolRubros models.ArbolRubros
// @Failure 403 :objectId is empty
// @router / [get]
func (j *ArbolRubrosController) GetAll() {
	session, _ := db.GetSession()

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

	obs := models.GetAllArbolRubross(session, query)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// Get obtiene un elemento por su id
// @Title Get
// @Description get ArbolRubros by nombre
// @Param	nombre		path 	string	true		"El nombre de la ArbolRubros a consultar"
// @Success 200 {object} models.ArbolRubros
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *ArbolRubrosController) Get() {
	id := j.GetString(":id")
	session, _ := db.GetSession()
	if id != "" {
		arbolrubros, err := models.GetArbolRubrosById(session, id)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = arbolrubros
		}
	}
	j.ServeJSON()
}

// @Title Borrar ArbolRubros
// @Description Borrar ArbolRubros
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ArbolRubrosController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteArbolRubrosById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear ArbolRubros
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros	true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router / [post]
func (j *ArbolRubrosController) Post() {
	var arbolrubros models.ArbolRubros
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubros)
	fmt.Println(arbolrubros)
	session, _ := db.GetSession()
	models.InsertArbolRubros(session, arbolrubros)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the ArbolRubros
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *ArbolRubrosController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var arbolrubros models.ArbolRubros
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubros)
	session, _ := db.GetSession()

	err := models.UpdateArbolRubros(session, arbolrubros, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros	true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ArbolRubrosController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Preflight options
// @Description Crear ArbolRubros
// @Param	body		body 	models.ArbolRubros true		"Body para la creacion de ArbolRubros"
// @Success 200 {int} ArbolRubros.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ArbolRubrosController) ArbolRubrosDeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// @Title Registra rubro
// @Description Convierte la estructura del api mid para registrarla en un documento mongo
// @Param	body		interface{}	true		"Body para la creacion de un nuevo rubro"
// @Success 201 {string} recibido
// @Failure 403 body is empty
// @router /registrarRubro [post]
func (j *ArbolRubrosController) RegistrarRubro() {
	try.This(func() {
		var (
			rubroData  interface{}
			rubroPadre string = ""
			err        error
		)
		session, _ := db.GetSession()
		json.Unmarshal(j.Ctx.Input.RequestBody, &rubroData)
		rubroDataHijo := rubroData.(map[string]interface{})["RubroHijo"].(map[string]interface{})
		nuevoRubro := models.ArbolRubros{
			Id:               rubroDataHijo["Codigo"].(string),
			Idpsql:           strconv.FormatFloat(rubroDataHijo["Id"].(float64), 'f', 0, 64),
			Nombre:           rubroDataHijo["Nombre"].(string),
			Descripcion:      rubroDataHijo["Descripcion"].(string),
			Hijos:            nil,
			Unidad_Ejecutora: strconv.FormatFloat(rubroDataHijo["UnidadEjecutora"].(float64), 'f', 0, 64)}

		if rubroData.(map[string]interface{})["RubroPadre"] != nil {
			rubroDataPadre := rubroData.(map[string]interface{})["RubroPadre"].(map[string]interface{})
			rubroPadre = rubroDataPadre["Codigo"].(string)
			nuevoRubro.Padre = rubroPadre
			updatedRubro, _ := models.GetArbolRubrosById(session, rubroPadre)
			updatedRubro.Hijos = append(updatedRubro.Hijos, rubroDataHijo["Codigo"].(string))
			session, _ = db.GetSession()
			err = models.RegistrarRubroTransacton(updatedRubro, nuevoRubro, session)
		} else {
			err = models.InsertArbolRubros(session, nuevoRubro)
		}

		if err != nil {
			panic(err.Error())
		} else {
			j.Data["json"] = map[string]interface{}{"Type": "success"}
		}
	}).Catch(func(e try.E) {
		fmt.Println(e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})

	j.ServeJSON()
}

// @Title Eliminar rubro
// @Description recibe el idPsql del rubro desde api mid para eliminar el rubro
// @Param	body		interface{}	true		"Body para la eliminación de un rubro"
// @Success 201 {string} sucess
// @Failure 403 body is empty
// @router /eliminarRubro/:idPsql [delete]
func (j *ArbolRubrosController) EliminarRubro() {
	try.This(func() {
		session, _ := db.GetSession()
		var (
			// rubroIdPsql interface{}
			err error
		)
		rubroIdPsql := j.GetString(":idPsql")
		rubroHijo, _ := models.GetArbolRubrosByIdPsql(session, rubroIdPsql)
		fmt.Println("rubroHijo: ", rubroHijo)
		session, _ = db.GetSession()
		if rubroHijo.Padre != "" {
			rubroPadre, _ := models.GetArbolRubrosById(session, rubroHijo.Padre)
			fmt.Println("rubroPadre sin modificar: ", rubroPadre)
			rubroPadre.Hijos = remove(rubroPadre.Hijos, rubroHijo.Id)
			fmt.Println("rubroPadre modificado: ", rubroPadre)
			session, _ = db.GetSession()
			err = models.EliminarRubroTransaccion(rubroPadre, rubroHijo, session)
		} else {
			_, err = models.DeleteArbolRubrosById(session, rubroHijo.Id)
		}

		if err != nil {
			panic(err.Error())
		} else {
			j.Data["json"] = map[string]interface{}{"Type": "success"}
		}

	}).Catch(func(e try.E) {
		fmt.Println(e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	j.ServeJSON()
}

func remove(slice []string, object string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == object {
			slice = append(slice[:i], slice[i+1:]...)
			return slice
		}
	}
	return slice
}

// @Title RaicesArbol
// @Description RaicesArbol
// @Param body body models.Rubro true "Body para la creacion de Rubro"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /RaicesArbol/:unidadEjecutora [get]
func (j *ArbolRubrosController) RaicesArbol() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	session, _ := db.GetSession()
	var roots []map[string]interface{}
	rubros, err := models.GetRaices(session, ueStr)
	for i := 0; i < len(rubros); i++ {
		idPsql, _ := strconv.Atoi(rubros[i].Idpsql)
		root := map[string]interface{}{
			"Id":              idPsql,
			"Codigo":          rubros[i].Id,
			"Nombre":          rubros[i].Nombre,
			"Hijos":           rubros[i].Hijos,
			"IsLeaf":          true,
			"UnidadEjecutora": rubros[i].Unidad_Ejecutora,
		}
		if len(rubros[i].Hijos) > 0 {
			var hijos []map[string]interface{}
			root["IsLeaf"] = false
			for j := 0; j < len(root["Hijos"].([]string)); j++ {
				hijo := GetHijoRubro(root["Hijos"].([]string)[j], ueStr)
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

// @Title Preflight options
// @Description Construye el árbol a un nivel dependiendo de la raíz
// @Param body body stringtrue "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /ArbolRubro/:raiz/:unidadEjecutora [get]
func (j *ArbolRubrosController) ArbolRubro() {
	nodoRaiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	session, _ := db.GetSession()
	var arbolRubrosGrande []map[string]interface{}
	raiz, err := models.GetNodo(session, nodoRaiz, ueStr)
	if err == nil {

		arbolRubros := make(map[string]interface{})
		arbolRubros["Id"], _ = strconv.Atoi(raiz.Idpsql)
		arbolRubros["Codigo"] = raiz.Id
		arbolRubros["Nombre"] = raiz.Nombre
		arbolRubros["IsLeaf"] = true
		arbolRubros["UnidadEjecutora"] = raiz.Unidad_Ejecutora
		var hijos []interface{}
		for j := 0; j < len(raiz.Hijos); j++ {
			hijo := GetHijoRubro(raiz.Hijos[j], ueStr)
			if len(hijo) > 0 {
				arbolRubros["IsLeaf"] = false
				hijos = append(hijos, hijo)
			}
		}
		arbolRubros["Hijos"] = hijos
		arbolRubrosGrande = append(arbolRubrosGrande, arbolRubros)

		j.Data["json"] = arbolRubrosGrande
	} else {
		j.Data["json"] = err
	}

	j.ServeJSON()
}

func GetHijoRubro(id, ue string) map[string]interface{} {
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetNodo(session, id, ue)
	hijo := make(map[string]interface{})

	if rubroHijo.Id != "" {
		hijo["Id"], _ = strconv.Atoi(rubroHijo.Idpsql)
		hijo["Codigo"] = rubroHijo.Id
		hijo["Nombre"] = rubroHijo.Nombre
		hijo["IsLeaf"] = false
		hijo["UnidadEjecutora"] = rubroHijo.Unidad_Ejecutora
		if len(rubroHijo.Hijos) == 0 {
			hijo["IsLeaf"] = true
			hijo["Hijos"] = nil
			return hijo
		}
	}
	return hijo
}
