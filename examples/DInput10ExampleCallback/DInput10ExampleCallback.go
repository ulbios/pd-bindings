package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"bitbucket.org/pickdata-fw/emod_controller_binding_go/modules/DInput10Module"
)

func module_callback(data []byte, idFunction uint8, ctx interface{}) {
	buf := bytes.NewReader(data)

	// In this low level callback we have to find out if the callback is triggered by a input status change or an input counter.
	if idFunction == DInput10Module.IdFunctionINPUTS {
		// Decode data as big endian 2 bytes value

		var inputs_status uint16

		if err := binary.Read(buf, binary.BigEndian, &inputs_status); err != nil {
			fmt.Printf("binary.Read failed: %s\n", err)
		}

		// Print the digital inputs level
		for i := 0; i < int(DInput10Module.NUMBER_OF_DI_INPUTS); i++ {
			fmt.Printf("DI%02d status = %t\n", i+1, (inputs_status&(1<<i)) != 0)
		}
		fmt.Printf("\n")
	} else if idFunction >= DInput10Module.IdFunctionCOUNTER01 && idFunction <= DInput10Module.IdFunctionCOUNTER10 {
		// Decode data as big endian 4 bytes value

		var value uint32

		if err := binary.Read(buf, binary.BigEndian, &value); err != nil {
			fmt.Printf("binary.Read failed: %s\n", err)
		}

		// Get the specific input counter which triggered the callback.
		input := idFunction - DInput10Module.IdFunctionCOUNTER01 + 1

		// Print the counter value
		fmt.Printf("DI%02d counter = %d\n\n", input, value)
	} else {
		fmt.Printf("Unkown callback with idFunction %d and %d bytes data\n", idFunction, len(data))
	}
}

func main() {
	module := DInput10Module.NewDInput10Module()
	defer module.Destroy()

	var err error

	// Intialize the module with a callback
	if err = module.Init(module_callback, module); err != nil {
		fmt.Printf("Error initializing module: %s\n", err)
		return
	}

	// Reset the event configuration just in case, even if we are reconfiguring later
	if err := module.ResetEventConfig(); err != nil {
		fmt.Printf("Error resetting module event configuration: %s\n", err)
		return
	}

	// Configure the module to send a reading when any digital input level changes
	if err := module.ConfigEventOnNewData(); err != nil {
		fmt.Printf("Error setting module event at status change: %s\n", err)
		return
	}

	// Do nothing, the work is done by the callback
	for {
		time.Sleep(time.Second)
	}
}
