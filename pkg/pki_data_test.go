package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestEcd(t *testing.T) {
	Ecd()
}

func TestGenKeysRsa(t *testing.T) {
	privateKey, publicKey := GenKeysRsa()
	assert.Nil(t, privateKey.Validate())
	assert.NotNil(t, publicKey.N)
}

func TestGenKeysEcdsa(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)
}

func TestMarshalRsa(t *testing.T) {
	privateKey, publicKey := GenKeysRsa()
	mPrivateKey, mPublicKey := MarshalRsa(privateKey, publicKey)
	assert.NotNil(t, mPrivateKey)
	assert.NotNil(t, mPublicKey)
}

func TestMarshalEcdsa(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	mPrivateKey, mPublicKey := MarshalEcdsa(privateKey, publicKey)
	assert.NotNil(t, mPrivateKey)
	assert.NotNil(t, mPublicKey)
}

func TestEncodePem(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	mPrivateKey, mPublicKey := MarshalEcdsa(privateKey, publicKey)
	ePrivateKey, ePublicKey := EncodePem(mPrivateKey, mPublicKey, ECDSA)
	assert.NotNil(t, ePrivateKey)
	assert.NotNil(t, ePublicKey)
}

func TestEncodePemToFile(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	mPrivateKey, mPublicKey := MarshalEcdsa(privateKey, publicKey)
	ePrivateKey, ePublicKey := EncodePem(mPrivateKey, mPublicKey, ECDSA)
	EncodePemToFile(ePrivateKey, ePublicKey, "", "")
	assert.NotNil(t, ePrivateKey)
	assert.NotNil(t, ePublicKey)
	// read file
	file, err := os.Open(path.Join("", "private.pem"))
	assert.Nil(t, err)
	assert.NotNil(t, file)

	file, err = os.Open(path.Join("", "public.pem"))
	assert.Nil(t, err)
	assert.NotNil(t, file)

	EncodePemToFile(ePrivateKey, ePublicKey, "", "test")
	file, err = os.Open(path.Join("", "test_private.pem"))
	assert.Nil(t, err)
	assert.NotNil(t, file)

	//delete files
	err = os.Remove(path.Join("", "private.pem"))
	assert.Nil(t, err)
	err = os.Remove(path.Join("", "public.pem"))
	assert.Nil(t, err)
	err = os.Remove(path.Join("", "test_private.pem"))
	assert.Nil(t, err)
	err = os.Remove(path.Join("", "test_public.pem"))
	assert.Nil(t, err)
}

func TestDecodePem(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	mPrivateKey, mPublicKey := MarshalEcdsa(privateKey, publicKey)
	ePrivateKey, ePublicKey := EncodePem(mPrivateKey, mPublicKey, ECDSA)
	EncodePemToFile(ePrivateKey, ePublicKey, "", "")
	dPublicKey := DecodePublicPemFromFile("public.pem")
	assert.NotNil(t, dPublicKey)
	err := os.Remove(path.Join("", "public.pem"))
	assert.Nil(t, err)
	err = os.Remove(path.Join("", "private.pem"))
	assert.Nil(t, err)
}

func TestDecodePemFromFile(t *testing.T) {
	privateKey, publicKey := GenKeysEcdsa()
	mPrivateKey, mPublicKey := MarshalEcdsa(privateKey, publicKey)
	ePrivateKey, ePublicKey := EncodePem(mPrivateKey, mPublicKey, ECDSA)
	EncodePemToFile(ePrivateKey, ePublicKey, "", "")
	dPublicKey := DecodePublicPemFromFile("public.pem")
	assert.NotNil(t, dPublicKey)

	rsa := UnmarshalPublicRsa(dPublicKey)
	assert.NotNil(t, rsa)
}

func TestUnmarshalPublicRsa(t *testing.T) {
	privateKey, publicKey := GenKeysRsa()
	mPrivateKey, mPublicKey := MarshalRsa(privateKey, publicKey)
	ePrivateKey, ePublicKey := EncodePem(mPrivateKey, mPublicKey, RSA)
	EncodePemToFile(ePrivateKey, ePublicKey, "", "")
	dPublicKey := DecodePublicPemFromFile("public.pem")
	assert.NotNil(t, dPublicKey)

	rsa := UnmarshalPublicRsa(dPublicKey)
	assert.NotNil(t, rsa)
}
