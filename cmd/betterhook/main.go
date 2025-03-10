package betterhook

import (
	"fmt"
	"log"

	Hook "github.com/yashGoyal40/BetterHook/pkg"
)

// SyncHook installs all Git hooks from the .betterhook directory
func SyncHook() {
	err := Hook.LoadAllHooks()
	if err != nil {
		log.Fatalf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ All hooks successfully synced 🎉")
	}
}

func SyncOneHook(hookName string) {
	err := Hook.LoadHook(hookName)
	if err != nil {
		log.Fatalf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Hook %q successfully synced 🎉\n", hookName)
	}
}
