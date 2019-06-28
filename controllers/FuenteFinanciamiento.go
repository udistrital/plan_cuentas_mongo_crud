package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
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

	try.This(func() {
		session, err := db.GetSession()
		if err != nil {
			fmt.Println("error en la sesión")
			panic(err)
		}

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

			valorOriginal := calcularValorOriginal(v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{}))
			err, op := crearFuentePadre(infoFuente, valorOriginal)
			if err != nil {
				panic(err)
			}
			options = append(options, op)

			rubroAfecta := map[string]interface{}{
				"Rubro":      v["Rubro"].(string),
				"Dedepencia": int(v["Dependencia"].(float64)),
			}

			RubrosAfecta := []map[string]interface{}{rubroAfecta}

			movimiento := models.Movimiento{
				IDPsql:         "3",
				RubrosAfecta:   RubrosAfecta,
				ValorOriginal:  v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Valor"].(float64),
				Tipo:           "fuente_financiamiento_" + v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["TipoMovimiento"].(map[string]interface{})["Nombre"].(string),
				Vigencia:       "2018",
				DocumentoPadre: strconv.Itoa(int(infoFuente["Id"].(float64))),
				FechaRegistro:  v["MovimientoFuenteFinanciamientoApropiacion"].([]interface{})[0].(map[string]interface{})["Fecha"].(string),
			}

			op, err = models.EstrctTransaccionMov(session, &movimiento)
			if err != nil {
				fmt.Println("Error en estructura de movimiento para fuente de financiamiento")
				panic(err)
			}
			options = append(options, op)
		}
		err = models.TrRegistroFuente(session, options)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		c.Data["json"] = map[string]interface{}{"Type": "success"}
	}).Catch(func(e try.E) {
		fmt.Println("error en Post() ", e)
		c.Data["json"] = map[string]interface{}{"Type": "error"}
	})
	c.ServeJSON()
}

// crearFuentePadre busca la fuente de financiamiento padre, en caso de que exista, devuelve vacio,
// de lo contario devuelve un objeto de tipo transaccion con la información del registro de la fuente
func crearFuentePadre(informacionFuente map[string]interface{}, valorOriginal float64) (err error, op interface{}) {
	var tipoFuente string

	session, err := db.GetSession()
	if err != nil {
		return
	}
	defer session.Close()

	fuentePadre := models.GetFuenteFinanciamientoPadreByIDPsql(session, int(informacionFuente["Id"].(float64)))
	if fuentePadre != nil {
		return
	}

	err = formatdata.FillStructDeep(informacionFuente, "TipoFuenteFinanciamiento.Nombre", &tipoFuente)
	if err != nil { // error convirtiendo a tipo fuente
		panic(err)
	}

	fuentePadre = &models.FuenteFinaciamientoPadre{
		ID:            informacionFuente["Codigo"].(string),
		Descripcion:   informacionFuente["Descripcion"].(string),
		IDPsql:        int(informacionFuente["Id"].(float64)),
		Nombre:        informacionFuente["Nombre"].(string),
		TipoFuente:    tipoFuente,
		ValorOriginal: valorOriginal,
	}
	op, err = models.EstructaRegistroFuentePadreTransaccion(session, fuentePadre)
	if err != nil {
		fmt.Println("Error al crearse estructura de fuente padre")
		panic(err)
	}
	return
}

// calcularValorOriginal recorre todas las afectaciones de la fuente, sumando el valor y retornando un acumulador del mismo
func calcularValorOriginal(afectaciones []interface{}) (totalFuente float64) {
	for _, v := range afectaciones {
		totalFuente += v.(map[string]interface{})["Valor"].(float64)
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
