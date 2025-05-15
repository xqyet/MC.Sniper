package main

import (
	"fmt"
	"os"

	"github.com/Kqzz/MCsniperGO/pkg/mc"
)

// go run ./cmd/cli/login.go
// for token OAuth

// go run ./cmd/cli
// for manual bearer token only (temporary?)


func main() {
	// Create a Microsoft account type (used for bearer auth)
	acc := &mc.MCaccount{Type: mc.Ms}

	// Trigger the device code login flow
		err := acc.OauthFlow()
    	if err != nil {
    		panic(err)
    	}

    	if acc.Bearer == "" {
            fmt.Println("[!] Bearer token is empty — Minecraft authentication likely failed.")
            fmt.Println("[!] Check the output above for the Minecraft login error.")
            return
        }

	// Print the token to the console for verification
	fmt.Println("[*] Bearer token:")
	fmt.Println(acc.Bearer)

	// Save the bearer token to ms.txt
	fmt.Println("[*] Saving token to ms.txt...")
	f, err := os.OpenFile("ms.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(acc.Bearer + "\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("[*] Bearer token saved to ms.txt ✅")
}
