package interanal

import (
	"fmt"
	"slices"
)

func InsertSlice() {
	oldSlice := []int8{50, 75, 66, 20, 32, 90}
	findIndex := slices.Index(oldSlice, 66) + 1
	leftSlice := oldSlice[:findIndex]
	leftSlice = slices.Grow(leftSlice, len(oldSlice))
	leftSlice = append(leftSlice, 88)
	oldSlice = append(leftSlice, oldSlice[findIndex:]...)

	for i, v := range oldSlice {
		fmt.Printf("%d. %d\n", i, v)
	}
}
