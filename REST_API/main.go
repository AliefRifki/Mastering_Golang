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
	Tipe  string `json:"type"`
	Setup string `json:"setup"`
	Jokes string `json:"delivery"`
}

func main() {
	// HTTP request mengembalikan request dan error
	// 1. Buat Request
	req, err := http.NewRequest("GET", "https://v2.jokeapi.dev/joke/Programming?type=twopart", nil)
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

	fmt.Println("Tipe:", factResponse.Tipe)
	fmt.Println("Setup:", factResponse.Setup)
	fmt.Println("Punchline:", factResponse.Jokes)

}
