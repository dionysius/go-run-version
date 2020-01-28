package main

import "fmt"

const (
	branch  = ""
	tag     = ""
	comment = "I'm the first commit, without tag or currently on a branch (at the commit time obviously branch master, but this branch will move once I continue)"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
