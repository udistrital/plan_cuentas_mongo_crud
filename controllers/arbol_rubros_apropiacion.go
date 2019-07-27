package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	// "github.com/manucorporat/try"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
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
	session, _ := db.GetSession()
	if id != "" {
		vigenciaInt, _ := strconv.Atoi(vigencia)
		arbolrubroapropiacion, err := models.GetNodoRubroApropiacionById(session, id, unidadEjecutora, vigenciaInt)
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
		j.Data["json"] = err
	}

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
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetNodoRubroApropiacionById(session, id, ue, vigencia)
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

//RegistrarApropiacionInicial
// @Title RegistrarApropiacionInicial...
// @Description Crear NodoRubroApropiacion2018
// @Param	body		body 	models.NodoRubroApropiacion2018 true		"Body para la creacion de ApropiacionInicial"
// @Success 200 {int} NodoRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router RegistrarApropiacionInicial/:vigencia [post]
// func (j *NodoRubroApropiacionController) RegistrarApropiacionInicial() {
// 	var (
// 		dataApropiacion map[string]interface{}
// 		rubro           models.NodoRubro
// 	)
// 	try.This(func() {
// 		vigenciaStr := j.Ctx.Input.Param(":vigencia")
// 		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataApropiacion); err == nil {
// 			session, _ := db.GetSession()

// 			codigoRubro := dataApropiacion["Codigo"].(string)
// 			unidadEjecutora := dataApropiacion["UnidadEjecutora"].(string)
// 			if rubro, err = models.GetNodoRubroById(session, codigoRubro); err != nil {
// 				panic(err.Error())
// 			}
// 			vigencia, _ := strconv.Atoi(vigenciaStr)
// 			general := models.General{
// 				codigoRubro,
// 				vigencia,
// 				dataApropiacion["Nombre"].(string),
// 				"",
// 				int(dataApropiacion["Id"].(float64)),
// 				nil,
// 			}

// 			nodoRubro := models.NodoRubro{
// 				&general,
// 				rubro.Hijos,
// 				rubro.Padre,
// 				dataApropiacion["UnidadEjecutora"].(string),
// 			}

// 			nuevaApropiacion := models.NodoRubroApropiacion{
// 				&nodoRubro,
// 				dataApropiacion["ApropiacionInicial"].(float64),
// 			}

// 			// nuevaApropiacion := models.NodoRubroApropiacion{
// 			// 	&General.ID:         codigoRubro,
// 			// 	Idpsql:              strconv.Itoa(int(dataApropiacion["Id"].(float64))),
// 			// 	Nombre:              dataApropiacion["Nombre"].(string),
// 			// 	Descripcion:         "",
// 			// 	Unidad_ejecutora:    dataApropiacion["UnidadEjecutora"].(string),
// 			// 	Padre:               rubro.Padre,
// 			// 	Hijos:               rubro.Hijos,
// 			// 	Apropiacion_inicial: int(dataApropiacion["ApropiacionInicial"].(float64)),
// 			// }

// 			if nuevaApropiacion.Padre == "" { // Si el rubro actual es una raíz, se hace un registro sencillo
// 				session, _ = db.GetSession()
// 				models.InsertNodoRubroApropiacion(session, &nuevaApropiacion, unidadEjecutora, vigencia)
// 			} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
// 				if err = construirRama(nuevaApropiacion.General.ID, unidadEjecutora, vigencia, nuevaApropiacion.IDPsql, nuevaApropiacion.ApropiacionInicial); err != nil {
// 					fmt.Println("error en construir rama: ", err.Error())
// 					panic(err.Error())
// 				}
// 			}
// 			defer session.Close()
// 			j.Data["json"] = map[string]interface{}{"Type": "success"}
// 		} else {
// 			panic(err.Error())
// 			fmt.Println("unmarshal error: ", err.Error())
// 		}

// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error: ", e)
// 		j.Data["json"] = map[string]interface{}{"Type": "error"}
// 	})

// 	j.ServeJSON()
// }

// Construye la rama a partir de un registro de apropiación inicial
// func construirRama(codigoRubro, ue string, vigencia, idApr int, nuevaApropiacion float64) error {
// 	var (
// 		actualRubro                         models.NodoRubro
// 		padreApropiacion, actualApropiacion *models.NodoRubroApropiacion
// 		err                                 error
// 	)

// 	try.This(func() {
// 		session, _ := db.GetSession()
// 		defer session.Close()
// 		actualRubro, err = models.GetNodoRubroById(session, codigoRubro)
// 		actualRubro.UnidadEjecutora = ue
// 		session, _ = db.GetSession()
// 		padreApropiacion, _ = models.GetNodoRubroApropiacionById(session, actualRubro.Padre, ue, vigencia)

// 		if padreApropiacion == nil {
// 			session, _ = db.GetSession()
// 			actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
// 			models.InsertNodoRubroApropiacion(session, actualApropiacion, ue, vigencia)
// 			if actualApropiacion.Padre != "" {
// 				construirRama(actualRubro.Padre, ue, vigencia, actualRubro.IDPsql, actualApropiacion.ApropiacionInicial)
// 			}
// 		} else {
// 			session, _ = db.GetSession()
// 			apropiacionActualizada, _ := models.GetNodoRubroApropiacionById(session, codigoRubro, ue, vigencia)
// 			apropiacionAnterior := 0.0
// 			session, _ = db.GetSession()
// 			if apropiacionActualizada != nil {
// 				apropiacionAnterior = apropiacionActualizada.ApropiacionInicial
// 				apropiacionActualizada.ApropiacionInicial = nuevaApropiacion
// 				models.UpdateNodoRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.ID, ue, vigencia)
// 			} else {
// 				actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
// 				models.InsertNodoRubroApropiacion(session, actualApropiacion, ue, vigencia)
// 			}

// 			propagarCambio(padreApropiacion.ID, ue, vigencia, nuevaApropiacion-apropiacionAnterior)

// 		}

// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error: ", e)
// 	})
// 	return err
// }

// Propaga el cambio de la apropiación desde la hoja hasta la raiz,
// verificando recursivamente si el rubro que se está obteniendo tiene un padre o no
// func propagarCambio(codigoRubro, ue string, vigencia int, valorPropagado float64) error {
// 	var err error

// 	try.This(func() { // try catch para recibir errores

// 		session, _ := db.GetSession()
// 		apropiacionActualizada, err := models.GetNodoRubroApropiacionById(session, codigoRubro, ue, vigencia)
// 		apropiacionActualizada.ApropiacionInicial += valorPropagado

// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		session, _ = db.GetSession()
// 		models.UpdateNodoRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.ID, ue, vigencia)

// 		if apropiacionActualizada.Padre != "" {
// 			propagarCambio(apropiacionActualizada.Padre, ue, vigencia, valorPropagado)
// 		}
// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error: ", e)
// 		err = errors.New("unknow error")
// 	})
// 	return err
// }

// func crearNuevaApropiacion(actualRubro models.NodoRubro, aprId int, nuevaApropiacion float64) *models.NodoRubroApropiacion {
// 	general := models.General{
// 		actualRubro.ID,
// 		0,
// 		actualRubro.Nombre,
// 		actualRubro.Descripcion,
// 		aprId,
// 		nil,
// 	}

// 	nodoRubro := models.NodoRubro{
// 		&general,
// 		actualRubro.Hijos,
// 		actualRubro.Padre,
// 		actualRubro.UnidadEjecutora,
// 	}

// 	nodoRubroApropiacion := models.NodoRubroApropiacion{
// 		&nodoRubro,
// 		nuevaApropiacion,
// 	}
// 	// actualApropiacion := &models.NodoRubroApropiacion{
// 	// 	Id:                  actualRubro.ID,
// 	// 	Idpsql:              aprId,
// 	// 	Nombre:              actualRubro.Nombre,
// 	// 	Descripcion:         actualRubro.Descripcion,
// 	// 	Unidad_ejecutora:    actualRubro.Unidad_Ejecutora,
// 	// 	Padre:               actualRubro.Padre,
// 	// 	Hijos:               actualRubro.Hijos,
// 	// 	Apropiacion_inicial: nuevaApropiacion,
// 	// }
// 	return &nodoRubroApropiacion
// }

// @Title FullArbolRubroApropiaciones
// @Description Construye el árbol a un nivel dependiendo de la raíz
// @Param body body stringtrue "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /FullArbolRubroApropiaciones/:unidadEjecutora [get]
func (j *NodoRubroApropiacionController) FullArbolRubroApropiaciones() {
	ueStr := j.GetString(":unidadEjecutora")
	fmt.Println(ueStr)
	// tree := rubroHelper.BuildTree(ueStr)

	var tree, childrens []map[string]interface{}

	forkData := make(map[string]interface{})

	//forkData["Codigo"] = "3"

	children := make(map[string]interface{})
	children["data"] = map[string]interface{}{"Codigo": "3-1", "ApropiacionInicial": 500, "children": []map[string]interface{}{
		map[string]interface{}{"data": map[string]interface{}{"Codigo": "3-1-1", "ApropiacionInicial": 300, "children": []map[string]interface{}{}}},
		map[string]interface{}{"data": map[string]interface{}{"Codigo": "3-1-2", "ApropiacionInicial": 200, "children": []map[string]interface{}{}}},
	},
	}

	childrens = append(childrens, children)
	forkData["data"] = map[string]interface{}{"Codigo": 3, "ApropiacionInicial": 500, "children": childrens}
	// forkData["data"]["children"] = childrens

	tree = append(tree, forkData)
	j.Data["json"] = tree
}
