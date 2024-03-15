package interfaces

// интерфейс, предоставляющий описание функционала по выполнению шифрования и дешифрования
// симметричным алгоритмом (параметр методов: [де]шифруемый блок (массив байтов))
// с преднастроенными отдельным методом раундовыми ключами (параметр метода:
// ключ [де]шифрования (массив байтов));

type SymmetricEncryption interface {
	SetKey(key []byte) error
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}
