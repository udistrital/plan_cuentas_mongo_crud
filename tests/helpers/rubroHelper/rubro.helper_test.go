package rubroHelper

import (
	"testing"

	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestBuildTree(t *testing.T) {
	var expected interface{}

	general := models.General{Vigencia: 2019}
	raiz := models.NodoRubro{ID: "3", General: &general, Hijos: []string{"3-1", "3-2"}}

	expected = rubroHelper.BuildTree(&raiz)

	if expected.([]map[string]interface{})[0] == nil {
		t.Error("Se esperaba hacer aserción de tipo correcta", expected)
		t.Fail()
	} else {
		t.Log("TestBuildTree Finalizado Correctamente (OK)")
	}
}
