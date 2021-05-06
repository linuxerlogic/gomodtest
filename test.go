package gomodtest

import (
	"fmt"
)

func SayHi(a string) {
	fmt.Printf("Hi %s, the is gomod test from github.com/linuxerlogic/gomodtest", a)
}

func Add(m, n int) int {
	return m + n
}
