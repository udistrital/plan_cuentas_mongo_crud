package presupuestalasignationtestcases

import "testing"

func TestPresupuestalAssignationPipeline(t *testing.T) {
	// <setup code>
	t.Run("Rubro's registration success process test", TestRubroRegistrationSuccess)
	t.Run("Rubro's registration node code fail process test", TestRubroRegistrationNodeCodeFail)
	t.Run("Apropiation's success registration success process test", TestApropiationRegistrationSuccess)
	t.Run("Apropiation's registration without rubro test", TestApropiationRegistrationWithOutRubroFail)
	t.Run("Apropiation balanced after registration success", TestCheckForApropiationTreeBalance)
	t.Run("Apropiation does not balanced after registration success", TestCheckForApropiationTreeBalanceFail)

	// <tear-down code>
}
