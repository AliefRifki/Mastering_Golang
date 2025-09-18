package main

import "fmt"

const word = "Hello Dunia" // isi data tidak bisa di ubah

func main() {
	var l bool = true // var [nama] tipedata
	fmt.Println("this is bool ")
	fmt.Println(true)
	fmt.Println()

	fmt.Println("this is string with double tick")
	fmt.Println("Hello World!")
	fmt.Println()

	fmt.Println("this is string with backtick")
	fmt.Println(`AWOWKAWOWAK
	a
	l
	i
	e
	f`)

	fmt.Println()
	fmt.Println("Nama ku Alief, Umurku", 20, "Tahun")

	fmt.Printf("iam %v %T \n", true, true) // v untuk melihat nilai dari tipe data apapun, T untuk melihat nama tipe data

	fmt.Printf("Aku semester %d \n", 5)            // %d untuk integer decimal
	fmt.Printf("iam binary %b\n", 5)               // %b untuk integer to binary
	fmt.Printf("\tAlief DA GOAT %t\n", l)          // %t untuk print nilai boolean
	fmt.Printf("GPA ku semester ini %.2f\n", 3.02) // %f untuk float

	var kata = "Alief Ganteng" // go lang bisa langsung menentukan tipe data tanpa harus di specify dengan syarat harus memberi nilai secara langsung
	fmt.Println(kata)

	var kata1, kata2 string = "THE", "GOAT" // jika membuat seperti ini dan di specify tipe data nya maka nilai nya harus sama
	fmt.Println(kata1, kata2)

	var (
		yap1 = 12
		yap2 = "Kerkez"
	)

	fmt.Println(yap1, yap2)

	n := 5
	fmt.Println(n)

	// nama variabel mempengaruhi penggunaan, jika menggunakan huruf kapital bisa di gunakan di package lain?
	// var alief = "ganteng" -> ini hanya bisa di akses di package main
	// var Alief = "ganteng" -> bisa diakses di package lain

	fmt.Println(word)
	const lel = 404
	fmt.Println(lel)
}
