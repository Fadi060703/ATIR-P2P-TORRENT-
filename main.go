package main

import (
	"fmt"
	"main/internals/bencode"
)

func main() {
	var line string = "i3e 4:test l3:wowi17ee"
	result := bencode.EVALUATE(line)
	for _, res := range result {
		fmt.Printf("%v\n", res)
	}
}
