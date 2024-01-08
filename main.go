package main

import (
	machinePkg "turing-machine/machine"
)

func main() {
	configs := make(map[byte]machinePkg.Configuration, 10)
	op := machinePkg.Operation{Function: b}
	operations := make([]machinePkg.Operation, 1)
	operations = append(operations, op)
	config := machinePkg.Configuration{Operations: operations, NextConfiguration: 0}
	configs['b'] = config
	machine := machinePkg.Machine{Configs: configs}
	machine.Start()
}

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
