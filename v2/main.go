package main

import "fmt"

const (
	branch  = ""
	tag     = "v2.0.0"
	comment = "opening third branch"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
