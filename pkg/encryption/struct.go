package encryption

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

func EncryptStruct(key []byte, message any) (string, error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(message)
	if err != nil {
		return "", err
	}

	keyStr := hex.EncodeToString(key)

	cryptoText, err := encrypt([]byte(keyStr), string(network.Bytes()))
	if err != nil {
		return "", err
	}
	return cryptoText, nil
}

func DecryptStruct(key []byte, message string) (any, error) {
	keyStr := hex.EncodeToString(key)

	text, err := decrypt([]byte(keyStr), message)
	if err != nil {
		return nil, err
	}
	byteBuffer := bytes.NewBuffer([]byte(text))
	dec := gob.NewDecoder(byteBuffer)
	var decodedMessage Auth
	err = dec.Decode(&decodedMessage)
	if err != nil {
		return nil, err
	}
	return decodedMessage, nil
}
