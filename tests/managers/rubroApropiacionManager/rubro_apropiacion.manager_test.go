package rubroApropiacionManager

import (
	"testing"
	"time"

	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroApropiacionManager"
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
	}

	raiz := models.NodoRubroApropiacion{
		NodoRubro:            &nodoRubro,
		ID:                   "4",
		ApropiacionInicial:   30000,
		ApropiacionUtilizada: 0,
		Movimientos:          []string{},
		Productos:            map[string]map[string]interface{}{},
		Estado:               "Aprobado",
	}

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(&raiz, "1", 2019); err != nil {
		t.Error("error: ", err)
		t.Fail()
	} else {
		t.Log("TestTrRegistrarNodoHoja Finalizado Correctamente (OK)")
	}
}
