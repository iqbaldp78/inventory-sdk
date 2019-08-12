package postgres

import (
	"log"
	"strings"
	"testing"
)

func TestNewConnection(t *testing.T) {
	input := []struct {
		dbname   string
		username string
		password string
		host     string
		port     int
	}{
		{"", "", "", "", 0},
		{"postgres", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort},
	}
	output := []interface{}{"connect", nil}

	for index := range input {
		result := New(input[index].dbname, input[index].username, input[index].password, input[index].host, input[index].port)
		log.Println(input[index], result)
		if result != output[index] {
			if !strings.Contains(result.Error(), output[index].(string)) {
				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
			}
		}
	}
}

func TestGenerateTestDB(t *testing.T) {
	output := []interface{}{nil, nil}
	for index := range output {
		_, result := GenerateTestDB()
		if result != output[index] {
			if result.Error() != output[index].(error).Error() {
				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
			}
		}
	}
}
