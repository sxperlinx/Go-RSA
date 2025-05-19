package rsa

import (
	"log"

	"github.com/sxperlinx/Go-RSA/math/gcd"
	"github.com/sxperlinx/Go-RSA/math/modular"
)

type Key struct {
	Exponent int64
	Modulus  int64
}

type KeyPair struct {
	PrivateKey Key
	PublicKey  Key
}

func GenerateKeyPair() (*KeyPair, error) {
	var p int64 = 7919
	var q int64 = 1009
	var e int64

	n := p * q
	phi := (p - 1) * (q - 1)

	for e = 2; e < phi; e++ {
		if (gcd.Iterative(e, phi) == 1) && (gcd.Recursive(e, phi) == 1) {
			break
		}
	}

	d, err := modular.Inverse(e, phi)

	if err != nil {
		log.Fatalf("Error calculating modular inverse: %v", err)
		return nil, err
	}

	return &KeyPair{
		PrivateKey: Key{ d, n },
		PublicKey:  Key{ e, n },
	}, nil
}
