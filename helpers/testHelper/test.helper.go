package testhelper

import "testing"

// TestsDeferFucntionForFailedComprobation you must use this function if the test is for check if a wrong paramater give an error in the operation.
func TestsDeferFucntionForFailedComprobation(t *testing.T, successMessage, errorMessage string) {
	if r := recover(); r != nil {
		t.Log("OK")
	} else {
		t.Error(errorMessage, r)
		t.Fail()
	}
}

// TestsDeferFucntionForSuccessComprobation you must use this for a normañ success comprobation.
func TestsDeferFucntionForSuccessComprobation(t *testing.T, successMessage, errorMessage string) {
	if r := recover(); r != nil {
		t.Error(errorMessage, r)
		t.Fail()
	} else {
		t.Log(successMessage)
	}
}
