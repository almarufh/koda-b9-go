package openfiles

import (
	"fmt"
	"io"
	"os"
)

func Open(url string) (result string, err error) {
	file, err := os.Open(url)

	if err != nil {
		return "git ", err
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	defer file.Close()
	cnt, err := io.ReadAll(file)

	if err != nil {
		panic("Terjadi panic saat mebaca file !")
	}

	// fmt.Println(string(cnt))

	return string(cnt), nil
}
