package rubroManager

import (
	"testing"
	"time"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestTrRegistrarNodoHoja(t *testing.T) {
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
		Hijos:           []string{"4-1", "4-2"},
		UnidadEjecutora: "1",
		Padre:           "",
		ID:              "4",
	}

	if err := rubroManager.TrRegistrarNodoHoja(&nodoRubro, models.NodoRubroCollection); err != nil {
		t.Error("error: ", err)
		t.Fail()
	} else {
		t.Log("TestTrRegistrarNodoHoja Finalizado Correctamente (OK)")
	}
}

func TestGetRaices(t *testing.T) {
	roots := rubroManager.GetRaices("1")

	if len(roots) > 0 {
		t.Log("TestGetRaices Finalizado Correctamente (OK)")
	} else {
		t.Error("Se esperaba un arreglo de al menos una posición: ", roots)
		t.Fail()
	}
}

func TestGetNodo(t *testing.T) {
	node := rubroManager.GetNodo("4", "1")

	if node == nil {
		t.Error("Se esperaba un nodo del árbol de rubros ", node)
		t.Fail()
	} else {
		t.Log("TestGetNodo Finalizado Correctamente (OK)")
	}Y
}

func TestGetHijoRubro(t *testing.T) {
	nodeSon := rubroManager.GetHijoRubro("4", "1")

	if nodeSon == nil {
		t.Error("Se esperaba un nodo del árbol de rubros ", nodeSon)
		t.Fail()
	} else {
		t.Log("TestGetHijoRubro Finalizado Correctamente (OK)")
	}
}
