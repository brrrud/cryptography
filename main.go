package main

import (
	"cryptography/utils"
	"fmt"
)

// / сеть фейстеля на основе интерфейса 3
func permuteBits(num uint8, positions []int) uint8 {
	result := uint8(0)

	for i, pos := range positions {
		bit := (num >> uint(7-pos)) & 1 // Получаем бит на позиции pos из числа num
		result |= bit << uint(7-i)      // Ставим бит на соответствующую позицию в результирующем числе
	}

	return result
}

func main() {
	//input := uint8(105)                          // Пример входного числа: 01101001
	//permutation := []int{0, 6, 5, 3, 2, 4, 1, 7} // Пример перестановки битов
	//
	//output := permuteBits(input, permutation)
	//
	//fmt.Printf("Input:  %08b\n", input)
	//fmt.Printf("Output: %08b\n", output)

	// 0 1 1 0 1 0 0 1
	// 0 1 2 3 4 5 6 7
	// 0 0 0 0 1 1 1 1

	arr := []byte{1, 0, 1, 0, 1, 0, 1, 0}
	arr1 := []byte{0, 1, 0, 1, 1, 0, 1, 1}

	byte1 := utils.BitsArray8ToByte(arr)
	sl := []byte{byte1}
	sl = append(sl, utils.BitsArray8ToByte(arr1))

	fmt.Println(utils.GetStringBitArrayRepresentation(sl))
	//lab1.PermuteBits(sl, []uint64{1, 2, 3}, lab1.NUMERATE_FROM_START)
	//fmt.Println(utils.GetStringBitArrayRepresentation(sl))

	//
	//utils.SetBit(sl, 14, true)
	//
	fmt.Println(utils.GetStringBitArrayRepresentation(sl))
	//
	//fmt.Println(utils.GetBit(sl, 5))

	//fmt.Println(utils.GetBit(arr))
	//fmt.Printf("0x%X\n", utils.BitsArray8ToByte(arr)) //0xAA
	//utils.GetStringBitArrayRepresentation([]byte{utils.BitsArray8ToByte(arr)})
	//utils.SetBit(arr, 2, true)
	//fmt.Println(arr)
	//utils.GetStringBitArrayRepresentation(By)
}
