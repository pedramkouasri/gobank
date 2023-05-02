package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	NewAPIServer(":3000", store).Run()

	fmt.Println("Yeh Build")
}
