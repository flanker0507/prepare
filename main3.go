package main

import (
	"fmt"
)

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}
func introduce(name string, kelamin string, job string, age string) string {
	panggil := "Pak"
	if kelamin == "perempuan" {
		panggil = "Bu"
	}

	return fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggil, name, job, age)
}

func buahFavorit(name string, buah ...string) string {
	output := fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah ", name)
	for i, s := range buah {
		output += fmt.Sprintf("\"%s\"", s)
		if i < len(buah)-1 {
			output += ", "
		}
	}
	return output
}
func main() {

	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	fmt.Println("=====***1***=====")

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	fmt.Println("=====***2***=====")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	fmt.Println("=====***3***=====")

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm1 {
		fmt.Println(item)
	}

}

var dataFilm1 = []map[string]string{}

// buatlah closure function disini
func tambahDataFilm(tittle, durasi, genre, tahun string) {
	film := map[string]string{"genre": genre, "jam": durasi, "tahun": tahun, "title": tittle}
	dataFilm1 = append(dataFilm1, film)
}
