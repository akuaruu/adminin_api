package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// in production lvl, should be taken from Env Variable (.env)
var secretKey =[]byte("Secret-Key")

type Claims struct{
	AdminID int64 `json:"admin_id"`
	jwt.RegisteredClaims
}

//Generate token for "access card"
func GenerateToken(adminID int64) (string, error){
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
			AdminID: adminID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt : jwt.NewNumericDate(expirationTime),
			},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil{
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error){
	claims := &Claims{}

	token, err:= jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error){
		
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid token algorithm")
		}
		
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid{
		return nil, errors.New("invalid token")
	}

	return  claims, nil
} 


