package rsa

import (
	"errors"
	"github.com/sxperlinx/Go-RSA/math/modular"
)

var ErrorFailedToEncrypt = errors.New("Failed to encrypt")
var ErrorFailedToDecrypt = errors.New("Failed to decrypt")

func Encrypt(message []rune, publicExponent, modulus int64) ([]rune, error) {
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), publicExponent, modulus)
		if err != nil {
			return nil, ErrorFailedToEncrypt
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}

	return encrypted, nil
}

func EncryptPub(message []rune, publicKey Key) ([]rune, error) {
	return Encrypt(message, publicKey.Exponent, publicKey.Modulus)
}

func Decrypt(encrypted []rune, privateExponent, modulus int64) (string, error) {
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedLetter, err := modular.Exponentiation(int64(letter), privateExponent, modulus)
		if err != nil {
			return "", ErrorFailedToDecrypt
		}
		decrypted = append(decrypted, rune(decryptedLetter))
	}

	return string(decrypted), nil
}

func DecryptPriv(encrypted []rune, privateKey Key) (string, error) {
	return Decrypt(encrypted, privateKey.Exponent, privateKey.Modulus)
}