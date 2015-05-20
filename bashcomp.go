package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if strings.HasPrefix(pair[0], "COMP_") {
			fmt.Fprintf(os.Stderr, "%s=%q\n", pair[0], pair[1])
		}
	}
}

// Example:
// COMP_TYPE="9"
// COMP_LINE="abc def"
// COMP_POINT="7"
// COMP_KEY="9"
