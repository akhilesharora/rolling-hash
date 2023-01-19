package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akhilesharora/rolling-hash/pkg/rollinghash"
)

func main() {
	original, err := os.Open("testData/original.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	updated, err := os.Open("testData/updated.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer updated.Close()

	r := rollinghash.NewRollingHash(1024)

	delta, err := r.ComputeDelta(original, updated)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(delta))
}
