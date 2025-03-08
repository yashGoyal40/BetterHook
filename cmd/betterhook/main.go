package betterhook

import (
	"fmt"
	"log"

	Hook "github.com/yashGoyal40/BetterHook/pkg"
)

func SyncHook(message string) {
	err := Hook.LoadHook(message)
	if err != nil {
		log.Fatalf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Hook %q successfully synced 🎉\n", message)
	}
}
