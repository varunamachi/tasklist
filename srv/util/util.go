package util

import (
	"encoding/json"
	"fmt"
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
