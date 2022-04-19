package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.com");
	if (err != nil) {
		fmt.Println("Error in API Call", err);
		os.Exit(1);
	}
	fmt.Println("Response from API Call\n", resp);
}