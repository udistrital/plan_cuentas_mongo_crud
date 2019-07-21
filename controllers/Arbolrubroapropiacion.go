package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// ArbolRubroApropiacionController struct del controlador, utiliza los atributos y funciones de un controlador de beego
type ArbolRubroApropiacionController struct {
	beego.Controller
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Success 200 ArbolRubroApropiacion models.ArbolRubroApropiacion
// @Failure 403 :objectId is empty
// @router /:vigencia/:unidadEjecutora [get]
func (j *ArbolRubroApropiacionController) GetAll() {
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

	obs := models.GetAllArbolRubroApropiacion(session, query, unidadEjecutora, vigencia)

	if len(obs) == 0 {
		j.Data["json"] = []string{}
	} else {
		j.Data["json"] = &obs
	}

	j.ServeJSON()
}

// Get Método Get de HTTP
// @Title Get
// @Description get ArbolRubroApropiacion2018 by nombre
// @Param	nombre		path 	string	true		"El nombre de la ArbolRubroApropiacion2018 a consultar"
// @Success 200 {object} models.ArbolRubroApropiacion2018
// @Failure 403 :uid is empty
// @router /:id/:vigencia/:unidadEjecutora [get]
func (j *ArbolRubroApropiacionController) Get() {
	id := j.GetString(":id")
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	session, _ := db.GetSession()
	if id != "" {
		arbolrubroapropiacion, err := models.GetArbolRubroApropiacionById(session, id, unidadEjecutora, vigencia)
		if err != nil {
			j.Data["json"] = err.Error()
		} else {
			j.Data["json"] = arbolrubroapropiacion
		}
	}
	j.ServeJSON()
}

// Delete elimina
// @Title Delete ArbolRubroApropiacion2018
// @Description Borrar ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *ArbolRubroApropiacionController) Delete() {
	session, _ := db.GetSession()
	objectID := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteArbolRubroApropiacion2018ById(session, objectID)
	j.Data["json"] = result
	j.ServeJSON()
}

// Post Método Post de HTTP
// @Title Crear ArbolRubroApropiacion2018
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018	true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:vigencia/:unidadEjecutora [post]
func (j *ArbolRubroApropiacionController) Post() {
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	if vigencia != "" {
		var arbolrubroapropiacion *models.ArbolRubroApropiacion
		json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
		fmt.Println(arbolrubroapropiacion)
		session, _ := db.GetSession()
		models.InsertArbolRubroApropiacion(session, arbolrubroapropiacion, unidadEjecutora, vigencia)
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = "vigencia null"
	}

	j.ServeJSON()
}

// Put de HTTP
// @Title Update
// @Description update the ArbolRubroApropiacion2018
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId/:vigencia/:unidadEjecutora [put]
func (j *ArbolRubroApropiacionController) Put() {
	objectID := j.Ctx.Input.Param(":objectId")
	vigencia := j.Ctx.Input.Param(":vigencia")
	unidadEjecutora := j.Ctx.Input.Param(":unidadEjecutora")
	var arbolrubroapropiacion models.ArbolRubroApropiacion
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
	session, _ := db.GetSession()

	err := models.UpdateArbolRubroApropiacion(session, arbolrubroapropiacion, objectID, unidadEjecutora, vigencia)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}

// Options options
// @Title Preflight options
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018	true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router / [options]
func (j *ArbolRubroApropiacionController) Options() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// ArbolRubroApropiacion2018DeleteOptions ArbolRubroApropiacion2018DeleteOptions
// @Title Preflight options
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router /:objectId [options]
func (j *ArbolRubroApropiacionController) ArbolRubroApropiacion2018DeleteOptions() {
	j.Data["json"] = "success!"
	j.ServeJSON()
}

// ArbolApropiacion devuelve un árbol desde la raiz indicada
// @Title Preflight ArbolApropiacion
// @Description Devuelve un nivel del árbol de apropiaciones
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ArbolRubroApropiacion2018"
// @Success 200 {object} models.Object
// @Failure 403 body is empty
// @router /ArbolApropiacion/:raiz/:unidadEjecutora/:vigencia [get]
func (j *ArbolRubroApropiacionController) ArbolApropiacion() {
	nodoRaiz := j.GetString(":raiz")
	ueStr := j.GetString(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var arbolApropacionessGrande []map[string]interface{}

	raiz, err := models.GetNodoApropiacion(session, nodoRaiz, ueStr, vigencia)

	if err == nil {
		arbolApropiaciones := make(map[string]interface{})
		arbolApropiaciones["Codigo"] = raiz.Id
		arbolApropiaciones["Nombre"] = raiz.Nombre
		arbolApropiaciones["IsLeaf"] = true
		arbolApropiaciones["UnidadEjecutora"] = raiz.Unidad_ejecutora
		arbolApropiaciones["ApropiacionInicial"] = raiz.Apropiacion_inicial

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
func (j *ArbolRubroApropiacionController) RaicesArbolApropiacion() {
	ueStr := j.Ctx.Input.Param(":unidadEjecutora")
	vigencia := j.GetString(":vigencia")
	session, _ := db.GetSession()
	var roots []map[string]interface{}
	raices, err := models.GetRaicesApropiacion(session, ueStr, vigencia)
	for i := 0; i < len(raices); i++ {
		root := map[string]interface{}{
			"Codigo":             raices[i].Id,
			"Nombre":             raices[i].Nombre,
			"Hijos":              raices[i].Hijos,
			"IsLeaf":             true,
			"UnidadEjecutora":    raices[i].Unidad_ejecutora,
			"ApropiacionInicial": raices[i].Apropiacion_inicial,
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
func getHijoApropiacion(id, ue, vigencia string) map[string]interface{} {
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetArbolRubroApropiacionById(session, id, ue, vigencia)
	hijo := make(map[string]interface{})
	if rubroHijo != nil {
		if rubroHijo.Id != "" {
			hijo["Codigo"] = rubroHijo.Id
			hijo["Nombre"] = rubroHijo.Nombre
			hijo["IsLeaf"] = false
			hijo["UnidadEjecutora"] = rubroHijo.Unidad_ejecutora
			hijo["ApropiacionInicial"] = rubroHijo.Apropiacion_inicial
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
// @Description Crear ArbolRubroApropiacion2018
// @Param	body		body 	models.ArbolRubroApropiacion2018 true		"Body para la creacion de ApropiacionInicial"
// @Success 200 {int} ArbolRubroApropiacion2018.Id
// @Failure 403 body is empty
// @router RegistrarApropiacionInicial/:vigencia [post]
func (j *ArbolRubroApropiacionController) RegistrarApropiacionInicial() {
	var (
		dataApropiacion map[string]interface{}
		rubro           models.ArbolRubros
	)
	session, err := db.GetSession()
	if err != nil {
		fmt.Println("error en la sesión")
		panic(err)
	}
	defer func() {
		session.Close()
		if r := recover(); r != nil {
			logs.Error(r)
			logManager.LogError(r)
			panic(r)
		}
	}()
	vigencia := j.Ctx.Input.Param(":vigencia")
	if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataApropiacion); err == nil {
		codigoRubro := dataApropiacion["Codigo"].(string)
		unidadEjecutora := dataApropiacion["UnidadEjecutora"].(string)
		if rubro, err = models.GetArbolRubrosById(session, codigoRubro); err != nil {
			panic(err.Error())
		}

		nuevaApropiacion := models.ArbolRubroApropiacion{
			Id:                  codigoRubro,
			Nombre:              dataApropiacion["Nombre"].(string),
			Descripcion:         "",
			Unidad_ejecutora:    dataApropiacion["UnidadEjecutora"].(string),
			Padre:               rubro.Padre,
			Hijos:               rubro.Hijos,
			Estado:              dataApropiacion["Estado"].(string),
			Apropiacion_inicial: int(dataApropiacion["ApropiacionInicial"].(float64)),
		}

		if nuevaApropiacion.Padre == "" { // Si el rubro actual es una raíz, se hace un registro sencillo
			session, _ = db.GetSession()
			models.InsertArbolRubroApropiacion(session, &nuevaApropiacion, unidadEjecutora, vigencia)
		} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
			if err = construirRama(nuevaApropiacion.Id, unidadEjecutora, vigencia, nuevaApropiacion.Id, nuevaApropiacion.Apropiacion_inicial); err != nil {
				fmt.Println("error en construir rama: ", err.Error())
				panic(err.Error())
			}
		}

		j.Data["json"] = map[string]interface{}{"Type": "success"}
	} else {
		panic(err.Error())
		fmt.Println("unmarshal error: ", err.Error())
	}

	j.ServeJSON()
}

// Construye la rama a partir de un registro de apropiación inicial
func construirRama(codigoRubro, ue, vigencia, idApr string, nuevaApropiacion int) error {
	var (
		actualRubro                         models.ArbolRubros
		padreApropiacion, actualApropiacion *models.ArbolRubroApropiacion
		err                                 error
	)

	session, err := db.GetSession()
	if err != nil {
		fmt.Println("error en la sesión")
		panic(err)
	}
	defer session.Close()
	actualRubro, err = models.GetArbolRubrosById(session, codigoRubro)
	actualRubro.Unidad_Ejecutora = ue
	session, _ = db.GetSession()
	padreApropiacion, _ = models.GetArbolRubroApropiacionById(session, actualRubro.Padre, ue, vigencia)

	if padreApropiacion == nil {
		session, _ = db.GetSession()
		actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
		models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
		if actualApropiacion.Padre != "" {
			construirRama(actualRubro.Padre, ue, vigencia, actualRubro.Id, actualApropiacion.Apropiacion_inicial)
		}
	} else {
		session, _ = db.GetSession()
		apropiacionActualizada, _ := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
		apropiacionAnterior := 0
		session, _ = db.GetSession()
		if apropiacionActualizada != nil {
			apropiacionAnterior = apropiacionActualizada.Apropiacion_inicial
			apropiacionActualizada.Apropiacion_inicial = nuevaApropiacion
			models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia)
		} else {
			actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
			models.InsertArbolRubroApropiacion(session, actualApropiacion, ue, vigencia)
		}

		propagarCambio(padreApropiacion.Id, ue, vigencia, nuevaApropiacion-apropiacionAnterior)

	}

	return err
}

// Propaga el cambio de la apropiación desde la hoja hasta la raiz,
// verificando recursivamente si el rubro que se está obteniendo tiene un padre o no
func propagarCambio(codigoRubro, ue, vigencia string, valorPropagado int) error {
	var err error

	session, err := db.GetSession()
	if err != nil {
		fmt.Println("error en la sesión")
		panic(err)
	}
	defer session.Close()
	apropiacionActualizada, err := models.GetArbolRubroApropiacionById(session, codigoRubro, ue, vigencia)
	apropiacionActualizada.Apropiacion_inicial += valorPropagado

	if err != nil {
		panic(err.Error())
	}
	session, _ = db.GetSession()
	models.UpdateArbolRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.Id, ue, vigencia)

	if apropiacionActualizada.Padre != "" {
		propagarCambio(apropiacionActualizada.Padre, ue, vigencia, valorPropagado)
	}
	return err
}

func crearNuevaApropiacion(actualRubro models.ArbolRubros, aprId string, nuevaApropiacion int) *models.ArbolRubroApropiacion {
	actualApropiacion := &models.ArbolRubroApropiacion{
		Id:                  actualRubro.Id,
		Nombre:              actualRubro.Nombre,
		Descripcion:         actualRubro.Descripcion,
		Unidad_ejecutora:    actualRubro.Unidad_Ejecutora,
		Padre:               actualRubro.Padre,
		Hijos:               actualRubro.Hijos,
		Estado:              "Aprobado",
		Apropiacion_inicial: nuevaApropiacion,
	}
	return actualApropiacion
}

// @Title FullArbolRubroApropiaciones
// @Description Construye el árbol a un nivel dependiendo de la raíz
// @Param body body stringtrue "Código de la raíz"
// @Success 200 {object} models.Object
// @Failure 404 body is empty
// @router /FullArbolRubroApropiaciones/:unidadEjecutora [get]
func (j *ArbolRubroApropiacionController) FullArbolRubroApropiaciones() {
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
