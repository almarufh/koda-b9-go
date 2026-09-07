package internal

import "fmt"

func StarWindow(panjang int8, lebar int8) {
	for i := range panjang {
		for j := range lebar {
			if i == 0 || i == panjang-1 || j == 0 || j == lebar-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
