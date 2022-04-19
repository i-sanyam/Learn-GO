package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error in API Call", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println("\nResponse from API Call\n", resp)
	fmt.Println("\nResponse Body\n", string(bs))
}
