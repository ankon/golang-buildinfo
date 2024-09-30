package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	if bi, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("%+v\n", bi)
	}
}
