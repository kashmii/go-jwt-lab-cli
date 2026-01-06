package jwtlab

import (
	"fmt"
	"strings"
)

func Inspect(token string) error {
	fmt.Println("Inspecting JWT:", token)
	// token を3つに分ける
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT token format")
		return fmt.Errorf("invalid JWT: expected 3 parts, got %d", len(parts))
	}

	return nil
}