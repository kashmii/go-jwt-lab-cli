package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kashmii/go-jwt-lab-cli/internal/jwtlab"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: jwt-lab <command>\n")
		fmt.Fprintf(os.Stderr, "Commands: inspect, sign, verify\n")
		os.Exit(1)
	}

	mode := args[0]
	switch mode {
	case "inspect":
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "Usage: jwt-lab inspect <token>\n")
			os.Exit(1)
		}
		token := args[1]
		jwtlab.Inspect(token)
	case "sign":
		// jwtlab.Sign() を実装予定
	case "verify":
		// jwtlab.Verify() を実装予定
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", mode)
		os.Exit(1)
	}
}
