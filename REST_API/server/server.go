package main

import (
	"encoding/json"
	"net/http"
)

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	price int    `json:"price"`
}

func main() {
	// 1. Buat route multiplexer
	mux := http.NewServeMux()

	// 3. Tambahkan handler ke mux
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", listProduct)
	mux.HandleFunc("PUT /products{id}", updateProduct)
	mux.HandleFunc("DELETE /products{id}", deleteProduct)

	// 4. Buat server
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 5. Jallankan server
	server.ListenAndServe()

}

var database = map[int]Products{}

// 2. Fungsi Handler
func listProduct(w http.ResponseWriter, r *http.Request) {
	var products []Products

	for _, v := range database {
		products = append(products, v)
	}

	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("Terjadi Kesalahan"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

}
