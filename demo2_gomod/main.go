package main

import (
	"fmt"

	layer1 "testMod/pkg/layer1"

	FavoriteNumber "github.com/bencodehard/go_sample_packages"
)

func main() {
	// Package from Github
	FavoriteNumber.FavoriteNumber()

	// Package from local
	layer1.Layer1()

	// End
	fmt.Println("Done Process")
}
