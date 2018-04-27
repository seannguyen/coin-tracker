package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

func Create() string {
	token := getTokenObject()
	apiSecret := viper.GetString("QUOINEX_API_SECRET")
	tokenString, err := token.SignedString([]byte(apiSecret))
	if err != nil {
		log.Panic(err)
	}
	return tokenString
}

func getTokenObject() *jwt.Token {
	apiId := viper.GetString("QUOINEX_API_ID")
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path":     "/accounts",
		"nonce":    time.Now().UnixNano(),
		"token_id": apiId,
	})
}
