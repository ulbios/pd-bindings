package main

import (
	"fmt"
	"time"

	"bitbucket.org/pickdata-fw/emod_controller_binding_go/modules/Relay8Module"
)

func main() {
	relay_module := Relay8Module.NewRelay8Module()
	defer relay_module.Destroy()

	var err error

	// Intialize the module
	if err = relay_module.Init(); err != nil {
		fmt.Printf("Error initializing module: %s\n", err)
		return
	}

	// Disable the pulsing for the first to relays so they stay with their configured values later
	if err = relay_module.ConfigPulseWidth(Relay8Module.RELAY1|Relay8Module.RELAY2, 0); err != nil {
		fmt.Printf("Error configuring pulse width: %s\n", err)
		return
	}

	// Alternatively toggle each relay individually
	for {
		if err = relay_module.Activate(Relay8Module.RELAY1); err != nil {
			fmt.Printf("Error activating relay 1: %s\n", err)
		}
		if err = relay_module.Deactivate(Relay8Module.RELAY2); err != nil {
			fmt.Printf("Error deactivating relay 2: %s\n", err)
		}
		time.Sleep(time.Second)

		if err = relay_module.Activate(Relay8Module.RELAY2); err != nil {
			fmt.Printf("Error activating relay 2: %s\n", err)
		}
		if err = relay_module.Deactivate(Relay8Module.RELAY1); err != nil {
			fmt.Printf("Error deactivating relay 1: %s\n", err)
		}
		time.Sleep(time.Second)
	}
}
