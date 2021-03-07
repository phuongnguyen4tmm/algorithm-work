package main

import (
	"fmt"
	"math"
)

func main() {
	str := "1001"
	fmt.Println(binary2Decimal(str))
	fmt.Println(binary2Decimal2(str))
}

// dùng số mũ
func binary2Decimal(input string) int32 {
	result := float64(0)

	strLen := len(input)
	for i := 0; i < len(input); i++ {
		power := float64(strLen - i - 1)
		number := float64(input[i] - 48)
		result += math.Pow(float64(2), power) * number
	}

	return int32(result)
}

// dùng nhân đôi
func binary2Decimal2(input string) int32 {
	result := int32(0)
	for _, char := range input {
		number := char - 48
		result = result*2 + number
	}

	return result
}
