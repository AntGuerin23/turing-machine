package machine

import (
	"fmt"
	"time"
)

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
		machine.printState()
		if machine.currentConfig == 0 {
			break
		}
		config := machine.Configs[machine.currentConfig]
		nextConfig := machine.runConfiguration(config)
		machine.currentConfig = nextConfig
		time.Sleep(time.Second) //Here for debugging
	}
}

func (machine *Machine) runConfiguration(config Configuration) byte {
	operations := machine.getOperationsBranchFromCurrentSymbol(config) // A config can have multiple arrays of operations, they're chosen using the current symbol, hence the branching
	machine.runOperations(operations)
	return config.NextConfiguration
}

func (machine *Machine) getOperationsBranchFromCurrentSymbol(config Configuration) []Operation {
	//A normal Turing machine configuration can check for
	//A particular written symbol (1, 0, etc)
	//"Any" (Any symbol),
	//"None" (Nothing written),
	//Nothing or "" (no branching)

	var operations []Operation
	symbol := machine.Head.Read()
	hasSymbol := symbol != ""

	if hasSymbol {
		operations = config.Operations[symbol]
		if operations == nil {
			operations = config.Operations["Any"]
		}
	} else {
		operations = config.Operations["None"]
	}
	if operations == nil {
		operations = config.Operations[""]
	}
	return operations
}

func (machine *Machine) runOperations(operations []Operation) {
	for _, op := range operations {
		if op.Function != nil {
			op.Function(machine)
		}
	}
}

// Prints a snapshot of the machine
func (machine *Machine) printState() {
	configString := string(machine.currentConfig)
	if machine.currentConfig == 0 {
		configString = "End"
	}
	fmt.Printf("(config: %v) %v:%v \n", configString, machine.Head.Read(), machine.Head.Tape)
}
