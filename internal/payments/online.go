package payments

import "fmt"

type Online struct{}

func (o Online) Payment(listPrice []uint32) error {
	var tagihan uint32

	for _, value := range listPrice {
		tagihan += value
	}

	fmt.Println(Struk(tagihan, "ONLINE"))

	return nil
}
