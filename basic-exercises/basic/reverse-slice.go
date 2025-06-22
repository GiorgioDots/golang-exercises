package basic

import (
	"fmt"
)

func ReverseSlice() {
	slice := make([]int, 21)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println("Before: ", slice)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println("After: ", slice)
}
