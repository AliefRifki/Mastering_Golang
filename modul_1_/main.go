package main

import (
	"fmt"
)

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

func main() {
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
}
