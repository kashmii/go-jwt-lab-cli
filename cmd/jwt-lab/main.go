package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kashmii/go-jwt-lab-cli/internal/jwtlab"
)

func main() {
	mode := flag.String("mode", "inspect", "inspect|sign|verify")
	flag.Parse()
	switch *mode {
case "inspect":
    jwtlab.Inspect()
case "sign":
    // jwtlab.Sign() を実装予定
case "verify":
    // jwtlab.Verify() を実装予定
default:
    fmt.Fprintf(os.Stderr, "Unknown mode: %s\n", *mode)
    os.Exit(1)
}
}
