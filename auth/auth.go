package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)



type JWTClaim struct {
	UserId string `json:"userId" bson:"userId"`
	Username string `json:"username" bson:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userId string) (tokenString string, err error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	jwtKey := os.Getenv("JWT_SECRET")

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		UserId: userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))
	return
}

func ValidateToken(signedToken string) (jwtClaim *JWTClaim ,err error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	jwtKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		jwtClaim = nil
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token expired")
		jwtClaim = nil
		return
	}

	jwtClaim = claims
	err = nil
	return 
}
