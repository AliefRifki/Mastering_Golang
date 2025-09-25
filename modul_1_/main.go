package main

import (
	"errors"
	"fmt"
)

var errorNol error = errors.New("tidak bisa membagi dengan nol")

// anonymous function,  fungsi yang di masukkan ke dalam variabel
var cetakHello = func() {
	fmt.Println("Hello World")
}

// fungsi sebagai parameter
func operasiAngka(angka int, fx func(int) int) {
	var hasil = fx(angka)

	fmt.Println(hasil)
}

func kuadrat(angka int) int {
	return angka * angka
}

func kubik(angka int) int {
	return angka * angka * angka
}

func cekSaldo() int {
	return 100000
}

// Error adalah variable yang ada di golang, biasanya error akan ada di belakang return pertama
func bagiAngka(x int, y int) (int, error) {
	if y == 0 {
		return 0, errorNol
	}

	if y < 0 {
		return 0, errors.New("tidak mendukung pembagian dengan angka negatif")
	}
	return x / y, nil
}

func bagiPanic(x, y int) int {
	if y == 0 {
		panic("Tidak bisa membagi dengan nol")
	}

	return x / y
}

func main() {
	// defer akan jalan terakhir setelah program selesai, dan ini berfungsi untuk merecover panic yang terjadi agar program tidak langsung di terminate (program tidak langsung force close) atau namanya adalah graceful shutdown
	defer func() {
		r := recover()
		fmt.Println("\nPanic berhasil di recover", r)
	}() // immidiate invoke func

	cetakHello()

	// immidiately invoke function, setelah fungsi dibuat langsung di panggil
	func() {
		fmt.Println("Halo Nigaz")
	}()

	func(nama string) {
		fmt.Println("Hi,", nama, "ini nigga ku")
	}("Rusdi")

	// fungsi sebagai parameter
	fmt.Println("Kuadrat dari 5 : ")
	operasiAngka(5, kuadrat)
	fmt.Println("Kubik dari 5 : ")
	operasiAngka(5, kubik)

	// scope variable
	var x = 5

	{
		var x = 7
		fmt.Println(x) // ini akan berisi 7 dan hanya disini berisi 7 (lokal)
	}

	fmt.Println(x) // ini akan menghasilkan 5

	// conditional statement, di Golang ada 2 yaitu if dan switch
	// const saldoMinimum = 50
	// const saldoSekarang = 100
	// if saldoSekarang >= saldoMinimum {
	// 	fmt.Println("Tarik tunai berhasil!")
	// } else {
	// 	fmt.Println("Saldo tidak mencukupi")
	// }

	const usernameTerdaftar = true
	const passwordSesuai = true

	if usernameTerdaftar && passwordSesuai {
		fmt.Println("Login berhasil")
	} else if usernameTerdaftar && !passwordSesuai {
		fmt.Println("Password salah")
	} else {
		fmt.Println("Username tidak terdaftar")
	}

	var saldoMinimum = 50000

	if saldoSekarang := cekSaldo(); saldoSekarang >= saldoMinimum {
		fmt.Printf("Tarik tunai berhasil, saldo sebelum ditarik %d, setelah ditarik %d \n", saldoSekarang, saldoSekarang-50000)
	} else {
		fmt.Printf("Saldo tidak mencukupi, sisa saldo %d", saldoSekarang)
	}

	hasilBagi, err := bagiAngka(5, 0)
	if err != nil {
		fmt.Printf("Terjadi kesalahan %s \n", err.Error())
		if errors.Is(err, errorNol) {
			fmt.Println("Jangan gunaan nol, karena menghasilkan tak hingga")
		}
	} else {
		fmt.Printf("Hasil bagi: %d", hasilBagi)
	}

	fmt.Println()
	fmt.Println("Contoh recover panic pada function pembagian")
	hasilPanic := bagiPanic(5, 5)
	fmt.Printf("Hasil bagi : %d \n", hasilPanic)

	// Iteration
	fmt.Println()
	var warna = [...]string{"Merah", "Kuning", "Hijau", "Kelabu"}

	for i := 0; i < len(warna); i++ {
		fmt.Println(warna[i])
	}

	fmt.Println()
	var club = [...]string{"Liverpool", "celeng", "demit"}
	for i := range club {
		fmt.Println(club[i])
	}

	// array
	fmt.Println("\nArray")
	var arr = [...]string{"satu", "dua", "tiga", "empat"}
	fmt.Println(len(arr))

	var i = 0

	// for i != 10 {
	// 	fmt.Println("Nilai I : ", i)
	// 	i++
	// }

	for {
		if i%2 != 0 {
			i++
			continue
		} else if i == 10 {
			break
		}
		fmt.Println("Nilai i: ", i)
		i++
	}

	// Array multidimensi
	fmt.Println()

	// Range pada string
	var kata string = "Golang"

	for _, k := range kata {
		fmt.Printf("Nilai %c Tipe data %T \n", k, k)
	}

	// Slice
	fmt.Println("\nSlice :")

	var color = []string{"Merah", "Kuning", "Hijau"}
	fmt.Println(color)
	var kolor = make([]string, 3, 5)
	fmt.Println(len(kolor))
	fmt.Println(cap(kolor))

	color = append(color, "Biru", "Putih", "Kuning")

	fmt.Println(color)
}
