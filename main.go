package main

import (
	"fmt"
	"turing-machine/head"
)

func main() {
	machine := head.Head{Tape: make([]string, 50)}
	machine.Write("salut")
	fmt.Println(machine.Read())
}
