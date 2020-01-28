package main

import "fmt"

const (
	branch  = "second"
	tag     = "v0.2.1"
	comment = "both branch tag and tag v0.2.1"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
