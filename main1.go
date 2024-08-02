package main

import (
	"fmt"
	"strconv"
)

func main() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	panjangInt, err1 := strconv.Atoi(panjangPersegiPanjang)
	if err1 != nil {
		return
	}

	lebarInt, err2 := strconv.Atoi(lebarPersegiPanjang)
	if err2 != nil {
		return
	}

	alasInt, err3 := strconv.Atoi(alasSegitiga)
	if err3 != nil {
		return
	}

	tinggiInt, err3 := strconv.Atoi(tinggiSegitiga)
	if err3 != nil {
		return
	}

	fmt.Println("panjangPersegiPanjang: ", panjangInt)
	fmt.Println("lebarPersegiPanjang: ", lebarInt)
	fmt.Println("alasSegitiga: ", alasInt)
	fmt.Println("tinggiSegitiga", tinggiInt)

	luasPersegiPanjang = panjangInt * lebarInt
	fmt.Println("luasPersegiPanjang: ", luasPersegiPanjang)

	kelilingPersegiPanjang = 2 * (panjangInt + lebarInt)
	fmt.Println("kelilingPersegiPanjang: ", kelilingPersegiPanjang)

	luasSegitiga = (alasInt * tinggiInt) / 2
	fmt.Println("luasSegitiga: ", luasSegitiga, "cm")

	fmt.Println("=====******=====")

	arr := []int{100, 200, 1000, 300, 200, 75, 78, 200, 253}
	i := 0
	sum := 0

	for sum <= 1000 && i < len(arr) {
		sum += arr[i]
		fmt.Println(sum)
		i++
	}

	fmt.Println("hasil dari sum: ", sum)
	fmt.Println("berhenti di index: ", i-1)

	fmt.Println("=====******=====")

	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiDoe >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("indeksnya c")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("indeksnya D")
	} else {
		fmt.Println("indeksnya E")
	}

	if nilaiJohn >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeksnya c")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeksnya D")
	} else {
		fmt.Println("indeksnya E")
	}

	fmt.Println("=====******=====")

	tanggal := 5
	bulan := 7
	tahun := 2000

	switch bulan {
	case 1:
		fmt.Printf("%d Januari %d", tanggal, tahun)
	case 2:
		fmt.Printf("%d February %d", tanggal, tahun)
	case 3:
		fmt.Printf("%d Maret %d", tanggal, tahun)
	case 4:
		fmt.Printf("%d April %d", tanggal, tahun)
	case 5:
		fmt.Printf("%d May %d", tanggal, tahun)
	case 6:
		fmt.Printf("%d June %d", tanggal, tahun)
	case 7:
		fmt.Printf("%d July %d\n", tanggal, tahun)
	case 8:
		fmt.Printf("%d August %d", tanggal, tahun)
	case 9:
		fmt.Printf("%d September %d", tanggal, tahun)

	}

	fmt.Println("=====******=====")

	var kelahiran int = 2024

	switch {
	case kelahiran >= 1944 && kelahiran <= 1964:
		fmt.Println("Baby boomer")
	case kelahiran >= 1956 && kelahiran <= 1979:
		fmt.Println("Generasi X")
	case kelahiran >= 1980 && kelahiran <= 1994:
		fmt.Println("Generasi Y (Millenials)")
	case kelahiran >= 1995 && kelahiran <= 2015:
		fmt.Println("Generasi Z")
	default:
		fmt.Println("TAHUN LAHIR ANDA TIDAK ADA")
	}
}

//Baby boomer, kelahiran 1944 s.d 1964
//Generasi X, kelahiran 1965 s.d 1979
//Generasi Y (Millenials), kelahiran 1980 s.d 1994
//Generasi Z, kelahiran 1995 s.d 2015
