package payments

import "fmt"

type Payments interface {
	Payment([]uint32) error
}

func Pay(listPrice []uint32, metode Payments) error {
	return metode.Payment(listPrice)
}

func Struk(total uint32, metode string) string {
	return fmt.Sprintf("----[ PEMBAYARAN %s SUKSES ]----\n\nTotal Pembayaran : %d\n\n-----------------------", metode, total)
}
