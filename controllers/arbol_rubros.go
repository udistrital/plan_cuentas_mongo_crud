package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/globalsign/mgo" // Inicializa mgo para poder usar sus métodos
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// NodoRubroController estructura para un controlador de beego
type NodoRubroController struct {
	beego.Controller
	response map[string]interface{}
}

func (j *NodoRubroController) URLMapping() {
	j.Mapping("Post", j.Post)
	j.Mapping("Put", j.Put)
	j.Mapping("Delete", j.Delete)
	j.Mapping("Get", j.Get)
	j.Mapping("GetAll", j.GetAll)
	j.Mapping("GetHojas", j.GetHojas)
}

// GetAll función para obtener todos los objetos
// @Title GetAll
// @Description get all objects
// @Param query        query  string    false  "Consulta"
// @Success 200 {object} []models.NodoRubro
// @Failure 403 :objectId is empty
// @router / [get]
func (j *NodoRubroController) GetAll() {
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

	err := errors.New("Bad info response")

	response := DefaultResponse(403, err, nil)

	if obs := models.GetAllNodoRubro(session, query); len(obs) > 0 {
		response = DefaultResponse(200, nil, &obs)
	}

	j.Data["json"] = response

	j.ServeJSON()
}

// Get obtiene un elemento por su id
// @Title Get
// @Description get NodoRubro by nombre
// @Param	id		path 	string	true		"El nombre de la NodoRubro a consultar"
// @Success 200 {object} models.NodoRubro
// @Failure 403 :uid is empty
// @router /:id [get]
func (j *NodoRubroController) Get() {
	id := j.GetString(":id")
	if id != "" {
		arbolrubros, err := models.GetNodoRubroById(id)
		if err == nil {
			j.response = DefaultResponse(200, nil, &arbolrubros)
		} else {
			j.response = DefaultResponse(403, err, nil)
		}
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// @Title Borrar NodoRubro
// @Description Borrar NodoRubro
// @Param	id		path 	string	true		"El id del objeto que se quiere borrar"
// @Success 200 {object} string
// @Failure 403 id is empty
// @router /:id [delete]
func (j *NodoRubroController) Delete() {
	objectID := j.Ctx.Input.Param(":id")
	if err := rubroManager.TrEliminarNodoHoja(objectID, models.NodoRubroCollection); err == nil {
		j.response = DefaultResponse(200, nil, "delete success")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// @Title Crear NodoRubro
// @Description Crear NodoRubro
// @Param	body		body 	models.NodoRubro	true		"Body para la creacion de NodoRubro"
// @Success 200 {object} string
// @Failure 403 body is empty
// @router / [post]
func (j *NodoRubroController) Post() {
	var nodoRubro models.NodoRubro
	json.Unmarshal(j.Ctx.Input.RequestBody, &nodoRubro)
	defer func() {
		if r := recover(); r != nil {
			j.response = DefaultResponse(500, fmt.Errorf(fmt.Sprintf("%s", r)), "insert error!")
		}
		j.Data["json"] = j.response
		j.ServeJSON()
	}()
	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro, models.NodoRubroCollection); err == nil {
		j.response = DefaultResponse(200, nil, "insert success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}

}

// @Title Update
// @Description update the NodoRubro
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.NodoRubro	true		"The body"
// @Success 200 {object} string
// @Failure 403 :id is empty
// @router /:id [put]
func (j *NodoRubroController) Put() {
	objectId := j.Ctx.Input.Param(":id")
	var arbolrubros models.NodoRubro
	json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubros)

	err := models.UpdateNodoRubro(arbolrubros, objectId)
	if err == nil {
		j.response = DefaultResponse(200, nil, "update success!")
	} else {
		j.response = DefaultResponse(403, err, nil)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// @Title FullArbolRubro
// @Description Construye el árbol a un nivel dependiendo de la raíz
// @Param raiz path string true "Código de la raíz"
// @Success 200 {object} []map[string]interface{}
// @Failure 404 body is empty
// @router /arbol/:raiz [get]
func (j *NodoRubroController) FullArbolRubro() {
	raiz := j.GetString(":raiz")

	raizRubro, err := models.GetNodoRubroById(raiz)
	if err != nil {
		j.response = DefaultResponse(403, err, nil)
	} else {
		tree := rubroHelper.BuildTree(&raizRubro)
		j.response = DefaultResponse(200, nil, &tree)
		j.Data["json"] = tree
		j.ServeJSON()
	}

}

func GetHijoRubro(id, ue string) map[string]interface{} {
	session, _ := db.GetSession()
	rubroHijo, _ := models.GetNodo(session, id, ue)
	hijo := make(map[string]interface{})

	if rubroHijo.ID != "" {
		hijo["Codigo"] = rubroHijo.ID
		hijo["Nombre"] = rubroHijo.Nombre
		hijo["IsLeaf"] = false
		hijo["UnidadEjecutora"] = rubroHijo.UnidadEjecutora
		if len(rubroHijo.Hijos) == 0 {
			hijo["IsLeaf"] = true
			hijo["Hijos"] = nil
			return hijo
		}
	}
	return hijo
}

// GetHojas ...
// @Title GetHojas
// @Description Devuelve un arreglo con todos los nodos hoja
// @Success 200 {object} []models.NodoRubroApropiacion
// @router /get_hojas [get]
func (j *NodoRubroController) GetHojas() {
	leafs, err := models.GetHojasRubro()

	if err != nil {
		j.response = DefaultResponse(404, err, nil)
	} else {
		j.response = DefaultResponse(200, nil, &leafs)
	}
	j.Data["json"] = j.response
	j.ServeJSON()
}

// FullArbolRubroReducido ...
// @Title FullArbolRubroReducido
// @Description Construye el árbol con solo el nombre, codigo e hijos a un nivel dependiendo de la raíz y nivel
// @Param raiz  path  string true  "Código de la raíz"
// @Param nivel query string false "Número de nivel (-1 = Todo el arbol, 0 = nivel 0 , 1 = Primer Nivel ... - Por Defecto: -1)"
// @Success 200 {object} []map[string]interface{}
// @Failure 404 body is empty
// @router /arbol_reducido/:raiz [get]
func (j *NodoRubroController) FullArbolRubroReducido() {
	var nivel int
	raiz := j.GetString(":raiz")
	query := j.GetString("nivel")
	if query == "" {
		nivel = -1
	} else {
		nivel, _ = strconv.Atoi(query)
	}
	raizRubro, err := models.GetNodoRubroReducidoById(raiz)
	if err != nil {
		j.response = DefaultResponse(403, err, nil)
	} else {
		tree := rubroHelper.BuildTreeReducido(&raizRubro, nivel)
		j.response = DefaultResponse(200, nil, &tree)
		j.Data["json"] = tree
		j.ServeJSON()
	}

}
