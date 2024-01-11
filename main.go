package main

import (
	machinePkg "turing-machine/machine"
)

func main() {
	//TODO : Replace with a file read of the machine's configs
	configs := make(map[byte]machinePkg.Configuration, 10)
	op := append(make([]machinePkg.Operation, 1), machinePkg.Operation{Function: b})
	operations := make(map[string][]machinePkg.Operation, 1)
	operations[""] = op
	config := machinePkg.Configuration{Operations: operations, NextConfiguration: 0}
	configs['b'] = config
	machine := machinePkg.Machine{Configs: configs}
	machine.Start()
}

// TODO : Remove this
func b(machine *machinePkg.Machine) {
	machine.Head.MoveRight()
	machine.Head.Write("5")
	machine.Head.MoveRight()
	machine.Head.MoveRight()
	machine.Head.Write("3")
	machine.Head.MoveRight()
	machine.Head.Write("1")
	machine.Head.MoveLeft()
	machine.Head.MoveRight()
}
