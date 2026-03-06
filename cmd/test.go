package main

import (
	"fmt"
	"processBlueprint/internal/config"
)

func main() {
	cfg := config.Get()
	fmt.Println(cfg.Environment)

	// c := make(chan struct{})
	// <-c
}
