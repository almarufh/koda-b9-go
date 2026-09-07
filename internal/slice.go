package internal

import (
	"fmt"
	"slices"
)

func InsertSlice() []int8 {
	oldSlice := []int8{50, 75, 66, 20, 32, 90}
	findIndex := slices.Index(oldSlice, 66) + 1
	leftSlice := oldSlice[:findIndex]
	leftSlice = slices.Grow(leftSlice, len(oldSlice))
	leftSlice = append(leftSlice, 88)
	finalSlice := append(leftSlice, oldSlice[findIndex:]...)

	for i, v := range finalSlice {
		fmt.Printf("%d. %d\n", i, v)
	}

	return finalSlice
}
