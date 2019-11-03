package presupuestalasignationtestcases

import (
	"testing"
	"time"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroApropiacionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestApropiationRegistrationSuccess(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		} else {
			t.Log("OK")
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

	nodoRubro := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "3",
	}

	raiz := models.NodoRubroApropiacion{
		NodoRubro:    &nodoRubro,
		ID:           "3-01",
		ValorInicial: 30000,
		ValorActual:  0,
		Productos:    map[string]map[string]interface{}{},
		Estado:       models.EstadoRegistrada,
		Padre:        "3",
	}

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(&raiz, "1", 2019); err != nil {
		panic(err.Error())
	}

	raiz2 := models.NodoRubroApropiacion{
		NodoRubro:    &nodoRubro,
		ID:           "2-01",
		ValorInicial: 30000,
		ValorActual:  0,
		Productos:    map[string]map[string]interface{}{},
		Estado:       models.EstadoRegistrada,
		Padre:        "2",
	}

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(&raiz2, "1", 2019); err != nil {
		panic(err.Error())
	}

}

func TestApropiationRegistrationWithOutRubroFail(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("OK")
		} else {
			t.Error("error: Apropiacion restriction does not work ")
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

	nodoRubro := models.NodoRubro{
		General:         &general,
		UnidadEjecutora: "1",
		Padre:           "3",
	}

	raiz := models.NodoRubroApropiacion{
		NodoRubro:    &nodoRubro,
		ID:           "3-2",
		ValorInicial: 30000,
		ValorActual:  0,
		Productos:    map[string]map[string]interface{}{},
		Estado:       models.EstadoRegistrada,
	}

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(&raiz, "1", 2019); err != nil {
		panic(err.Error())
	}

}
