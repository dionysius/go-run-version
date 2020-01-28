package main

import "fmt"

const (
	branch  = "third"
	tag     = ""
	comment = "ending third branch, one commit after v2.0.0"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
