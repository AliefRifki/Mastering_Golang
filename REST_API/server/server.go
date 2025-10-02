package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// 1. Buat route multiplexer
	mux := http.NewServeMux()

	// 3. Tambahkan handler ke mux
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// 4. Buat server
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 5. Jallankan server
	server.ListenAndServe()

}

var database = map[int]Products{}

// id untuk enumerasi di database
var lastID = 0

// 2. Fungsi Handler
func listProduct(w http.ResponseWriter, r *http.Request) {
	// Slice untuk response
	var products []Products

	// itersai pada map untuk menambahkan nilai produk dalam map ke slice response
	for _, v := range database {
		products = append(products, v)
	}

	// ubah menjadi json
	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	// beri response json tadi
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Terjadi Kesalahan Dalam Request"))
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Terjadi Kesalahan Dalam Request"))
	}

	// Increment nomor urut
	lastID++

	products.ID = lastID

	// Tambahkan ke db
	database[products.ID] = products

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("Request Berhasil Diproses dan ditambahkan"))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Terjadi Kesalahan Dalam Request"))
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("Terjadi Kesalahan Dalam Request"))
	}

	// Supaya ID terbaru tidak tergantikan dengan ID dari req body
	products.ID = productIDInt
	database[productIDInt] = products

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	// Hapus dari map
	delete(database, productIDInt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
