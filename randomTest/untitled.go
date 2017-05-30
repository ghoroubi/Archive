package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	uuid := make([]byte, 16)
	num, err := io.ReadFull(rand.Reader, uuid)
	n := fmt.Sprintf("%x-%x-%x-%x", uuid[0:3], uuid[4:6], uuid[7:10], uuid[11:15])
	fmt.Print(n, "\n")
	if num != len(uuid) || err != nil {
		panic(err)
	}
}
