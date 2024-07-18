package utils

// jwt.NewWithClaims() - To create token
// jwt.Parse() - To verify token

import (
	"firstproject/cmd/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "id": id, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(config.Envs.JWTSecret)
    if err != nil {
    return "", err
    }

 	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return config.Envs.JWTSecret, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }