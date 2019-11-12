package presupuestalasignationtestcases

import (
	"testing"
	"time"

	arbolrubroapropiacioncompositor "github.com/udistrital/plan_cuentas_mongo_crud/compositors/arbolRubroApropiacionCompositor"

	testhelper "github.com/udistrital/plan_cuentas_mongo_crud/helpers/testHelper"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/rubroApropiacionManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

func TestApropiationRegistrationSuccess(t *testing.T) {

	defer testhelper.TestsDeferFucntionForSuccessComprobation(t, "OK", "Error")

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

	defer testhelper.TestsDeferFucntionForFailedComprobation(t, "OK", "error: Apropiacion restriction does not work ")

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

func TestCheckForApropiationTreeBalance(t *testing.T) {
	defer testhelper.TestsDeferFucntionForSuccessComprobation(t, "OK", "error: Balance check fails.")
	var movimientos []models.Movimiento
	balanced, _, _, err := arbolrubroapropiacioncompositor.CheckForAprTreeBalanceWithSimulation(movimientos, "2019", "1")

	if err != nil {
		panic(err.Error())
	}

	if !balanced {
		panic("The balanced flag from the function espected was true but got false.")
	}

}

func TestCheckForApropiationTreeBalanceFail(t *testing.T) {
	defer testhelper.TestsDeferFucntionForFailedComprobation(t, "OK", "The balanced flag from the function espected was false but got true.")
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

	raiz2 := models.NodoRubroApropiacion{
		NodoRubro:    &nodoRubro,
		ID:           "2-02",
		ValorInicial: 1000,
		ValorActual:  0,
		Productos:    map[string]map[string]interface{}{},
		Estado:       models.EstadoRegistrada,
		Padre:        "2",
	}

	if err := rubroApropiacionManager.TrRegistrarNodoHoja(&raiz2, "1", 2019); err != nil {
		panic(err.Error())
	}

	var movimientos []models.Movimiento
	balanced, _, _, err := arbolrubroapropiacioncompositor.CheckForAprTreeBalanceWithSimulation(movimientos, "2019", "1")

	if err != nil {
		panic(err.Error())
	}

	if !balanced {
		panic("not balanced")
	}
}
