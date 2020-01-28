package main

import "fmt"

const (
	branch  = ""
	tag     = ""
	comment = "A yeah, forgot, if you go get with @master, you should get me, I should be the strange go preversion v0.1.1-0"
)

func main() {
	fmt.Printf("branch: %s\n", branch)
	fmt.Printf("tag: %s\n", tag)
	fmt.Printf("comment: %s\n", comment)
}
