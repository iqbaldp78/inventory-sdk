package jose

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

var conf *config

//Payload used for encode/decode token
type Payload struct {
	Pay     string `json:"pay"`
	Payload Data   `json:"-"`
}

type config struct {
	secret []byte
	issuer string
}

//Data used for extract payload from Payload.Pay Variable
type Data struct {
	ID      int            `json:"id"`
	GroupID int            `json:"group_id"`
	Role    map[string]int `json:"role"`
}

//Extract used for extract token to payload
func Extract(token string) (result Payload, err error) {
	tokenString := strings.Split(token, " ")
	if len(tokenString) < 2 {
		err = errors.New("Bearer token should be have prefix 'Bearer'")
		return
	}

	claims, err := GetClaims(tokenString[1])
	if err != nil {
		return
	}

	byteJSON, _ := json.Marshal(claims)

	err = json.Unmarshal(byteJSON, &result)
	if err == nil {
		result.Payload, err = ValidateJWS(result.Pay)
	}
	return
}

//Setup used for setup jwt configuration
func Setup(secret, issuer string) {
	conf = &config{[]byte(secret), issuer}
}

//IsInitial used for check this object empty or not
func (obj Data) IsInitial() bool {
	if reflect.DeepEqual(obj, Data{}) {
		return true
	}
	return false
}
