package concurrency

import (
	"fmt"
	"sync"
)

type sender struct {
	name    string
	message string
}

func SendMessage() {
	wg := sync.WaitGroup{}
	messages := []sender{
		{
			name:    "Fajar",
			message: "Selamat malam",
		},
		{
			name:    "Dimas",
			message: "Selamat Pagi",
		},
		{
			name:    "Fajar",
			message: "Selamat malam",
		},
		{
			name:    "Dimas",
			message: "Selamat Pagi",
		},
	}

	board := make(chan sender, len(messages))

	wg.Go(func() { PrintMessage(board) })

	for _, v := range messages {
		board <- v
	}

	close(board)

	wg.Wait()

	fmt.Println("Seluruh pesan berhasil dikirim dan di catat di papan tulis")
}

func PrintMessage(board chan sender) {
	defer fmt.Printf("\n\nSelesai mencetak seluruh pesan")
	fmt.Printf("\nSiap menerima dan menctat kedalam papan tulis\n\n")
	for data := range board {
		fmt.Printf("Sender  : %s\n", data.name)
		fmt.Printf("Message : %s\n", data.message)
		fmt.Println()
	}
}
