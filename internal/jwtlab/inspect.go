package jwtlab

import "fmt"

func Inspect(token string) {
	fmt.Println("Inspecting JWT:", token)
}