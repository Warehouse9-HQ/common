package address__test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"strings"
	"testing"

	"github.com/Warehouse9HQ/common/address"
)

func TestAddressPublicConversions(t *testing.T) {
	curve := elliptic.P256
	privateKey, err := ecdsa.GenerateKey(curve(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Generated private key")

	public := &privateKey.PublicKey

	msg := "say hello to the new age of storage services!"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		t.Fatal(err)
	}

	valid := ecdsa.VerifyASN1(public, hash[:], sig)

	if !valid {
		t.Fatal("Failed to verify original signature")
	}

	t.Log("Verified original signature")

	addr := address.PublicToAddress(public)
	if !strings.HasPrefix(addr, "0s") {
		t.Fatal("PublicToAddress failed - address doesn't begin with '0s'")
	}

	t.Log("Address:", addr)

	publicKey, err := address.AddressToPublic(curve(), addr)

	if err != nil {
		log.Fatal(err)
	}

	t.Log("Derived public key from address")

	valid = ecdsa.VerifyASN1(publicKey, hash[:], sig)

	if !valid {
		t.Fatal("Failed to verify signature with key derived from address")
	}

	t.Log("Verified signature with derived key")

}

func TestPrivateKeyConversions(t *testing.T) {
	curve := elliptic.P256
	privateKey, err := ecdsa.GenerateKey(curve(), rand.Reader)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Generated private key")

	public := &privateKey.PublicKey

	msg := "say hello to the new age of storage services!"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		t.Fatal(err)
	}

	valid := ecdsa.VerifyASN1(public, hash[:], sig)

	if !valid {
		t.Fatal("Failed to verify original signature")
	}

	t.Log("Verified original signature")

	encoded := address.EncodePrivate(privateKey)

	t.Log("Encoded private key", encoded)

	decoded, err := address.DecodePrivate(curve(), encoded)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Decoded private key")

	sig2, err := ecdsa.SignASN1(rand.Reader, decoded, hash[:])
	if err != nil {
		t.Fatal(err)
	}

	valid = ecdsa.VerifyASN1(public, hash[:], sig2)

	if !valid {
		t.Fatal("Failed to verify new signature with old key")
	}

	t.Log("Verified new signature with old key")

	valid = ecdsa.VerifyASN1(&decoded.PublicKey, hash[:], sig)

	if !valid {
		t.Fatal("Failed to verify old signature with decoded key")
	}

	t.Log("Verified old signature with decoded key")
}
