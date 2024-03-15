package encryptionModes

type ContextInfo struct {
	key         []byte
	blockMode   BlockCipherMode
	paddingMode PaddingMode
	ivVector    []byte
}

func (cntInfo ContextInfo) Encrypt(plaintext []byte) ([]byte, error) {
	return nil, nil
}

func (cntInfo ContextInfo) Decrypt(ciphertext []byte) ([]byte, error) {
	return nil, nil
}

func (cntInfo ContextInfo) SetKey(key []byte) error {
	return nil
}

func NewContextInfo(key []byte, blockMode BlockCipherMode, paddingMode PaddingMode, ivVector []byte, anotherVariadic ...any) *ContextInfo {
	return &ContextInfo{key: key,
		blockMode:   blockMode,
		paddingMode: paddingMode,
		ivVector:    ivVector}
}
