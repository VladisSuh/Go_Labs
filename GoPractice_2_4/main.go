package main

import (
	popcount "GoPractice_2_4/PopCount"
	"fmt"
)

func main() {
	var x uint64 = 101212323
	result := popcount.NewPopCount(x)
	fmt.Println("Количество установленных битов в числе:", result)
}
