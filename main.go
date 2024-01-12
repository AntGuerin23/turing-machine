package main

import (
	machinePkg "turing-machine/machine"
)

func main() {
	machineTest()
	//fileTest()
}

// TODO : Remove this
func b(machine *machinePkg.Machine) {
	machine.Head.MoveRight()
	machine.Head.Write('5')
	machine.Head.MoveRight()
	machine.Head.MoveRight()
	machine.Head.Write('3')
	machine.Head.MoveRight()
	machine.Head.Write('1')
	machine.Head.MoveLeft()
	machine.Head.MoveRight()
}
