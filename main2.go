package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	fmt.Println("=====******=====")

	//tingi := 7
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	fmt.Println("=====******=====")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2], kalimat[3], kalimat[4], kalimat[5], kalimat[6])

	fmt.Println("=====******=====")

	var sayuran = []string{}
	vegetables := append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i := 0; i < len(vegetables); i++ {
		fmt.Printf("%d. %s\n", i+1, vegetables[i])
	}
	fmt.Println("=====******=====")

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	volume := 1
	for s, i := range satuan {
		fmt.Printf("%s: %d\n", s, i)
		volume = volume * i
	}

	fmt.Println("volume balok: ", volume)
}
