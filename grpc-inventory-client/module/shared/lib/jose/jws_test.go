package jose

import (
	"testing"
)

func TestGenerateJWS(t *testing.T) {
	input := Data{1, 1, map[string]int{}}
	output := ""

	result := GenerateJWS(input)
	if result == output {
		t.Errorf("Expected result not to be `%v`. Got `%v`", output, result)
	}
}

func TestValidateJWS(t *testing.T) {
	input := []string{
		"",
		`{"payload":"WyJhIiwiYiIsImMiXQ","protected":"eyJhbGciOiJIUzI1NiJ9","signature":"4j1eG5yf74DD8WdrdRAcSF1whi6SMRS98wytsFjusxk"}`,
		`{"payload":"eyJpZCI6MSwiZ3JvdXBfaWQiOjEsInJvbGUiOnt9fQ","protected":"eyJhbGciOiJIUzI1NiJ9","signature":"fTDJmjBBgFeOIqswSw_CXtKcTFvzQUHA4Uf9z3qM5x"}`,
		`{"payload":"eyJpZCI6MSwiZ3JvdXBfaWQiOjEsInJvbGUiOnt9fQ","protected":"eyJhbGciOiJIUzI1NiJ9","signature":"fTDJmjBBgFeOIqswSw_CXtKcTFvzQUHA4Uf9z3qMu5Y"}`,
	}
	output := []bool{true, true, true, false}

	for index := range input {
		result, _ := ValidateJWS(input[index])
		if result.IsInitial() != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result.IsInitial())
		}
	}
}
