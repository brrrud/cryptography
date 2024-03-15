package encryptionModes

type PaddingMode int

const (
	ZEROS PaddingMode = iota
	ANSIX923
	PKCS7
	ISO10126
)
