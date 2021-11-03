package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"bitbucket.org/pickdata-fw/emod_controller_binding_go/modules/AInput7Relay2Module"
)

func module_callback(data []byte, idFunction uint8, ctx interface{}) {
	// In this low level callback we have to find out which input triggered it with idFunction parameter.
	if idFunction >= AInput7Relay2Module.IdFunctionINPUT01 && idFunction <= AInput7Relay2Module.IdFunctionINPUT07 {
		// Decode data as big endian 2 bytes value
		buf := bytes.NewReader(data)
		var samples uint16

		if err := binary.Read(buf, binary.BigEndian, &samples); err != nil {
			fmt.Printf("binary.Read failed: %s\n", err)
		}

		// Get the specific input which triggered the callback.
		input := idFunction - AInput7Relay2Module.IdFunctionINPUT01 + 1

		// Get module reference through ctx parameter type assertion
		if module, ok := ctx.(*AInput7Relay2Module.AInput7Relay2Module); ok {
			// Get actual voltage value from discrete samples reading
			voltage := module.SamplesToVoltage(samples)

			fmt.Printf("AI%02d = %2.2f V\n", input, voltage)
		} else {
			fmt.Printf("Failed type assertion\n")
		}
	} else {
		fmt.Printf("Unkown callback with idFunction %d and %d bytes data\n", idFunction, len(data))
	}
}

func main() {
	module := AInput7Relay2Module.NewAInput7Relay2Module()
	defer module.Destroy()

	// Intialize the module with a callback
	if err := module.Init(module_callback, module); err != nil {
		fmt.Printf("Error initializing module: %s\n", err)
		return
	}

	// Reset the event configuration just in case, even if we are reconfiguring later
	if err := module.ResetEventConfiguration(); err != nil {
		fmt.Printf("Error resetting module event configuration: %s\n", err)
		return
	}

	// Configure the module to send a reading every second for each analog input
	if err := module.ConfigEventAtTimeInterval(time.Second*1, AInput7Relay2Module.ALL_INPUT); err != nil {
		fmt.Printf("Error setting module event at time interval: %s\n", err)
		return
	}

	// Manually cycle the relays
	for {
		if err := module.Activate(AInput7Relay2Module.ALL_RELAY); err != nil {
			fmt.Printf("Error activating relays: %s\n", err)
			return
		}

		time.Sleep(time.Millisecond * 500)

		if err := module.Deactivate(AInput7Relay2Module.ALL_RELAY); err != nil {
			fmt.Printf("Error deactivating relays: %s\n", err)
			return
		}

		time.Sleep(time.Millisecond * 500)
	}
}
