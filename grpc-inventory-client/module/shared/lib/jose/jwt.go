//Package jose used for implementation of JOSE standards (JWE, JWS, JWT)
package jose

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateJWT used for generate JWT token
func GenerateJWT(data Data, expireHour time.Duration) (string, error) {
	subject := fmt.Sprintf("access granted for %d hour(s)", expireHour)
	payload := GenerateJWS(data)

	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * expireHour).Unix(),
		"iss": conf.issuer,
		"pay": payload,
		"sub": subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(conf.secret)
}

//GetClaims used for extract JWT token payload
func GetClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return conf.secret, nil
	})

	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, err
		}
	}
	return nil, err
}
