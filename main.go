package main

import "fmt"

const (
	branch  = ""
	tag     = ""
	comment = "since I'm now on v2, I shouldn't be used"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
