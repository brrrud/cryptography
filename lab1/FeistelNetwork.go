package lab1

import "cryptography/lab1/interfaces"

type FeistelNetwork struct {
	keyExpander interfaces.KeyExpander
	encrypter   interfaces.Encrypter
}

func NewFeistelNetwork(keyExpander interfaces.KeyExpander, encrypter interfaces.Encrypter) *FeistelNetwork {
	return &FeistelNetwork{keyExpander: keyExpander,
		encrypter: encrypter,
	}
}

func (*FeistelNetwork) Encrypt(input []byte, key []byte) []byte {
	return nil
}
