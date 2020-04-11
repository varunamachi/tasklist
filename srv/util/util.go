package util

import (
	"encoding/json"
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

//DumpJSON - dumps JSON representation of given data to stdout
func DumpJSON(o interface{}) {
	b, err := json.MarshalIndent(o, "", "    ")
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println("Failed to marshal data to JSON", err)
	}
}

// AskSecret - asks password from user, does not echo charectors
func AskSecret(prompt string) (secret string, err error) {
	fmt.Print(prompt, ": ")

	var pbyte []byte
	pbyte, err = terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		secret = string(pbyte)
		fmt.Println()
	}
	return secret, err
}
