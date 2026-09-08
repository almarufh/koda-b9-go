package payments

import "fmt"

type Fiktif struct {
	Total []uint32
}

func (o *Fiktif) newTFiktif(total []uint32) {
	o.Total = append(o.Total, total...)
}

func (o *Fiktif) Payment(listPrice []uint32) error {

	if len(listPrice) < 1 {
		return fmt.Errorf("List tagihan kosong")
	}

	var tagihan uint32 = 0
	for _, v := range listPrice {
		tagihan += v
	}

	if tagihan <= 0 {
		return fmt.Errorf("Tagihan sama dengan atau kurang dari 0 (%d)", tagihan)
	}
	o.newTFiktif(listPrice)
	return nil
}

func (o *Fiktif) CetakStrukFiktif() {

	var tagihan uint32

	for _, v := range o.Total {
		tagihan += v
	}

	fmt.Println(Struk(tagihan, "FIKTIF"))
}
