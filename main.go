package main

import (
	"jwt/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		return
	}

	/*simple := generateSimple()
	//fmt.Println(simple)
	count := strings.Count(simple, ".")
	if count < 3 {
		simple += "."
	}

	parse, _ := jwt.Parse(simple, func(token *jwt.Token) (interface{}, error) {
		raw := token.Header
		fmt.Println(raw["alg"])
		fmt.Println("------------")
		//fmt.Println(raw)
		return token, nil
	})
	fmt.Println(parse)

	signed := generateSigned()
	jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		i := token.Header["alg"]
		fmt.Println(i)
		return token, nil
	})
	//fmt.Println(signed)

	symmetric := generateSymmetric()
	jwt.Parse(symmetric, func(token *jwt.Token) (interface{}, error) {
		i := token.Header["alg"]
		fmt.Println(i)
		return token, nil
	})*/
	//fmt.Println(symmetric)
	/*fmt.Println("---------------------------------------------")
	pk, pp := privateAndPublicKeyInMemory()
	withClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	fromPEM, err := jwt.ParseRSAPrivateKeyFromPEM(pk)
	signedString, err := withClaims.SignedString(fromPEM)
	if err != nil {
		panic(err)
	}

	keyFromPEM, err := jwt.ParseRSAPublicKeyFromPEM(pp)
	parse, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return keyFromPEM, nil
	})
	fmt.Println(parse)
	fmt.Println("---------------------------------------------")

	key, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("===============================================")
	fmt.Println(key.PublicKey)
	privateKey, err := x509.MarshalECPrivateKey(key)
	fmt.Println(privateKey)
	//os.WriteFile("C:/Users/davidul/private.key", privateKey, 0644)
	pkRsa("testKey")*/
}
