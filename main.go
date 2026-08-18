package main

import (
	"fmt"
	"main/internals/bencode"
)

func main() {
	var line string = "l3:bar4:spam3:fooi42ee"
	result := bencode.EVALUATE(line)
	for _, res := range result {
		fmt.Printf("%v\n", res)
	}
}
