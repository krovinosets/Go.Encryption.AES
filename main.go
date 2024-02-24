package main

import (
	"Encryption/pkg/encryption"
	"fmt"
)

const (
	Key = "17~jiKH3y8B0cKtN"
)

func main() {
	key := []byte(Key)

	en, err := encryption.EncryptStruct(key, encryption.Auth{UserId: 1003})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(en)

	de, err := encryption.DecryptStruct(key, "kCMsU3HQtCt22cPWPQCRS8DimqjLsYEXfkyZLPMw5lRJk3xqSfWmRBKxBjj9H8Fp4SDoons")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(de)

	de, err = encryption.DecryptStruct(key, "dQ3/94TAK7wfHskEjtU8+LPm5EupI6l4w0O9WRVsRZa6MVjwI8x+9BhdeH9olrJowJaqQ/w")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(de)

	de, err = encryption.DecryptStruct(key, "Dvlulsr2Pi8zUgOvbgtNLBsSLutqFQYirdxGPaGS2V+YKaYbRf2+WzQlmQKITzajQEO2caQ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(de)
}
