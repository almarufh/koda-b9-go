package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func GoWork() {
	defer fmt.Println("Otw ngantor bosku")
	wg := sync.WaitGroup{}
	wg.Go(takeAbath)
	wg.Go(drinkCofee)
	wg.Go(breakFast)
	// wg.Go(cleanRoom)
	wg.Add(1)
	go cleanRoom(&wg)
	wg.Wait()

}

func takeAbath() {
	defer fmt.Println("Selesai mandi")
	fmt.Println("Mulai mandi")
	time.Sleep(20 * time.Millisecond)
}

func drinkCofee() {
	defer fmt.Println("Selesai Ngopi")
	fmt.Println("Mulai Ngopi")
	time.Sleep(20 * time.Millisecond)
}

func breakFast() {
	defer fmt.Println("Selesai sarapan")
	fmt.Println("Mulai menyiapkan sarapan")
	time.Sleep(20 * time.Millisecond)
}

func cleanRoom(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Selesai membereskan kammar tidur")
	fmt.Println("Mulai mebereskan kamar tidur")
	time.Sleep(20 * time.Millisecond)
}
