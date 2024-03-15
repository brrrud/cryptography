package encryptionModes

type BlockCipherMode int

const (
	ECB BlockCipherMode = iota
	CBC
	PCBC
	CFB
	OFB
	CTR
	RANDOM_DELTA
)
