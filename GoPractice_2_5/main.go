package main

import (
	"GoPractice_2_5/popcount"
	"fmt"
)

func main() {
	var x uint64 = 101212323
	result := popcount.PopCountDrop(x)
	fmt.Println("Количество установленных битов в числе:", result)
}
