package movimientoCompositor

import (
	"testing"

	"github.com/udistrital/plan_cuentas_mongo_crud/compositors/movimientoCompositor"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestAddMovimientoTransaction(t *testing.T) {
	var expected []interface{}

	movimiento := models.Movimiento{
		IDPsql:         3,
		Valor:          30000,
		Tipo:           "CDP",
		DocumentoPadre: 3,
		FechaRegistro:  "2019-01-08",
		Movimientos:    map[string]float64{},
	}

	expected = movimientoCompositor.AddMovimientoTransaction(movimiento)

	if len(expected) > 0 {
		t.Log("TestAddMovimientoTransaction Finalizado Correctamente (OK)")
	} else {
		t.Error("Se esperaba un arreglo con más de una posición ", expected)
		t.Fail()
	}
}

func TestBuildPropagacionValoresTr(t *testing.T) {
	var expected []interface{}

	movimiento := models.Movimiento{
		IDPsql:         3,
		Valor:          30000,
		Tipo:           "RP",
		DocumentoPadre: 3,
		FechaRegistro:  "2019-01-08",
		Movimientos:    map[string]float64{},
	}

	expected = movimientoCompositor.BuildPropagacionValoresTr(movimiento)

	if len(expected) > 0 {
		t.Log("TestBuildPropagacionValoresTr Finalizó Correctamente (OK)")
	} else {
		t.Error("Se esperaba un arreglo con al menos una posición ", expected)
		t.Fail()
	}
}
