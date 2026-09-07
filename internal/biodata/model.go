package biodata

type Pendidikan struct {
	Nama    string
	Jurusan string
}

type Biodata struct {
	Nama       string
	Foto       string
	Email      string
	Umur       int8
	Telepon    string
	Status     string
	Pendidikan Pendidikan
}
