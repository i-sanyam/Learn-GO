package main

import (
	"fmt"
	"net/http"
)

var card string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		card = "Ace of Spade"
		card2 := "Ace of Diamond"
		card2 = "Five of Club"
		fmt.Println(card, card2)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Println(r.URL.Path)
	})

	http.ListenAndServe(":3000", nil)
}
