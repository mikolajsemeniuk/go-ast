package main

import "fmt"

type One struct {
	Check string `json:"check,omitempty"`
	Hi    Hi     `json:"hi,omitempty"`
}

func Valid() bool {
	return true
}

type Hi struct {
	Hello int   `json:"hello"`
	First First `json:"first"`
}

type First string

//go:generate go run ./cmd/tool.go --type Pill
func main() {
	fmt.Println("s")
}
