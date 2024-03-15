package utils

import "fmt"

func SetBit(byteArray []byte, pos uint64, value bool) {
	if value {
		byteArray[(pos-1)/8] |= (byte)(1 << (7 - (pos-1)%8))
	} else {
		byteArray[(pos-1)/8] &= (byte)(^(1 << (7 - (pos-1)%8)))
	}
}

func GetBit(byteArray []byte, pos uint64) bool {
	return (byteArray[(pos-1)/8] & (1 << (7 - (pos-1)%8))) != 0
}

func ByteToBinaryRepresentationString(byteElem byte) string {
	return fmt.Sprintf("%08b", byteElem)
}

func GetStringBitArrayRepresentation(byteArray []byte) string {
	binaryString := ""
	if byteArray != nil {
		for _, b := range byteArray {
			binaryString += "[" + ByteToBinaryRepresentationString(b) + "]"
		}
	}
	return binaryString
}

func BitsArray8ToByte(bitsArray []byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		result |= bitsArray[i] << uint(7-i)
	}
	return result
}
