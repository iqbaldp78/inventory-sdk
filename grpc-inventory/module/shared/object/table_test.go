package object

import (
	"errors"
	"testing"
)

type foo struct {
	f1 string `db:"f1"`
	f2 string `db:"f2"`
}

func TestTableOperationValid(t *testing.T) {
	temp := foo{}
	input := []struct{
		to *TableOperation
		target foo
	}{
		{&TableOperation{"", "", 0, 0}, temp},
		{&TableOperation{"", "asc", 0, 0}, temp},
		{&TableOperation{"f1", "asc", 0, 0}, temp},
		{&TableOperation{"f1", "basc", 0, 0}, temp},
		{&TableOperation{"f3", "asc", 0, 0}, temp},
	}
	err := errors.New("Raise custom validation")
	output := []interface{}{nil, nil, nil, err, err}

	for index := range input {
		result := input[index].to.Valid(input[index].target)
		if result != output[index] {
			if result.Error() != output[index].(error).Error() {
				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
			}
		}
	}
}
