package main

import (
	"fmt"
)

type buah struct {
	name, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	luas := (s.alas * s.tinggi) / 2
	return luas
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

func main() {
	fmt.Println("=====***1***=====")

	fruit := []buah{
		{name: "nanas", warna: "kuning", adaBijinya: false, harga: 9000},
		{name: "Jeruk", warna: "oranye", adaBijinya: true, harga: 8000},
		{name: "semangka", warna: "hijau & merah", adaBijinya: true, harga: 10000},
		{name: "pisang", warna: "kuning", adaBijinya: false, harga: 5000},
	}

	for _, b := range fruit {
		fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %t, Harga: %d\n", b.name, b.warna, b.adaBijinya, b.harga)
	}
	fmt.Println("=====***2***=====")

	sg := segitiga{10, 10}
	fmt.Println("Luas Segitiga", sg.luas())

	p := persegi{10}
	fmt.Println("Luas Perdegi", p.luas())

	pp := persegiPanjang{10, 23}
	fmt.Println("Luas Persegi panjang", pp.luas())

	fmt.Println("=====***3***=====")

	ph := phone{
		name:   "iPhone X",
		brand:  "Apple",
		year:   2017,
		colors: []string{"Black", "White"},
	}

	// Menambahkan warna ke objek phone
	ph.addColor("Red")
	ph.addColor("Blue")

	// Mencetak informasi phone termasuk warna-warna yang telah ditambahkan
	fmt.Printf("Name: %s\n", ph.name)
	fmt.Printf("Brand: %s\n", ph.brand)
	fmt.Printf("Year: %d\n", ph.year)
	fmt.Printf("Colors: %v\n", ph.colors)

	fmt.Println("=====***4***=====")

	tambahDataFilm1("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm1("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm1("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm1("juon", 120, "horror", 2004, &dataFilm)

	for _, film := range dataFilm {
		fmt.Printf("Title: %s, Duration: %d min, Genre: %s, Year: %d\n",
			film.title, film.duration, film.genre, film.year)

	}

}

type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

var dataFilm = []movie{}

// Fungsi untuk menambahkan data film ke slice dataFilm menggunakan pointer
func tambahDataFilm1(title string, duration int, genre string, year int, dataFilm *[]movie) {
	film := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}

	*dataFilm = append(*dataFilm, film)
}
