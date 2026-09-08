package payments

import "fmt"

type Bank struct{}

func (o Bank) Payment(listPrice []uint32) error {
	var tagihan uint32

	for _, value := range listPrice {
		tagihan += value
	}

	fmt.Println(Struk(tagihan, "BANK"))

	return nil
}
