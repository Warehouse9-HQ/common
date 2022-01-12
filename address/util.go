package address

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"math/big"
	"strings"
)

func AddressToPublic(c elliptic.Curve, key string) (*ecdsa.PublicKey, error) {
	if !strings.HasPrefix(key, "0s") {
		return nil, ErrInvalidAddress
	}
	key = key[2:]
	pubBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	var public = new(ecdsa.PublicKey)
	public.X, public.Y = elliptic.Unmarshal(c, pubBytes)
	public.Curve = c
	return public, nil
}

func PublicToAddress(public *ecdsa.PublicKey) string {
	return "0s" + base64.StdEncoding.EncodeToString(elliptic.Marshal(public.Curve, public.X, public.Y))
}

func EncodePrivate(private *ecdsa.PrivateKey) string {
	return base64.StdEncoding.EncodeToString(private.D.Bytes())
}

func DecodePrivate(c elliptic.Curve, key string) (*ecdsa.PrivateKey, error) {
	kBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	k := new(big.Int).SetBytes(kBytes)
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve, priv.D = c, k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}
