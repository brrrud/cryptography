package interfaces

type IFeistelFunction interface {
	FeistelFunction(block []byte, key []byte) []byte
}
