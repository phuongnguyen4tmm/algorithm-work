package main

import "fmt"

func main() {
	fmt.Println(simpleSolution(10))
	fmt.Println(betterSolution(10))
}

func betterSolution(num int) []int {
	if num == 0 {
		return []int{0}
	}

	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = res[i/2] + i%2
	}

	return res
}

func simpleSolution(num int) []int {
	if num == 0 {
		return []int{0}
	}

	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		res[i] = countBit1(i)
	}

	return res
}

func countBit1(num int) int {
	var count int

	for num > 0 {
		count += num & 1
		num >>= 1
	}

	return count
}
