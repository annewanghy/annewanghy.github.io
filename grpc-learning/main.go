package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go docker tutorial")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "helloo world")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
