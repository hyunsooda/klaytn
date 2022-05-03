package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

type KeyTyp uint

const (
	PRIVKEY KeyTyp = iota
	PUBKEY
)

const (
	// 100 years
	DefaultCertExpiration  = time.Hour * 24 * 365 * 100
	DefaultRootKeyPEM      = "rootkey.pem"
	DefaultClientKeyPEM    = "clientkey.pem"
	DefaultClientPubKeyPEM = "clientpubkey.pem"
	DefaultRootCertPEM     = "rootcert.pem"
	DefaultClientCertPEM   = "clientcert.pem"
)

var (
	ErrCertParse = errors.New("Failed to parse certificate")
	ErrDecodePEM = errors.New("Failed to parse pem")
)

func genKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 4096)
}

func genRootCertTemplate(orgName string) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		Subject: pkix.Name{
			Organization: []string{orgName},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(DefaultCertExpiration),
		IsCA:      true,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
}

func genClientCertTemplate(orgName string) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: big.NewInt(2022),
		Subject: pkix.Name{
			Organization: []string{orgName},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(DefaultCertExpiration),

		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
}

func certToFile(certBytes []byte, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	pem.Encode(f, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	return nil
}

func keyToFile(keyTyp KeyTyp, key *rsa.PrivateKey, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	switch keyTyp {
	case PRIVKEY:
		pem.Encode(f, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
	case PUBKEY:
		pem.Encode(f, &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
		})
	}
	return nil
}

func loadPEM(filename string) (*pem.Block, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, ErrDecodePEM
	}
	return block, nil
}

func LoadCert(filename string) (*x509.Certificate, error) {
	block, err := loadPEM(filename)
	if err != nil {
		return nil, err
	}
	return x509.ParseCertificate(block.Bytes)
}

func LoadKey(filename string) (*rsa.PrivateKey, error) {
	block, err := loadPEM(filename)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func GenRootCert(orgName string) error {
	rootKey, err := genKey()
	if err != nil {
		return err
	}
	caCert := genRootCertTemplate(orgName)
	// Self sign
	caCertBytes, err := x509.CreateCertificate(rand.Reader, caCert, caCert, &rootKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	if err = certToFile(caCertBytes, DefaultRootCertPEM); err != nil {
		return err
	}
	if err = keyToFile(PRIVKEY, rootKey, DefaultRootKeyPEM); err != nil {
		return err
	}
	return nil
}

func GenClientCert(orgName string) error {
	rootKey, err := LoadKey(DefaultRootKeyPEM)
	if err != nil {
		return err
	}
	rootcaCert, err := LoadCert(DefaultRootCertPEM)
	if err != nil {
		return err
	}
	clientKey, err := genKey()
	if err != nil {
		return err
	}
	clientCert := genClientCertTemplate(orgName)
	clientCertBytes, err := x509.CreateCertificate(rand.Reader, clientCert, rootcaCert, &clientKey.PublicKey, rootKey)
	if err != nil {
		return err
	}
	if err = certToFile(clientCertBytes, DefaultClientCertPEM); err != nil {
		return err
	}
	if err = keyToFile(PRIVKEY, clientKey, DefaultClientKeyPEM); err != nil {
		return err
	}
	if err = keyToFile(PUBKEY, clientKey, DefaultClientPubKeyPEM); err != nil {
		return err
	}
	return nil
}
