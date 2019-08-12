package jose

import (
	"encoding/json"

	"gopkg.in/square/go-jose.v2"
)

/*GenerateJWS used for generate JWS*/
func GenerateJWS(data Data) string {
	signer, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: conf.secret}, nil)
	payload, _ := json.Marshal(data)
	object, _ := signer.Sign(payload)

	return object.FullSerialize()
}

/*ValidateJWS used for check JWS*/
func ValidateJWS(signed string) (Data, error) {
	data := Data{}
	object, err := jose.ParseSigned(signed)
	if err != nil {
		return data, err
	}

	output, err := object.Verify(conf.secret)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(output, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
