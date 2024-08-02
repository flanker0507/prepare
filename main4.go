//package main
//
//import "fmt"
//
//func hitungLuas(luasLigkaran, kelilingLingkaran *float64, jari float64) {
//	const Pi = 3.14
//	*luasLigkaran = Pi * jari * jari
//	*kelilingLingkaran = 2 * jari * jari
//}
//
//func introducee(sentence *string, name, gender, work, age string) {
//	panggil := "Pak"
//	if gender == "perempuan" {
//		panggil = "Bu"
//	}
//
//	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggil, name, work, age)
//}
//
//func fruits(buah *[]string, tamahBuah ...string) {
//	for _, s := range tamahBuah {
//		*buah = append(*buah, s)
//	}
//}
//
//func main() {
//	fmt.Println("=====***1***=====")
//
//	var luasLigkaran float64
//	var kelilingLingkaran float64
//
//	hitungLuas(&luasLigkaran, &kelilingLingkaran, 5)
//
//	fmt.Printf("Luas Lingkaran adalah: %.2f\n", luasLigkaran)
//	fmt.Printf("Keliling Lingkaran adalah: %.2f\n", kelilingLingkaran)
//
//	fmt.Println("=====***2***=====")
//
//	var sentence string
//	introducee(&sentence, "John", "laki-laki", "penulis", "30")
//
//	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
//	introducee(&sentence, "Sarah", "perempuan", "model", "28")
//
//	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
//
//	fmt.Println("=====***3***=====")
//
//	var buah = []string{}
//
//	fruits(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
//
//	for i, s := range buah {
//		fmt.Printf("%d. %s\n", i+1, s)
//	}
//
//	fmt.Println("=====***4***=====")
//
//	tambahDataFilmm("LOTR", "2 jam", "action", "1999", &dataFilm) //Cannot use '&dataFilm' (type *[]map[string]string) as the type *[]string
//	tambahDataFilmm("avenger", "2 jam", "action", "2019", &dataFilm)
//	tambahDataFilmm("spiderman", "2 jam", "action", "2004", &dataFilm)
//	tambahDataFilmm("juon", "2 jam", "horror", "2004", &dataFilm)
//
//	for i, m := range dataFilm {
//		fmt.Printf("%d.Title: %s,\n Duration: %s,\n Genre: %s,\n Year: %s\n", i+1, m["title"], m["durasi"], m["genre"], m["tahun"])
//
//	}
//
//}
//
//var dataFilm = []map[string]string{}
//
//func tambahDataFilmm(tittle, durasi, genre, tahun string, dataFilm *[]map[string]string) {
//	film := map[string]string{
//		"title":  tittle,
//		"durasi": durasi,
//		"genre":  genre,
//		"tahun":  tahun,
//	}
//
//	*dataFilm = append(*dataFilm, film)
//
//}
