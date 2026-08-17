package main

import (
	"fmt"
	"main/internals/bencode"
)

func main() {
	var line string = "2:Hi i22e 2:im 7:looking i69e 4:with 3:you"
	result := bencode.EVALUATE(line)
	for _, res := range result {
		fmt.Printf("%v\n", res)
	}
}
