package inmemorygenerator

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"

	"code.cloudfoundry.org/cf-operator/pkg/credsgen"
	"github.com/pkg/errors"
)

// GenerateRSAKey generates an RSA key using go's standard crypto library
func (g InMemoryGenerator) GenerateRSAKey(name string) (credsgen.RSAKey, error) {
	log.Println("Generating RSA key ", name)

	// generate private key
	private, err := rsa.GenerateKey(rand.Reader, g.Bits)
	if err != nil {
		return credsgen.RSAKey{}, errors.Wrap(err, "Generating private key")
	}
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(private),
	}
	privatePEM := pem.EncodeToMemory(privateBlock)

	// Calculate public key
	public := private.Public().(*rsa.PublicKey)
	publicSerialized, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return credsgen.RSAKey{}, errors.Wrap(err, "Generating public key")
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicSerialized,
	}
	publicPEM := pem.EncodeToMemory(publicBlock)

	key := credsgen.RSAKey{
		PrivateKey: privatePEM,
		PublicKey:  publicPEM,
	}
	return key, nil
}
