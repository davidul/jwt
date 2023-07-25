package pkg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"
)

func GenKeysEcdsa() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.PublicKey
	return privateKey, &publicKey
}

// generate RSA keys
func GenKeysRsa() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.PublicKey
	return privateKey, &publicKey
}

// marshal RSA keys to PEM
func MarshalRsa(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) ([]byte, []byte) {
	mPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	mPublicKey := x509.MarshalPKCS1PublicKey(publicKey)

	return mPrivateKey, mPublicKey
}

// marshal ECDSA keys to PEM
func MarshalEcdsa(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, []byte) {
	mPrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
	}

	mPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
	}

	return mPrivateKey, mPublicKey
}

// encode PEM to memory
func EncodePem(mPrivateKey []byte, mPublicKey []byte) ([]byte, []byte) {
	privateBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: mPrivateKey}
	publicBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: mPublicKey}

	memoryPrivate := pem.EncodeToMemory(privateBlock)
	memoryPublic := pem.EncodeToMemory(publicBlock)

	return memoryPrivate, memoryPublic
}

// encode PEM to file
func EncodePemToFile(mPrivateKey []byte, mPublicKey []byte, filePath string, prefix string) {
	privateBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: mPrivateKey}
	publicBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: mPublicKey}

	privatePrefix := prefix + "_private.pem"
	publicPrefix := prefix + "_public.pem"

	if prefix == "" {
		privatePrefix = "private.pem"
		publicPrefix = "public.pem"
	}

	privateFile, err := os.Create(path.Join(filePath, privatePrefix))
	if err != nil {
		fmt.Println(err)
	}

	publicFile, err := os.Create(path.Join(filePath, publicPrefix))
	if err != nil {
		fmt.Println(err)
	}

	err = pem.Encode(privateFile, privateBlock)
	if err != nil {
		return
	}
	err = pem.Encode(publicFile, publicBlock)
	if err != nil {
		return
	}
}

func DecodePrivatePemFromFile(path string) *pem.Block {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	privatePem, _ := pem.Decode(bytes)

	return privatePem
}

func DecodePublicPemFromFile(path string) *pem.Block {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	publicPem, _ := pem.Decode(bytes)

	return publicPem
}

// unmarshal PEM to RSA keys
func UnmarshalPublicRsa(publicPem *pem.Block) *rsa.PublicKey {
	publicKey, err := x509.ParsePKCS1PublicKey(publicPem.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	return publicKey
}
func UnmarshalPrivateRsa(privatePem *pem.Block) *rsa.PrivateKey {
	privateKey, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	return privateKey
}

func Ecd() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.PublicKey
	mPrivateKey, err := x509.MarshalECPrivateKey(privateKey)

	mPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)

	privateBlock := &pem.Block{Type: "PRIVATE KEY", Bytes: mPrivateKey}
	publicBlock := &pem.Block{Type: "PUBLIC KEY", Bytes: mPublicKey}
	pem.EncodeToMemory(privateBlock)
	pem.EncodeToMemory(publicBlock)

	file, err := os.Create("path")
	pem.Encode(file, privateBlock)
	pem.Encode(file, publicBlock)
}
