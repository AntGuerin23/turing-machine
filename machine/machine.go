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
	Operations        map[string][]Operation
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
		nextConfig, err := machine.runConfiguration(config)
		if err == nil {
			machine.currentConfig = nextConfig
		} else {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}

func (machine *Machine) runConfiguration(config Configuration) (byte, error) {
	//TODO : Refactor

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
	for _, op := range operations {
		if op.Function != nil {
			op.Function(machine)
		}
	}
	return config.NextConfiguration, nil
}

// Prints a snapshot of the machine
func (machine *Machine) printCompleteConfiguration() {
	configString := string(machine.currentConfig)
	if machine.currentConfig == 0 {
		configString = "Ending"
	}
	fmt.Printf("(config: %v) %v:%v \n", configString, machine.Head.Read(), machine.Head.Tape)
}
