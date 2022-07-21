package main

import "fmt"

type One struct {
	Check string `json:"check,omitempty"`
	Hi    Hi     `json:"hi,omitempty"`
}

type Hi struct {
	Hello int   `json:"hello"`
	First First `json:"first"`
}

func (h *Hi) Valid() bool {
	return true
}

type First string

func (f *First) Valid() bool {
	return true
}

//go:generate go run ./cmd/tool.go --type Pill
func main() {
	fmt.Println("s")
}
