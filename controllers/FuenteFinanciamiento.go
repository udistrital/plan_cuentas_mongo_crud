package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
}

// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (c *FuenteFinanciamientoController) Post() {
	var (
		fuente, infoFuente map[string]interface{}
		movimientosFuente  []map[string]interface{}
		options            []interface{}
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

	json.Unmarshal(c.Ctx.Input.RequestBody, &fuente)

	err = formatdata.FillStruct(fuente["AfectacionFuente"], &movimientosFuente)
	if err != nil {
		panic(err)
	}
	for _, v := range movimientosFuente {
		err := formatdata.FillStruct(v["FuenteFinanciamiento"], &infoFuente)
		if err != nil {
			panic(err)
		}

		valorOriginal := v["ValorOriginal"].(float64)
		valorAcumulado := v["ValorAcumulado"].(float64)
		err, op := crearFuente(infoFuente, valorOriginal, valorAcumulado)
		if err != nil {
			panic(err)
		}
		options = append(options, op)

		// rubroAfecta := map[string]interface{}{
		// 	"Rubro":      v["Rubro"].(string),
		// 	"Dedepencia": int(v["Dependencia"].(float64)),
		// }

		// Afectacion := []map[string]interface{}{rubroAfecta}

		// movimiento := models.Movimiento{
		// 	IDPsql:         "3",
		// 	Afectacion:   Afectacion,
		// 	ValorOriginal:  v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Valor"].(float64),
		// 	Tipo:           "fuente_financiamiento_" + v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["TipoMovimiento"].(map[string]interface{})["Nombre"].(string),
		// 	Vigencia:       "2018",
		// 	DocumentoPadre: strconv.Itoa(int(infoFuente["Id"].(float64))),
		// 	FechaRegistro:  v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Fecha"].(string),
		// }

		// op, err = models.EstrctTransaccionMov(session, &movimiento)
		// if err != nil {
		// 	fmt.Println("Error en estructura de movimiento para fuente de financiamiento")
		// 	panic(err)
		// }
		// options = append(options, op)
	}
	models.TrRegistroFuente(session, options)
}

// crearFuente busca la fuente de financiamiento padre, en caso de que exista, devuelve vacio,
// de lo contario devuelve un objeto de tipo transaccion con la información del registro de la fuente
func crearFuente(informacionFuente map[string]interface{}, valorOriginal float64, valorAcumulado float64) (err error, op interface{}) {
	var tipoFuente string

	session, err := db.GetSession()
	if err != nil {
		return
	}
	defer session.Close()

	fuenteFinanciamiento := models.GetFuenteFinanciamientoByIDPsql(session, int(informacionFuente["Id"].(float64)))
	if fuenteFinanciamiento != nil {
		return
	}

	err = formatdata.FillStructDeep(informacionFuente, "TipoFuenteFinanciamiento.Nombre", &tipoFuente)
	if err != nil { // error convirtiendo a tipo fuente
		panic(err)
	}

	fuenteFinanciamiento = &models.FuenteFinanciamiento{
		ID:             informacionFuente["Codigo"].(string),
		Descripcion:    informacionFuente["Descripcion"].(string),
		IDPsql:         int(informacionFuente["Id"].(float64)),
		Nombre:         informacionFuente["Nombre"].(string),
		TipoFuente:     tipoFuente,
		ValorOriginal:  valorOriginal,
		ValorAcumulado: valorAcumulado,
	}

	op, err = models.PostFuentePadreTransaccion(session, fuenteFinanciamiento)
	if err != nil {
		fmt.Println("Error al crearse estructura de fuente padre")
		panic(err)
	}
	return
}

// GetOne ...
// @Title GetOne
// @Description get FuenteFinanciamiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FuenteFinanciamiento
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FuenteFinanciamientoController) GetOne() {

}
