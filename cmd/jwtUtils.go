package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.StandardClaims
}

func GenerateSimple() string {
	claims := customClaims()

	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	signingString, err := token.SigningString()
	if err != nil {
		panic(err)
	}

	return signingString
}

func GenerateSymmetric(secretKey string) string {
	claims := customClaims()
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := withClaims.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return signedString
}

func GenerateSigned() string {
	customClaims := customClaims()
	privateKey, _ := privateAndPublicKeyInMemory()
	jwtWithClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, customClaims)
	fromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		panic(err)
	}
	signedString, err := jwtWithClaims.SignedString(fromPEM)

	return signedString
}

func validate(t jwt.Token) {
	method := t.Method
	fmt.Println(method)
}

func customClaims() CustomClaims {
	now := time.Now()
	//duration := time.Duration(30 * time.Second)  //nanosecond
	datePlus3 := now.AddDate(0, 0, 3)

	return CustomClaims{
		"David",
		"Ulicny",
		jwt.StandardClaims{
			Audience:  "Z",
			ExpiresAt: datePlus3.Unix(),
			Id:        "123",
			IssuedAt:  now.Unix(),
			Issuer:    "asd",
			NotBefore: 123,
			Subject:   "asd",
		},
	}
}
