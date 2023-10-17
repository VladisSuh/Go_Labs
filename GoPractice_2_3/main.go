package main

import (
	"GoPractice_2_3/popcount"
	"fmt"
)

func main() {
	var x uint64 = 101212323
	result := popcount.PopCount(x)
	fmt.Println("Количество установленных битов в числе:", result)
}
