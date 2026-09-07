package main

import "fmt"

func hitungLuas(panjang int8, lebar int8) int16 {
	return int16(panjang * lebar)
}

func hitungKeliling(panjang int8, lebar int8) int16 {
	return int16(2 * panjang * lebar)
}

func luasKeliling(panjang int8, lebar int8) (luas int16, keliling int16) {
	return int16(panjang * lebar), int16(2 * panjang * lebar)
}

func main() {
	const panjang int8 = 4
	const lebar int8 = 8
	luas := hitungLuas(panjang, lebar)
	keliling := hitungKeliling(panjang, lebar)
	l, k := luasKeliling(panjang, lebar)
	fmt.Printf("Luas : %d\n", luas)
	fmt.Printf("Keliling : %d\n", keliling)
	fmt.Printf("L : %d\nK : %d ", l, k)
}
