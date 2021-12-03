package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path"
)

func PkRsa(keyPath string, keyName string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	_, err = os.Stat(keyPath)
	if os.IsNotExist(err) {
		log.Fatalln("Folder does not exist")
	}

	privatePem, err := os.Create(path.Join(keyPath, keyName+"private.pem"))
	if err != nil {
		panic(err)
	}

	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		panic(err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create(path.Join(keyPath, keyName+"public.pem"))
	if err != nil {
		panic(err)
	}

	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		panic(err)
	}

}

func privateAndPublicKeyInMemory() ([]byte, []byte) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &key.PublicKey
	publicKeyMemory := public_key(publicKey)

	privateKey := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKey,
	}

	var privateKeyMemory []byte = pem.EncodeToMemory(block)
	return privateKeyMemory, publicKeyMemory
}

func public_key(key *rsa.PublicKey) []byte {
	keyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		panic(err)
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keyBytes,
	}
	return pem.EncodeToMemory(block)
}
