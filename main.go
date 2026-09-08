package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/almarufh/koda-b9-go/internal"
	"github.com/almarufh/koda-b9-go/internal/biodata"
	"github.com/almarufh/koda-b9-go/internal/openfiles"
	"github.com/almarufh/koda-b9-go/internal/person"
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

func StringToInt8(input string) int8 {
	val, err := strconv.ParseUint(input, 10, 8)
	if err != nil {
		return 0
	}
	return int8(val)
}

func inputParameter() (panjang int8, lebar int8) {
	for {
		utils.Clear()
		fmt.Printf("\nInput Panjang : ")
		inputPanjang := utils.Input()
		utils.Clear()
		panjang = StringToInt8(inputPanjang)
		fmt.Printf("\nInput Lebar : ")
		inputLebar := utils.Input()
		lebar = StringToInt8(inputLebar)
		utils.Clear()
		fmt.Printf("\nPanjang  : %d", panjang)
		fmt.Printf("\nLebar    : %d\n\n", lebar)
		fmt.Printf("1. Procces\n")
		fmt.Printf("2. Change\n")
		fmt.Printf("\n0. Exit\n")
		fmt.Printf("\nChoose a menu : ")
		menu := utils.Input()
		switch menu {
		case "1":
			utils.Clear()
			return panjang, lebar
		case "2":
			continue
		case "0":
			os.Exit(0)
		default:
			wrongInput(menu)
		}
	}
}

func showFile(rest string, err error) {
	utils.Clear()
	if err != nil {
		fmt.Printf("\n-----[ ERROR ]-----\n\n%s", err)
		fmt.Printf("\n\n-----------------------")
	} else {
		fmt.Printf("\n-----[ RESULTS ]-----\n\n%s", rest)
		fmt.Printf("\n\n-----------------------")
	}
}

func dashboard() {
	for {
		utils.Clear()
		menu := []string{
			"Calculate Rectangle Area",
			"Calculate Rectangle Perimeter",
			"Calculate Rectangle Area and Perimeter",
			"Create Window",
			"Insert Slice",
			"My Biodata",
			"Open File",
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
			p, l := inputParameter()
			luas := rectangle.Area(p, l)
			fmt.Printf("Luas : %d\n", luas)
			confirm()
			continue
		case "2":
			utils.Clear()
			p, l := inputParameter()
			keliling := rectangle.Circumference(p, l)
			fmt.Printf("Keliling : %d", keliling)
			confirm()
			continue
		case "3":
			utils.Clear()
			p, lb := inputParameter()
			l, k := rectangle.AreaAndCircumference(p, lb)
			fmt.Printf("Luas     : %d\n", l)
			fmt.Printf("Keliling : %d\n", k)
			confirm()
			continue
		case "4":
			utils.Clear()
			p, l := inputParameter()
			internal.StarWindow(p, l)
			confirm()
			continue
		case "5":
			utils.Clear()
			internal.InsertSlice()
			confirm()
			continue
		case "6":
			utils.Clear()
			data := biodata.MyBiodata()
			fmt.Printf("\nNama       : %s\n", data.Nama)
			fmt.Printf("Foto       : %s\n", data.Foto)
			fmt.Printf("Email      : %s\n", data.Email)
			fmt.Printf("Umur       : %d\n", data.Umur)
			fmt.Printf("Telepon    : %s\n", data.Telepon)
			fmt.Printf("Status     : %s\n", data.Status)
			fmt.Printf("Pendidikan :\n")
			fmt.Printf("  Kampus   : %s\n", data.Pendidikan.Nama)
			fmt.Printf("  Jurusan  : %s\n", data.Pendidikan.Jurusan)
			confirm()
			continue
		case "7":
			utils.Clear()
			fmt.Printf("\n\n1. Input Manual")
			fmt.Printf("\n2. Readme")
			fmt.Printf("\n3. Directory")
			fmt.Printf("\n\nChoose a menu : ")
			input := utils.Input()
			switch input {
			case "1":
				utils.Clear()
				fmt.Printf("Input path : ")
				text := utils.Input()
				val, err := openfiles.Open(text)
				showFile(val, err)
				confirm()
				continue
			case "2":
				utils.Clear()
				val, err := openfiles.Open("./readme.md")
				showFile(val, err)
				confirm()
				continue
			case "3":
				utils.Clear()
				val, err := openfiles.Open("./internal")
				showFile(val, err)
				confirm()
				continue
			default:
				wrongInput(input)
			}
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

func main() {
	// dashboard()
	dataPersons := person.NewPerson("Ma'ruf", "Palopo", "082393468568")

	data := dataPersons.PrintAll()
	println(data)

	greet := dataPersons.Greet()
	fmt.Println(greet)

	dataPersons.ChangeName("Niko")

	greet = dataPersons.Greet()
	fmt.Println(greet)
}
