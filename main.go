package main

import (
	"fmt"
	"os"

	"github.com/almarufh/koda-b9-go/internal"
	"github.com/almarufh/koda-b9-go/internal/biodata"
	"github.com/almarufh/koda-b9-go/internal/rectangle"
	"github.com/almarufh/koda-b9-go/utils"
)

func wrongInput(text string) {
	utils.Clear()
	fmt.Printf("\n\nWarning...!!!\n\n")
	fmt.Printf("Not found menu in (%s)", text)
	confirm()
}

func confirm() {
	fmt.Printf("\n\n1. Back to dashboard\n0. Exit")
	fmt.Printf("\n\nChoose a menu : ")
	input := utils.Input()
	switch input {
	case "1":
		return
	case "0":
		utils.Clear()
		os.Exit(0)
	default:
		wrongInput(input)
	}
}

func main() {
	for {
		utils.Clear()
		menu := []string{
			"Calculate Rectangle Area",
			"Calculate Rectangle Perimeter",
			"Calculate Rectangle Area and Perimeter",
			"Create Window",
			"Insert Slice",
			"My Biodata",
		}

		fmt.Printf("\n\n---[ DASHBOARD MINI TASK GOLANG ]---\n\n")
		for i, v := range menu {
			fmt.Printf("%d. %s\n", i+1, v)
		}

		fmt.Printf("\n0. Exit\n\n")

		fmt.Printf("Choose a menu : ")
		input := utils.Input()
		switch input {
		case "1":
			utils.Clear()
			luas := rectangle.Area(8, 8)
			fmt.Printf("Luas : %d\n", luas)
			confirm()
			continue
		case "2":
			utils.Clear()
			keliling := rectangle.Circumference(8, 8)
			fmt.Printf("Keliling : %d", keliling)
			confirm()
			continue
		case "3":
			utils.Clear()
			l, k := rectangle.AreaAndCircumference(8, 8)
			fmt.Printf("Luas : %d\n", l)
			fmt.Printf("Keliling : %d\n", k)
			confirm()
			continue
		case "4":
			utils.Clear()
			internal.StarWindow(8, 6)
			confirm()
			continue
		case "5":
			utils.Clear()
			internal.InsertSlice()
			confirm()
			continue
		case "6":
			utils.Clear()
			biodata.MyBiodata()
			confirm()
			continue
		case "0":
			utils.Clear()
			os.Exit(0)
		default:
			wrongInput(input)
			continue
		}
	}
}
