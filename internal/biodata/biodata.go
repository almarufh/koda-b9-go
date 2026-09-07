package biodata

import "fmt"

func MyBiodata() {
	biodata := Biodata{
		Nama:    "Alma'ruf Hidayat",
		Foto:    "@almarufh",
		Email:   "almarufhidayat99@gmail.com",
		Umur:    27,
		Telepon: "082393468568",
		Status:  "Menikah",
		Pendidikan: Pendidikan{
			Nama:    "Universitas Islam Negeri Palopo",
			Jurusan: "Hukum Ekonomi Syariah",
		},
	}

	fmt.Println(biodata)
}
