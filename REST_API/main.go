package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// struct dengan json struct tag
type FactResponse struct {
	Teks string `json:"text"`
	Tipe string `json:"type"`
}

func main() {
	// HTTP request mengembalikan request dan error
	// 1. Buat Request
	req, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts/random", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 2. Buat Client
	client := http.Client{}
	// 3. Panggil request dengan client
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// tutup response body
	defer res.Body.Close()

	// 4. Baru baca response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 5. Convert ke tipe data custom (FactResponse)
	var factResponse FactResponse
	err = json.Unmarshal(resBody, &factResponse)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Text", factResponse.Teks)
	fmt.Println("Tipe", factResponse.Tipe)

}
