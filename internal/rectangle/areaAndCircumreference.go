package rectangle

func AreaAndCircumference(panjang int8, lebar int8) (area int16, Circumference int16) {
	return int16(panjang * lebar), int16(2 * panjang * lebar)
}
