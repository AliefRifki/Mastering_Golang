package main

import "fmt"

const word = "Hello Dunia" // isi data tidak bisa di ubah
const (
	JENIS_KELAMIN_LAKI      = "Laki-Laki"
	JENIS_KELAMIN_PEREMPUAN = "Perempuan"
) // constant hanya bisa digunakan dengan tipe dasar

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

	fmt.Println(JENIS_KELAMIN_LAKI)
	fmt.Println(JENIS_KELAMIN_PEREMPUAN)

	// Pointer
	var kalimat = "Hello World"
	var pointerSaya *string = &kalimat // zero value dari pointer itu nil, dan nil bisa membuat panic dalam bahasa golang (program berhenti karena ada nya error)

	*pointerSaya = "Programming golang"

	fmt.Println(*pointerSaya)
	fmt.Println(kalimat)

	var pointerNew = new(string) // zero value dalam pointer bisa saja aman jika kita menggunakan metode new() dan zero value nya akan berisi sesuai dari tipe data yang di assign, misal new(string) maka zero value nya adalah "", jika new(int) zero valuenya adalah 0 dst.
	fmt.Printf("%v\n", pointerNew)

	// Operator
	var nilai = 5 / 2
	fmt.Println(nilai) // hasilnya akan dibulatkan

	var modNilai = 5 % 2
	fmt.Println(modNilai) // Modulus (%) akan menampilkan sisa bagi jadi 5/2 memiliki sisa 1 maka hasilnya akan 1

	var x = 5
	var y = 5
	x++
	y--
	fmt.Println("Incremet & Decrement x = 5 x-- =", x, "y = 5 y-- =", y)

}
