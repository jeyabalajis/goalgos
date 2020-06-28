package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Println("Havey Kabisa!")

	eq := hash("p")+hash("a")+hash("t") == hash("a")+hash("t")+hash("p")
	if eq {
		fmt.Println("They're equal.")
	} else {
		fmt.Println("They're unequal.")
	}
}
