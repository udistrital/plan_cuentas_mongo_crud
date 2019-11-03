package tests

import (
	"testing"
	"time"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestRubroRegistrationSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("TestTrRegistrarNodoHoja Finalizado Correctamente (OK)")
		}
	}()

	general := models.General{
		Vigencia:          2019,
		Nombre:            "Rubro de prueba",
		Descripcion:       "Rubro de prueba",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
		Activo:            true,
	}

	nodoRubro3 := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "",
		ID:              "3",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro3, models.NodoRubroCollection); err != nil {
		panic(err.Error())
	}

	nodoRubro2 := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "",
		ID:              "2",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro2, models.NodoRubroCollection); err != nil {
		panic(err.Error())
	}

	nodoRubro3_1 := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "3",
		ID:              "3-1",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro3_1, models.NodoRubroCollection); err != nil {
		panic(err.Error())
	}

	nodoRubro2_1 := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "2",
		ID:              "2-1",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro2_1, models.NodoRubroCollection); err != nil {
		panic(err.Error())
	}

}

func TestRubroRegistrationFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("TestTrRegistrarNodoHoja Finalizado Correctamente (OK)")
		} else {
			t.Error("error: Rubro restriction does not work ")
			t.Fail()
		}
	}()

	general := models.General{
		Vigencia:          2019,
		Nombre:            "Rubro de prueba",
		Descripcion:       "Rubro de prueba",
		FechaCreacion:     time.Now(),
		FechaModificacion: time.Now(),
		Activo:            true,
	}

	nodoRubro4 := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "",
		ID:              "4",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro4, models.NodoRubroCollection); err != nil {
		panic(err.Error())
	}

}
