package main

import (
	"errors"
	"fmt"
)

var errorNol error = errors.New("tidak bisa membagi dengan nol")

type BangunRuangI interface {
	luas() float64
}

type Lingkaran struct {
	diameter float64
}

type Persegi struct {
	sisi float64
}

type Segitiga struct {
	tinggi float64
	alas   float64
}

func (I Lingkaran) luas() float64 {
	r := 0.5 * I.diameter
	return 3.14 * r * r
}

func (s Segitiga) luas() float64 {
	return 0.5 * s.alas * s.tinggi
}

func (p Persegi) luas() float64 {
	return p.sisi * p.sisi
}

// embedded struct
type Geometri struct {
	luas     int
	keliling int
}

// jika struct menggunakan huruf kapital diawalnya maka dia bisa diakses di package yang berbeda, namun jika penamaan field nya tidak menggunakan kapital diawal maka fiel tersebut tidak bisa diakses dari package yang berbeda
type BangunRuang struct {
	Geometri
	panjang int
	lebar   int // panjang & lebar tidak bisa diakses dari package yg berberda karena tidak menggunakan kapital di depan nya
}

func (b BangunRuang) hitungLuas() int {
	return b.panjang * b.lebar
}

// Ini adalah fungsi dengan pointer receiver, fungsi ini bisa mengubah nilai yang disimpan didalam struct
func (b *BangunRuang) aturPanjang(p int) {
	b.panjang = p
}

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

	color = append(color, "Biru", "Putih", "Hitam")

	fmt.Println(color)

	color = color[0:2] // [index awal: index akhir] jadi [0:2] 0 sampai 2, namun memotong slice artinya sebelum index akhir nya jadi hanya akan didapat 0,1 dan 2 nya tidak ikut
	fmt.Println(color)

	var name = [3]string{"Alice", "Bob", "Charles"}
	var namaSlice = name[:] // mengubah array menjadi slice, lalu jika nantinya isi slice diubah maka isi array juka akan terubah

	fmt.Println("kapasitas awal", cap(namaSlice))

	namaSlice[2] = "Alief"

	// jika melakukan perubahan nilai menggunakan append maka hanya isi slice yang terpengaruh sedangkan array tidak, hal ini dinamakan realokasi
	namaSlice = append(namaSlice, "Albert")

	fmt.Println("Slice", namaSlice)
	fmt.Println("array", name)
	fmt.Println("Kapasitas akhir", cap(namaSlice))
	// isi index ke 2 dari array dan slice akan sama sama berisi Alief karena ada perubahan pada "namaSlice" menggunakan indexing

	// Peningkatan panjang & kapasitas slice
	fmt.Println("\nPeningkatan panjang & Kapasitas slice")
	var pokemon = []string{}

	fmt.Println("Panjang", len(pokemon))
	fmt.Println("Kapasitas", cap(pokemon))

	fmt.Println("Tambah Pikachu")
	pokemon = append(pokemon, "Pikachu")

	fmt.Println("Panjang", len(pokemon))
	fmt.Println("Kapasitas", cap(pokemon))

	fmt.Println("Tambah Gengar")
	pokemon = append(pokemon, "Gengar")

	fmt.Println("Panjang", len(pokemon))
	fmt.Println("Kapasitas", cap(pokemon))

	fmt.Println("Tambah Greninja")
	pokemon = append(pokemon, "Greninja")

	fmt.Println("Panjang", len(pokemon))
	fmt.Println("Kapasitas", cap(pokemon))

	// kesimpulannya adalah kapasitas akan bertambah 2x lipat jika sudah melebihi batas kapasitasnya

	// tapi hal ini tidak edial dalam konteks efisiensi memori, maka dari itu lebih bijak jika kita menentukan kapasitas dari awal membuat slice itu sendiri

	// MAP
	fmt.Println("\nMap")
	var hobi map[string]string = map[string]string{"Alief": "Dota"} // cara pertama membuat map
	fmt.Println(hobi)

	var hobiMake map[string]string = make(map[string]string)
	hobiMake["Alip"] = "Kopi"
	hobiMake["Vina"] = "Tiktok"
	// zero value map itu nil maka dari itu lebih baik menggunakan make pada saat pembuatan map atau bisa juga langsung diinisiasikan agar tidak nil

	// value map itu case sensitive
	// if nilai hobi vina -> true maka turun ke yang pertama -> hobi vina blbablaa, namun jika hobi vina false / tidak ada di map maka dia akan turun ke else
	fmt.Println("Mengecek isi map menggunakan if")
	if nilai, ada := hobiMake["Vina"]; ada {
		fmt.Println("Hobi Vina adalah", nilai)
	} else {
		fmt.Println("Hobi Vina tidak tercatat dalam map")
	}

	// menggunakan range mirip seperti array bedanya kunci yang dipakai disini bukanlah index, tapi key value dari map itu sendiri
	fmt.Println("\nMengecek isi map menggunakan range")
	for kunci, nilai := range hobiMake {
		fmt.Printf("Hobi %s adalah %s\n", kunci, nilai)
	}

	fmt.Println("\nStruct")
	var persegi = BangunRuang{
		panjang: 5,
		lebar:   10,
		Geometri: Geometri{
			luas:     50,
			keliling: 0,
		},
	}
	fmt.Println(persegi.luas)

	// Jika isi struct di deklarasikan maka akan ada zero value dan akan tetap bisa berjalan
	persegiPanjang := BangunRuang{panjang: 5}
	fmt.Println(persegiPanjang)

	// var hasil = persegiPanjang.luas()
	// fmt.Println(hasil)
	fmt.Println(persegi.hitungLuas())

	fmt.Println("\nPointer Receiver Struct")
	persegi.aturPanjang(10)
	fmt.Println(persegi.panjang)
	// nilai panjang dari persegi akan terubah karena menggunkan pointer receiver, oleh karena itu jika kita butuh untuk mengubah nilai struct kita gunakan pointer receiver

	// persegiPointer := &BangunRuang{5, 10}
	// fmt.Println((*persegiPointer).panjang)

	// anonymous struct
	var persegiStruct = struct {
		panjang int
		lebar   int
	}{
		panjang: 10,
		lebar:   10,
	}

	fmt.Println(persegiStruct)

}
