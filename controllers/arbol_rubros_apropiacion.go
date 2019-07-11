package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
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
// @router /:vigencia/:unidadEjecutora [post]
func (j *NodoRubroApropiacionController) Post() {
	vigencia := j.GetString(":vigencia")
	unidadEjecutora := j.GetString(":unidadEjecutora")
	if vigencia != "" {
		var arbolrubroapropiacion *models.NodoRubroApropiacion
		json.Unmarshal(j.Ctx.Input.RequestBody, &arbolrubroapropiacion)
		fmt.Println(arbolrubroapropiacion)
		session, _ := db.GetSession()
		vigenciaInt, _ := strconv.Atoi(vigencia)
		models.InsertNodoRubroApropiacion(session, arbolrubroapropiacion, unidadEjecutora, vigenciaInt)
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = "vigencia null"
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
		arbolApropiaciones["Codigo"] = raiz.General.ID
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
			"Codigo":             raices[i].General.ID,
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
		if rubroHijo.General.ID != "" {
			hijo["Id"] = rubroHijo.General.IDPsql
			hijo["Codigo"] = rubroHijo.General.ID
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
func (j *NodoRubroApropiacionController) RegistrarApropiacionInicial() {
	var (
		dataApropiacion map[string]interface{}
		rubro           models.NodoRubro
	)
	try.This(func() {
		vigenciaStr := j.Ctx.Input.Param(":vigencia")
		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataApropiacion); err == nil {
			session, _ := db.GetSession()

			codigoRubro := dataApropiacion["Codigo"].(string)
			unidadEjecutora := dataApropiacion["UnidadEjecutora"].(string)
			if rubro, err = models.GetNodoRubroById(session, codigoRubro); err != nil {
				panic(err.Error())
			}
			vigencia, _ := strconv.Atoi(vigenciaStr)
			general := models.General{
				codigoRubro,
				vigencia,
				dataApropiacion["Nombre"].(string),
				"",
				int(dataApropiacion["Id"].(float64)),
				nil,
			}

			nodoRubro := models.NodoRubro{
				&general,
				rubro.Hijos,
				rubro.Padre,
				dataApropiacion["UnidadEjecutora"].(string),
			}

			nuevaApropiacion := models.NodoRubroApropiacion{
				&nodoRubro,
				dataApropiacion["ApropiacionInicial"].(float64),
			}

			// nuevaApropiacion := models.NodoRubroApropiacion{
			// 	&General.ID:         codigoRubro,
			// 	Idpsql:              strconv.Itoa(int(dataApropiacion["Id"].(float64))),
			// 	Nombre:              dataApropiacion["Nombre"].(string),
			// 	Descripcion:         "",
			// 	Unidad_ejecutora:    dataApropiacion["UnidadEjecutora"].(string),
			// 	Padre:               rubro.Padre,
			// 	Hijos:               rubro.Hijos,
			// 	Apropiacion_inicial: int(dataApropiacion["ApropiacionInicial"].(float64)),
			// }

			if nuevaApropiacion.Padre == "" { // Si el rubro actual es una raíz, se hace un registro sencillo
				session, _ = db.GetSession()
				models.InsertNodoRubroApropiacion(session, &nuevaApropiacion, unidadEjecutora, vigencia)
			} else { // si el rubro actual no es una raíz, se itera para registrar toda la rama
				if err = construirRama(nuevaApropiacion.General.ID, unidadEjecutora, vigencia, nuevaApropiacion.IDPsql, nuevaApropiacion.ApropiacionInicial); err != nil {
					fmt.Println("error en construir rama: ", err.Error())
					panic(err.Error())
				}
			}
			defer session.Close()
			j.Data["json"] = map[string]interface{}{"Type": "success"}
		} else {
			panic(err.Error())
			fmt.Println("unmarshal error: ", err.Error())
		}

	}).Catch(func(e try.E) {
		fmt.Println("catch error: ", e)
		j.Data["json"] = map[string]interface{}{"Type": "error"}
	})

	j.ServeJSON()
}

// Construye la rama a partir de un registro de apropiación inicial
func construirRama(codigoRubro, ue string, vigencia, idApr int, nuevaApropiacion float64) error {
	var (
		actualRubro                         models.NodoRubro
		padreApropiacion, actualApropiacion *models.NodoRubroApropiacion
		err                                 error
	)

	try.This(func() {
		session, _ := db.GetSession()
		defer session.Close()
		actualRubro, err = models.GetNodoRubroById(session, codigoRubro)
		actualRubro.UnidadEjecutora = ue
		session, _ = db.GetSession()
		padreApropiacion, _ = models.GetNodoRubroApropiacionById(session, actualRubro.Padre, ue, vigencia)

		if padreApropiacion == nil {
			session, _ = db.GetSession()
			actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
			models.InsertNodoRubroApropiacion(session, actualApropiacion, ue, vigencia)
			if actualApropiacion.Padre != "" {
				construirRama(actualRubro.Padre, ue, vigencia, actualRubro.IDPsql, actualApropiacion.ApropiacionInicial)
			}
		} else {
			session, _ = db.GetSession()
			apropiacionActualizada, _ := models.GetNodoRubroApropiacionById(session, codigoRubro, ue, vigencia)
			apropiacionAnterior := 0.0
			session, _ = db.GetSession()
			if apropiacionActualizada != nil {
				apropiacionAnterior = apropiacionActualizada.ApropiacionInicial
				apropiacionActualizada.ApropiacionInicial = nuevaApropiacion
				models.UpdateNodoRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.ID, ue, vigencia)
			} else {
				actualApropiacion = crearNuevaApropiacion(actualRubro, idApr, nuevaApropiacion)
				models.InsertNodoRubroApropiacion(session, actualApropiacion, ue, vigencia)
			}

			propagarCambio(padreApropiacion.ID, ue, vigencia, nuevaApropiacion-apropiacionAnterior)

		}

	}).Catch(func(e try.E) {
		fmt.Println("catch error: ", e)
	})
	return err
}

// Propaga el cambio de la apropiación desde la hoja hasta la raiz,
// verificando recursivamente si el rubro que se está obteniendo tiene un padre o no
func propagarCambio(codigoRubro, ue string, vigencia int, valorPropagado float64) error {
	var err error

	try.This(func() { // try catch para recibir errores

		session, _ := db.GetSession()
		apropiacionActualizada, err := models.GetNodoRubroApropiacionById(session, codigoRubro, ue, vigencia)
		apropiacionActualizada.ApropiacionInicial += valorPropagado

		if err != nil {
			panic(err.Error())
		}
		session, _ = db.GetSession()
		models.UpdateNodoRubroApropiacion(session, *apropiacionActualizada, apropiacionActualizada.ID, ue, vigencia)

		if apropiacionActualizada.Padre != "" {
			propagarCambio(apropiacionActualizada.Padre, ue, vigencia, valorPropagado)
		}
	}).Catch(func(e try.E) {
		fmt.Println("catch error: ", e)
		err = errors.New("unknow error")
	})
	return err
}

func crearNuevaApropiacion(actualRubro models.NodoRubro, aprId int, nuevaApropiacion float64) *models.NodoRubroApropiacion {
	general := models.General{
		actualRubro.ID,
		0,
		actualRubro.Nombre,
		actualRubro.Descripcion,
		aprId,
		nil,
	}

	nodoRubro := models.NodoRubro{
		&general,
		actualRubro.Hijos,
		actualRubro.Padre,
		actualRubro.UnidadEjecutora,
	}

	nodoRubroApropiacion := models.NodoRubroApropiacion{
		&nodoRubro,
		nuevaApropiacion,
	}
	// actualApropiacion := &models.NodoRubroApropiacion{
	// 	Id:                  actualRubro.ID,
	// 	Idpsql:              aprId,
	// 	Nombre:              actualRubro.Nombre,
	// 	Descripcion:         actualRubro.Descripcion,
	// 	Unidad_ejecutora:    actualRubro.Unidad_Ejecutora,
	// 	Padre:               actualRubro.Padre,
	// 	Hijos:               actualRubro.Hijos,
	// 	Apropiacion_inicial: nuevaApropiacion,
	// }
	return &nodoRubroApropiacion
}

var tipoTotal string
var tipoMovimiento string
var tipoMovimientoPadre string

//@Title SaldoCDP
//

//RegistrarMovimiento ...
// @Title RegistrarMovimiento
// @Description Registra los movimientos (como cdp, rp, ver variable tipoMovimiento) y los propaga tanto en la colección
// arbolrubrosapropiacion_[vigencia]_[unidad_ejecutura], como en la colección movimientos. Utiliza la función registrarValores para registrar los valores,
// y se le envian como párametro el nombre de los movimientos que se van a guardar en el atributo movimiento de la colección arbolrubrosapropiacion,
// al igual que se envia la variable dataValor, que son los valores del movimiento enviados desde el api_mid_financiera
// @Param	body		body 	models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router RegistrarMovimiento/:tipoPago [post]
// func (j *ArbolRubroApropiacionController) RegistrarMovimiento() {
// 	var dataValor map[string]interface{}

// 	try.This(func() {

// 		if err := json.Unmarshal(j.Ctx.Input.RequestBody, &dataValor); err != nil {
// 			panic(err.Error())
// 		}

// 		switch tipoMovimiento = j.GetString(":tipoPago"); tipoMovimiento {
// 		//rp
// 		case "Cdp":
// 			tipoTotal = "TotalComprometidoCdp"
// 			tipoMovimientoPadre = "Apr"
// 			registrarValores(dataValor, "total_cdp", "mes_cdp")
// 		case "Rp":
// 			tipoTotal = "TotalComprometidoRp"
// 			tipoMovimientoPadre = "Cdp"
// 			registrarValores(dataValor, "total_rp", "mes_rp")
// 		case "AnulacionRp":
// 			tipoTotal = "TotalAnuladoRp"
// 			tipoMovimientoPadre = "Rp"
// 			registrarValores(dataValor, "total_anulado_rp", "mes_anulado_rp")
// 		case "AnulacionCdp":
// 			tipoTotal = "TotalAnuladoCdp"
// 			tipoMovimientoPadre = "Cdp"
// 			registrarValores(dataValor, "total_anulado_cdp", "mes_anulado_cdp")
// 		case "Adicion": //Adición a la apropiación inicial
// 			tipoTotal = "AdicionApr"
// 			tipoMovimientoPadre = ""
// 			registrarValores(dataValor, "total_adicion", "mes_modificacion")
// 		case "ModificacionApr": // traslado de apropiación
// 			fmt.Println("Modificación de apropiación.....")
// 			registrarModifacionApr(dataValor)
// 		}

// 		j.Data["json"] = map[string]interface{}{"Type": "success"}
// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error registrar movimiento: ", e)
// 		j.Data["json"] = map[string]interface{}{"Type": "error"}
// 	})
// 	j.ServeJSON()
// }

// // De acuerdo a los valores que recibe, se hacen las modificaciones en el arbolrubroapropiacion
// // y también en la colección de movimientos
// // Parámetros: Recibe los valores correspondientes a la modificación, el mes correspondiente de la modificaicón
// func registrarModifacionApr(dataValor map[string]interface{}) (err error) {
// 	var ops []interface{}

// 	try.This(func() {
// 		unidadEjecutora := strconv.Itoa(int(dataValor["UnidadEjecutora"].(float64)))
// 		fechaRegistro := dataValor["FechaMovimiento"].(string)
// 		vigencia := strconv.Itoa(int(dataValor["Vigencia"].(float64)))
// 		mes, _ := time.Parse("2006-01-02", fechaRegistro)

// 		opsApr := registrarValoresModf(dataValor["Afectacion"].([]interface{}), strconv.Itoa(int(mes.Month())), vigencia, unidadEjecutora)

// 		for _, v := range dataValor["Afectacion"].([]interface{}) {
// 			value := v.(map[string]interface{}) // Convierte el elemento v en un map[string]inerface{}, para evitar una conversión constante del mismo

// 			tipoMovimiento := value["TipoMovimiento"].(string)

// 			modificacionApr := models.Movimiento{
// 				IDPsql:          strconv.Itoa(int(dataValor["Id"].(float64))),
// 				Tipo:            tipoMovimiento,
// 				Vigencia:        vigencia,
// 				DocumentoPadre:  strconv.Itoa(int(value["Apropiacion"].(float64))),
// 				FechaRegistro:   fechaRegistro,
// 				UnidadEjecutora: unidadEjecutora,
// 			}
// 			modificacionApr.RubrosAfecta = append(modificacionApr.RubrosAfecta, value)

// 			session, _ := db.GetSession()
// 			op, err := models.EstrctTransaccionMov(session, &modificacionApr)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			ops = append(ops, op)

// 		}
// 		ops = append(ops, opsApr...)
// 		for i := range ops {
// 			fmt.Println(ops[i], "\n......")
// 		}
// 		session, _ := db.GetSession()
// err = models.RegistrarMovimiento(session, ops)
// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error registrar modificación apropiación")
// 		panic(e)
// 	})
// 	return err
// }

// Crea un CDP para las modificaciones de apropiación inicial que lo necesitan
// func crearCdp(dataMovimiento map[string]interface{}, unidadEjecutora, fechaRegistro, vigencia string) (op interface{}) {
// 	var err error // error handle variable

// 	try.This(func() {
// 		rubrosAfecta := make(map[string]interface{})
// 		rubrosAfecta["Rubro"] = dataMovimiento["CuentaCredito"].(string)
// 		if dataMovimiento["TipoMovimiento"] == "Traslado" {
// 			rubrosAfecta["Rubro"] = dataMovimiento["CuentaContraCredito"].(string)
// 		}

// 		rubrosAfecta["Valor"] = dataMovimiento["Valor"].(float64)
// 		rubrosAfecta["Apropiacion"] = strconv.Itoa(int(dataMovimiento["Apropiacion"].(float64)))
// 		cdp := models.Movimiento{
// 			IDPsql:          strconv.Itoa(int(dataMovimiento["Disponibilidad"].(float64))),
// 			Tipo:            "Cdp",
// 			Vigencia:        vigencia,
// 			DocumentoPadre:  "0",
// 			FechaRegistro:   fechaRegistro,
// 			UnidadEjecutora: unidadEjecutora,
// 		}
// 		cdp.RubrosAfecta = append(cdp.RubrosAfecta, rubrosAfecta)

// 		session, _ := db.GetSession()
// 		op, err = models.EstrctTransaccionMov(session, &cdp)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error en crearCdp")
// 		panic(e)
// 	})
// 	return
// }

// // Registrar valores de modificación en arbolrubroapropiacion
// func registrarValoresModf(dataModificacion []interface{}, mes, vigencia, ue string) (ops []interface{}) {
// 	// var err error
// 	nuevoValor := make(map[string]map[string]map[string]float64)

// 	try.This(func() {

// 		for _, d := range dataModificacion {
// 			data := d.(map[string]interface{})
// 			data["Mes"] = mes
// 			if nuevoValor[data["CuentaCredito"].(string)] == nil {
// 				nuevoValor[data["CuentaCredito"].(string)] = make(map[string]map[string]float64)
// 			}

// 			if nuevoValor[data["CuentaCredito"].(string)][mes] == nil {
// 				nuevoValor[data["CuentaCredito"].(string)][mes] = make(map[string]float64)
// 			}

// 			if data["CuentaContraCredito"].(string) != "" && nuevoValor[data["CuentaContraCredito"].(string)] == nil {
// 				nuevoValor[data["CuentaContraCredito"].(string)] = make(map[string]map[string]float64)
// 			}

// 			if data["CuentaContraCredito"].(string) != "" && nuevoValor[data["CuentaContraCredito"].(string)][mes] == nil {
// 				nuevoValor[data["CuentaContraCredito"].(string)][mes] = make(map[string]float64)
// 			}

// 			if data["TipoMovimiento"].(string) != "Traslado" {

// 				formatModifGeneral(data, nuevoValor)
// 			} else {
// 				formatModifTraslado(data, nuevoValor)
// 			}
// 			beego.Debug(data["TipoMovimiento"], data["CuentaContraCredito"].(string), nuevoValor)

// 		}
// 		for k, v := range nuevoValor {
// 			op, err := prograpacionValores(k, mes, vigencia, ue, v[mes])
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			ops = append(ops, op...)
// 		}
// 	}).Catch(func(e try.E) {
// 		beego.Error("catch error en registrarValoresModificaciones")
// 		panic(e)
// 	})
// 	return
// }

// Formatea las modificaciones de tipo: reducción, adición, suspensión
func formatModifGeneral(data map[string]interface{}, res map[string]map[string]map[string]float64) {
	// if res[data["CuentaContraCredito"].(string)] == nil {
	// 	res[data["CuentaContraCredito"].(string)] = make(map[string]map[string]float64)

	// 	if res[data["CuentaContraCredito"].(string)][data["Mes"].(string)] == nil {
	// 		res[data["CuentaContraCredito"].(string)][data["Mes"].(string)] = make(map[string]float64)
	// 	}

	// 	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_credito"] += data["Valor"].(float64)
	// 	res[data["CuentaContraCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_contra_credito"] += data["Valor"].(float64)
	// } else {
	// 	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)] += data["Valor"].(float64)
	// }

	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)] += data["Valor"].(float64)

}

// Formatea las modificaciones de traslado
func formatModifTraslado(data map[string]interface{}, res map[string]map[string]map[string]float64) {
	res[data["CuentaCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_credito"] += data["Valor"].(float64)
	res[data["CuentaContraCredito"].(string)][data["Mes"].(string)][data["TipoMovimiento"].(string)+"_cuenta_contra_credito"] += data["Valor"].(float64)

}

// Itera sobre cada uno de los objetos que estén en el atributo Afectacion enviado desde el api_mid_financiera, que tienen la información necesaria del movimiento.
// Mientras se itera en cada uno de los elementos, se crean las variable rubro, unidadEjecutora y vigencia, para que se pueda buscar el nodo correspondiente en
// la colección arbolrubrosapropiacion_[vigencia]_[unidadEjecutora], luego se comprueba si dicho nodo tiene movimientosAsociados a el. En caso de no tener ninguno
// se instancia un nuevo atributo para que tenga esos valores, luego se guardan los valores enviados desde el api_mid_finciera en la variable nuevoValor y se envian
// como parametro para la función propagarValores, la cuál propaga los valores en el arbolrubrosapropiaciones, devolviendo un arrreglo de interfaces op
// Para la transacción que se llevará acabo
// func registrarValores(dataValor map[string]interface{}, total, mes string) (err error) {
// 	try.This(func() {

// 		var (
// 			op  []interface{} // operación para la transacción
// 			ops []interface{} // todas las operaciones de la transacción
// 		)

// 		for _, v := range dataValor["Afectacion"].([]interface{}) {
// 			rubro := v.(map[string]interface{})["Rubro"].(string)
// 			unidadEjecutora := v.(map[string]interface{})["UnidadEjecutora"].(string)
// 			vigencia := dataValor["Vigencia"].(string)

// 			session, _ := db.GetSession()

// 			rubroApropiacion, err := models.GetArbolRubroApropiacionById(session, rubro, unidadEjecutora, vigencia)

// 			if err != nil {
// 				panic(err.Error())
// 			}

// 			nuevoValor := make(map[string]float64)

// 			if len(rubroApropiacion.Movimientos) == 0 {
// 				rubroApropiacion.Movimientos = make(map[string]map[string]float64)
// 				rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] = make(map[string]float64)
// 			}

// 			if rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] == nil {
// 				rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)] = make(map[string]float64)
// 			}

// 			nuevoValor[mes] = v.(map[string]interface{})["Valor"].(float64)
// 			nuevoValor[total] = v.(map[string]interface{})["Valor"].(float64)

// 			rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][mes] = v.(map[string]interface{})["Valor"].(float64)
// 			rubroApropiacion.Movimientos[dataValor["MesRegistro"].(string)][total] += v.(map[string]interface{})["Valor"].(float64)

// 			ops, err = prograpacionValores(rubroApropiacion.Id, dataValor["MesRegistro"].(string), vigencia, unidadEjecutora, nuevoValor)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 		}

// 		op, err = registrarDocumentoMovimiento(dataValor, total, mes)

// 		ops = append(ops, op...)

// 		session, _ := db.GetSession()
// 		models.RegistrarMovimiento(session, ops)
// 	}).Catch(func(e try.E) {
// 		fmt.Println("catch error registrar valores: ", e)
// 		panic(e)
// 	})
// 	return err
// }

// func registrarDocumentoMovimiento(dataValor map[string]interface{}, total, mes string) (ops []interface{}, err error) {
// 	try.This(func() {
// 		var rubrosAfecta []map[string]interface{}

// 		documentoPadre, _ := dataValor["Disponibilidad"].(float64)

// 		for _, rubroAfecta := range dataValor["Afectacion"].([]interface{}) {
// 			rubrosAfecta = append(rubrosAfecta, rubroAfecta.(map[string]interface{}))
// 		}
// 		movimiento := models.Movimiento{
// 			IDPsql:         strconv.Itoa(int(dataValor["Id"].(float64))),
// 			RubrosAfecta:   rubrosAfecta,
// 			Tipo:           tipoMovimiento,
// 			Vigencia:       dataValor["Vigencia"].(string),
// 			DocumentoPadre: strconv.Itoa(int(documentoPadre)), // si el documento padre esta vacio (no tiene) el valor guardado es 0 (?)
// 		}
// 		session, _ := db.GetSession()
// 		op, err := models.EstrctTransaccionMov(session, &movimiento)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		ops = append(ops, op)

// 		opp, err := propagarValorMovimientos(movimiento.DocumentoPadre, movimiento, tipoMovimiento) // opp son los movimientos a propagar en la tx de mongodb
// 		ops = append(ops, opp...)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}).Catch(func(e try.E) {
// 		fmt.Println("error en registrar RP ", e)
// 		panic(e)
// 	})
// 	return ops, err
// }

// H
// func propagarValorMovimientos(documentoPadre string, Rp models.Movimiento, tMovimiento string) (op []interface{}, err error) {
// 	session, _ := db.GetSession()
// 	selectTipoMovimientoPadre(tMovimiento)
// 	padre, _ := models.GetMovimientoByPsqlId(session, documentoPadre, tipoMovimientoPadre)

// 	if padre != nil {
// 		afectacionWalk(&Rp, padre)

// 		session, _ = db.GetSession()
// 		opM, err := models.EstrctUpdateTransaccionMov(session, padre) //opM es la tx del movimiento a actualizar
// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		op = append(op, opM)
// 		opp, err := propagarValorMovimientos(padre.DocumentoPadre, Rp, tipoMovimientoPadre)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		op = append(op, opp...)

// 	}

// 	return
// }

// afectacionWalk itera en todos los elementos de RubrosAfecta del apuntador rp, y luego itera en todos los elementos de RubroAfecta del apuntador RP
// hasta encontrar los elementos que el movimiento de tipo RP afecta al CDP, en este punto se pueden dar dos acciones dependiendo de la variable tipoTotal:
// 1. Que el movimiento aún no tenga el atributo tipoTotal registrado
// 2. Que el movimiento ya tenga el atributo tipoTotal registrado
// Si el movimiento aún no tiene registrado el atributo tipoTotal, se crea y se le asigna el valor que viene del RP (el cual afectaria al CDP)
// Si el movimiento ya tiene registrado el atributo tipoTotal, se modifica su valor sumándole el que viene del RP (el cual afectaria a los correspondientes rubros del CDP)
// Finalmente los apuntadores son modificados y continuan su proceso en la función propagarValorMovimientos
func afectacionWalk(Rp, Cdp *models.Movimiento) {
	for _, rubroRp := range Rp.RubrosAfecta {
		for i := 0; i < len(Cdp.RubrosAfecta); i++ {
			if Cdp.RubrosAfecta[i]["Rubro"].(string) == rubroRp["Rubro"].(string) {
				if Cdp.RubrosAfecta[i][tipoTotal] != nil {
					switch Cdp.RubrosAfecta[i][tipoTotal].(type) {
					case int:
						Cdp.RubrosAfecta[i][tipoTotal] = float64(Cdp.RubrosAfecta[i][tipoTotal].(int)) + rubroRp["Valor"].(float64)
					case float64:
						Cdp.RubrosAfecta[i][tipoTotal] = Cdp.RubrosAfecta[i][tipoTotal].(float64) + rubroRp["Valor"].(float64)
					default:
						panic("No suitable type assertion")
					}

				} else {
					Cdp.RubrosAfecta[i][tipoTotal] = rubroRp["Valor"].(float64)
				}
			}
		}
	}

}

func prograpacionValores(rubro, mes, ue string, vigencia int, valorPrograpado map[string]float64) (ops []interface{}, err error) {
	try.This(func() {

		session, _ := db.GetSession()

		apropiacionPadre, err := models.GetNodoRubroApropiacionById(session, rubro, ue, vigencia)

		var apropiacionesCdp []*models.NodoRubroApropiacion
		if err != nil {
			panic(err.Error())
		}

		for apropiacionPadre != nil {
			if apropiacionPadre.Movimientos[mes] == nil {
				apropiacionPadre.Movimientos[mes] = make(map[string]float64)
			}

			if len(apropiacionPadre.Movimientos) == 0 {
				apropiacionPadre.Movimientos = make(map[string]map[string]float64)
				apropiacionPadre.Movimientos[mes] = valorPrograpado
			} else {
				for key, value := range valorPrograpado {

					if apropiacionPadre.Movimientos[mes][key] != 0 {
						if strings.Contains(key, "mes") {
							apropiacionPadre.Movimientos[mes][key] = value

						} else {
							apropiacionPadre.Movimientos[mes][key] += value
						}
					} else {
						apropiacionPadre.Movimientos[mes][key] = value
					}
				}
			}

			apropiacionesCdp = append(apropiacionesCdp, apropiacionPadre)

			if apropiacionPadre.Padre != "" {
				session, _ = db.GetSession()
				apropiacionPadre, err = models.GetNodoRubroApropiacionById(session, apropiacionPadre.Padre, ue, vigencia)
			} else {
				apropiacionPadre = nil
			}

			if err != nil {
				panic(err.Error())
			}

		}
		session, _ = db.GetSession()
		options, err := models.EstrctTransaccionArbolApropiacion(session, apropiacionesCdp, ue, vigencia)
		if err != nil {
			panic(err.Error())
		}
		for _, obj := range options {
			ops = append(ops, obj)
		}

	}).Catch(func(e try.E) {
		fmt.Println("catch error prograpacionValores: ", e)
		panic(e)
	})

	return ops, err
}

func selectTipoMovimientoPadre(tipoHijo string) {
	switch tipoMovimiento = tipoHijo; tipoMovimiento {
	//rp
	case "Cdp":
		tipoMovimientoPadre = "Apr"
	case "Rp":
		tipoMovimientoPadre = "Cdp"
	case "AnulacionRp":
		tipoMovimientoPadre = "Rp"
	case "AnulacionCdp":
		tipoMovimientoPadre = "Cdp"
	default:
		tipoMovimientoPadre = ""
	}
}

// SaldoApropiacion ...
// @Title SaldoApropiacion
// @Description Devuelve el saldo de una apropiación especifica
// @Param	body		body 	models.Object true "json de movimientos enviado desde el api_mid_financiera"
// @Success 200 {string} success
// @Failure 403 error
// @router /SaldoApropiacion/:rubro/:unidadEjecutora/:vigencia [get]
func (j *NodoRubroApropiacionController) SaldoApropiacion() {
	try.This(func() {
		var (
			rubroParam    string
			unidadEParam  int
			vigenciaParam int
			err           error
		)
		response := make(map[string]float64)
		rubroParam = j.GetString(":rubro")
		if unidadEParam, err = j.GetInt(":unidadEjecutora"); err != nil {
			panic(err.Error())
		}

		if vigenciaParam, err = j.GetInt(":vigencia"); err != nil {
			panic(err.Error())
		}

		session, _ := db.GetSession()
		rubro, err := models.GetNodoRubroApropiacionById(session, rubroParam, strconv.Itoa(unidadEParam), vigenciaParam)

		for _, value := range rubro.Movimientos {
			for key, data := range value {
				response[key] += data
			}
		}
		response["valor_inicial"] = float64(rubro.ApropiacionInicial)
		if err != nil {
			panic(err.Error())
		}

		j.Data["json"] = response
	}).Catch(func(e try.E) {
		j.Data["json"] = e
	})

	j.ServeJSON()
}

// SaldoMovimiento Devuelve un objeto con el saldo del cdp
// @Title SaldoMovimiento
// @Description Devuelve el saldo de un CDP especifico
// @Param	idPsql		path 	int	true		"idPsql del documento"
// @Param	rubro		path 	string	true		"código del rubro"
// @Param	fuente		query	string false		"fuente de financiamiento"
// @Success 200 {string} success
// @Failure 403 error
// @router /SaldoMovimiento/:idPsql/:rubro/:tipoMovimiento [get]
func (j *NodoRubroApropiacionController) SaldoMovimiento() {
	try.This(func() {
		var (
			cdpID    int
			err      error
			response map[string]interface{}
		)

		cdpID, err = j.GetInt(":idPsql") // id psql del cdp
		if err != nil {
			panic(err.Error())
		}
		rubro := j.GetString(":rubro")
		fuente := j.GetString("fuente")
		tipoMovimiento := j.GetString(":tipoMovimiento")
		session, _ := db.GetSession()
		cdp, err := models.GetMovimientoByPsqlId(session, strconv.Itoa(cdpID), tipoMovimiento)
		if err != nil {
			panic(err.Error())
		}

		for _, value := range cdp.RubrosAfecta {
			if value["FuenteCodigo"] == nil && value["Rubro"].(string) == rubro && fuente == "" {
				response = value
			} else if value["Rubro"].(string) == rubro && value["FuenteCodigo"].(string) == fuente {
				response = value
			}
		}
		delete(response, "FuenteNombre")
		delete(response, "Rubro")
		delete(response, "Apropiacion")
		delete(response, "FuenteCodigo")
		delete(response, "UnidadEjecutora")

		j.Data["json"] = response
	}).Catch(func(e try.E) {
		j.Data["json"] = e
	})

	j.ServeJSON()
}
