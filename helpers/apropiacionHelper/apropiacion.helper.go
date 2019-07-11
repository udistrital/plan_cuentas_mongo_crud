package apropiacionHelper

import (
	"fmt"
	"strings"

	"github.com/manucorporat/try"
	"github.com/udistrital/plan_cuentas_mongo_crud/db"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// PropagacionValores ... Update the apropiation's tree balance over the data passed to this function as "valorPropagado"
func PropagacionValores(rubro, mes, vigencia, ue string, valorPrograpado map[string]float64) (ops []interface{}, err error) {
	try.This(func() {
		session, _ := db.GetSession()

		apropiacionPadre, err := models.GetArbolRubroApropiacionById(session, rubro, ue, vigencia)

		var apropiacionesCdp []*models.ArbolRubroApropiacion
		if err != nil {
			panic(err.Error())
		}

		for apropiacionPadre != nil {

			if len(apropiacionPadre.Movimientos) == 0 {
				apropiacionPadre.Movimientos = make(map[string]map[string]float64)
				apropiacionPadre.Movimientos[mes] = make(map[string]float64)
				apropiacionPadre.Movimientos[mes] = valorPrograpado
			} else {
				if apropiacionPadre.Movimientos[mes] == nil {
					apropiacionPadre.Movimientos[mes] = make(map[string]float64)
				}
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
				apropiacionPadre, err = models.GetArbolRubroApropiacionById(session, apropiacionPadre.Padre, ue, vigencia)
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
