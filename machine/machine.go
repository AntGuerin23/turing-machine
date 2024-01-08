package machine

import (
	"fmt"
	"time"
)

// TODO : Find a way to move operation and configuration
type Operation struct {
	Function func(*Machine)
}

type Configuration struct {
	Operations        []Operation
	NextConfiguration byte
}

type Machine struct {
	Configs       map[byte]Configuration
	Head          Head //TODO : Make private
	currentConfig byte
}

func (machine *Machine) Start() {
	machine.initialize()
	machine.run()
}

func (machine *Machine) initialize() {
	machine.currentConfig = 'b' //Begin
	machine.Head = Head{Tape: make([]string, 1)}
}

func (machine *Machine) run() {
	for {
		machine.printCompleteConfiguration()
		if machine.currentConfig == 0 {
			break
		}
		config := machine.Configs[machine.currentConfig]
		machine.currentConfig = machine.runConfiguration(config)
		time.Sleep(time.Second)
	}
}

func (machine *Machine) runConfiguration(config Configuration) byte {
	//TODO : Run different operations based on machine.Read (symbol, Any, None)
	// Make operations a map of ([Any, None, ...], []Operation)
	for _, op := range config.Operations {
		if op.Function != nil {
			op.Function(machine)
		}
	}
	return config.NextConfiguration
}

// Prints a snapshot of the machine
func (machine *Machine) printCompleteConfiguration() {
	configString := string(machine.currentConfig)
	if machine.currentConfig == 0 {
		configString = "Ending"
	}
	fmt.Printf("(config: %v) %v:%v \n", configString, machine.Head.Read(), machine.Head.Tape)
}
