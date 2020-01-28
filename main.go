package main

import "fmt"

const (
	branch  = ""
	tag     = "v0.2.0"
	comment = "I'm on the second branch, with tag v0.2.0 but not holding the branch tag"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
