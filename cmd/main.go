package main

import (
	"fmt"
	"github.com/akhilesharora/rolling-hash/pkg/rollinghash"
	"log"
	"os"
)

func main() {
	original, err := os.Open("tmp/original.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	updated, err := os.Open("tmp/updated.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer updated.Close()

	delta, err := rollinghash.ComputeDelta(original, updated)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(delta)
}
