package main

import (
	Hook "github.com/yashGoyal40/BetterHook/pkg"
	"log"
)

func betterhook(message string) {
	err := Hook.LoadHook(message)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
