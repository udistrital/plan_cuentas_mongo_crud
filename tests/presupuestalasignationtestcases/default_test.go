package presupuestalasignationtestcases

import "testing"

func TestPresupuestalAssignationPipeline(t *testing.T) {
	// <setup code>
	t.Run("Rubro's registration success process test", TestRubroRegistrationSuccess)
	t.Run("Rubro's registration node code fail process test", TestRubroRegistrationNodeCodeFail)
	t.Run("Apropiation's registration success process test", TestApropiationRegistrationSuccess)
	t.Run("Apropiation's registration without rubro test", TestApropiationRegistrationWithOutRubroFail)

	// <tear-down code>
}
