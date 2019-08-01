package movimientoManager

import (
	"testing"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestGetOneMovimientoBy(t *testing.T) {
	expected, err := movimientoManager.GetOneMovimientoByTipo(3, "CDP")

	if err != nil {
		t.Error("error: ", err)
		t.Fail()
	} else {
		if &expected != nil {
			t.Log("TestGetOneMovimientoByTipo Finalizado Correctamente (OK)")
		} else {
			t.Error("expeted is nil", expected)
			t.Fail()
		}
	}
}

func TestGetOneMovimientoParameterByHijo(t *testing.T) {
	expected, err := movimientoManager.GetOneMovimientoParameterByHijo("RP")

	if err != nil {
		t.Error("error: ", err)
		t.Fail()
	} else {
		if &expected != nil {
			t.Log("TestGetOneMovimientoParameterByHijo Finalizado Correctamente (OK)")
		} else {
			t.Error("expeted is nil", expected)
			t.Fail()
		}
	}
}

func TestSaveMovimientoParameter(t *testing.T) {
	data := models.MovimientoParameter{
		TipoMovimientoHijo:  "OrdenPago",
		TipoMovimientoPadre: "RP",
		Multiplicador:       -1,
	}

	if err := movimientoManager.SaveMovimientoParameter(&data); err != nil {
		t.Error("error: ", err)
		t.Fail()
	} else {
		t.Log("TestSaveMovimientoParameter Finalizado Correctamente (OK)")
	}
}
