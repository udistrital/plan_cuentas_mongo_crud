package rubroApropiacionHelper

import (
	"testing"

	"github.com/udistrital/plan_cuentas_mongo_crud/helpers/rubroApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestBuildTree(t *testing.T) {
	var expected interface{}

	general := models.General{Vigencia: 2019}
	nodoRubro := models.NodoRubro{General: &general, Hijos: []string{"3-8", "3-9"}}
	raiz := models.NodoRubroApropiacion{NodoRubro: &nodoRubro, ID: "3"}

	expected = rubroApropiacionHelper.BuildTree(&raiz)

	if expected.([]map[string]interface{})[0] == nil {
		t.Error("Se esperaba hacer aserción de tipo correcta", expected)
		t.Fail()
	} else {
		t.Log("TestBuildTree Finalizado Correctamente (OK)")
	}
}

// func TestGetHijoApropiacion(t *testing.T) {
// 	var expected interface{}

// 	expected = rubroApropiacionHelper.GetHijoApropiacion("3", "1", 2019)

// 	if expected.(map[string]interface{})["Codigo"] == nil {
// 		t.Error("Se esperaba un mpa con una llave Codigo", expected)
// 		t.Fail()
// 	} else {
// 		t.Log("TestGetHijoApropiacion Finalizado Correctamente (OK)")
// 	}
// }
