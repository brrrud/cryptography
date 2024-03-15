package lab1

import "cryptography/utils"

func PermuteBits(input []byte, pbox []uint64, mode PBLOCK_MODE) []byte {
	var ans []byte
	for j, newPos := range pbox {
		utils.SetBit(ans, newPos+1, utils.GetBit(input, uint64(j)))
	}
	return ans
}
