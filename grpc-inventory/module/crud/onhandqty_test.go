package crud

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/gocraft/dbr"
// 	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/lib/postgres"
// 	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/object"
// )

// var db *dbr.Session

// func TestSetup(t *testing.T) {
// 	var err error
// 	db, err = postgres.GenerateTestDB()
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// }

// func TestSampleUnmarshalTime(t *testing.T) {
// 	input := Sample{1, "a", 1.2, false, "", "", dbr.NewNullTime(time.Now()), dbr.NewNullTime(time.Now())}
// 	output := ""

// 	input.UnmarshalTime()
// 	if input.CreatedOn == output || input.UpdatedOn == output {
// 		t.Errorf("Expected result not to be `%v`. Got `%v`", output, input.CreatedOn)
// 	}
// }

// func TestSampleMarshalTime(t *testing.T) {
// 	err := errors.New(`parsing time "2008-07-03" as "2006-01-02 15:04:05": cannot parse "" as "15"`)
// 	input := []Sample{
// 		{1, "a", 1.2, false, "2008-07-03", "2008-07-03", dbr.NullTime{}, dbr.NullTime{}},
// 		{1, "a", 1.2, false, "2008-07-03 00:00:00", "2008-07-03", dbr.NullTime{}, dbr.NullTime{}},
// 		{1, "a", 1.2, false, "2008-07-03 00:00:00", "2008-07-03 00:00:00", dbr.NullTime{}, dbr.NullTime{}},
// 	}
// 	output := []interface{}{err, err, nil}

// 	for index := range input {
// 		result := input[index].MarshalTime()
// 		if result != output[index] {
// 			if result.Error() != output[index].(error).Error() {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 			}
// 		}
// 	}
// }

// func TestSampleInsert(t *testing.T) {
// 	input := []Sample{
// 		{0, "cde", 0.10, true, "", "", dbr.NewNullTime(time.Time{}), dbr.NewNullTime(time.Time{})},
// 	}
// 	output := []interface{}{nil}
// 	for index := range input {
// 		result := input[index].Insert(db)
// 		if result != output[index] {
// 			if result.Error() != output[index].(error).Error() {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 			}
// 		}
// 	}
// }

// func TestSampleSelect(t *testing.T) {
// 	err := errors.New("dbr: not found")
// 	input := []int64{2, 3}
// 	output := []interface{}{nil, err}

// 	for index := range input {
// 		temp := Sample{}
// 		result := temp.Select(db, input[index])
// 		if result != output[index] {
// 			if result.Error() != output[index].(error).Error() {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 			}
// 		}
// 	}
// }

// func TestSampleUpdate(t *testing.T) {
// 	input := []Sample{
// 		{2, "cde", 0.10, true, "", "", dbr.NewNullTime(time.Time{}), dbr.NewNullTime(time.Time{})},
// 	}
// 	output := []interface{}{nil}

// 	for index := range input {
// 		result := input[index].Update(db)
// 		if result != output[index] {
// 			if result.Error() != output[index].(error).Error() {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 			}
// 		}
// 	}
// }

// func TestSampleSearch(t *testing.T) {
// 	input := []object.SearchParam{
// 		{Key: "abc", Limit: 10, Offset: 0},
// 		{Key: "abc", Limit: 10, Offset: 1},
// 		{Key: "abc", Limit: 0, Offset: 0},
// 		{Key: "abc", Limit: 1001, Offset: 0},
// 	}
// 	output := []int{1, 0, 1, 1}

// 	for index := range input {
// 		temp := Sample{}
// 		result := temp.Search(db, input[index])
// 		if len(result.Data.([]Sample)) != output[index] {
// 			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], len(result.Data.([]Sample)))
// 		}
// 	}
// }

// func TestSampleDelete(t *testing.T) {
// 	input := []int64{2, 3}
// 	output := []interface{}{nil, nil}

// 	for index := range input {
// 		temp := Sample{}
// 		result := temp.Delete(db, input[index])
// 		if result != output[index] {
// 			if result.Error() != output[index].(error).Error() {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 			}
// 		}
// 	}
// }

// func TestSampleIsInitial(t *testing.T) {
// 	input := []Sample{
// 		{1, "a", 1.2, false, "", "", dbr.NewNullTime(time.Now()), dbr.NewNullTime(time.Now())},
// 		{},
// 	}
// 	output := []bool{false, true}

// 	for index := range input {
// 		result := input[index].IsInitial()
// 		if result != output[index] {
// 			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
// 		}
// 	}
// }

// func TestSampleClear(t *testing.T) {
// 	input := Sample{1, "a", 1.2, false, "", "", dbr.NewNullTime(time.Now()), dbr.NewNullTime(time.Now())}
// 	output := Sample{}

// 	input.Clear()
// 	if input.ID != output.ID || input.FieldA != output.FieldA || input.FieldB != output.FieldB || input.FieldC != output.FieldC {
// 		t.Errorf("Expected result to be `%v`. Got `%v`", output, input)
// 	}
// }

// func TestCleanUp(t *testing.T) {
// 	_, err := postgres.GenerateTestDB()
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// }
