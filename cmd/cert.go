package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

func createCert() (*x509.Certificate, []byte) {
	certificate := x509.Certificate{
		Signature:          nil,
		SignatureAlgorithm: 0,
		PublicKeyAlgorithm: 0,
		PublicKey:          nil,
		Version:            0,
		SerialNumber:       big.NewInt(1),
		Issuer: pkix.Name{
			Country:      []string{"CZ"},
			Organization: []string{"David"},
		},
		Subject: pkix.Name{
			Country:            []string{"CZ"},
			Organization:       []string{"SampleOrg"},
			OrganizationalUnit: []string{"SampleUnit"},
			PostalCode:         []string{"25091"},
		},
		NotBefore: time.Now().Add(-10 * time.Second),
		NotAfter:  time.Now().AddDate(10, 0, 0),
		KeyUsage:  x509.KeyUsageCertSign,
		//ExtKeyUsage:                 nil,
		IsCA:       false,
		MaxPathLen: 2,
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, &certificate, &certificate, &key.PublicKey, &key)
	if err != nil {
		panic(err)
	}

	parseCertificate, err := x509.ParseCertificate(certificateBytes)
	if err != nil {
		panic(err)
	}

	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificateBytes,
	}
	memory := pem.EncodeToMemory(&block)

	return parseCertificate, memory

}
