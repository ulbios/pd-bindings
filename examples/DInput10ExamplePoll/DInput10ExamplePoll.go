package main

import (
	"fmt"
	"time"

	"bitbucket.org/pickdata-fw/emod_controller_binding_go/modules/DInput10Module"
)

func main() {
	module := DInput10Module.NewDInput10Module()
	defer module.Destroy()

	// Intialize the module without any callback
	if err := module.Init(nil, nil); err != nil {
		fmt.Printf("Error initializing module: %s\n", err)
		return
	}

	// Reset the event configuration just in case, as we want to poll the inputs
	if err := module.ResetEventConfig(); err != nil {
		fmt.Printf("Error resetting module event configuration: %s\n", err)
		return
	}

	for {
		// Poll the inputs status every second and print it.
		if status, err := module.GetAllStatus(); err == nil {
			for i := 0; i < int(DInput10Module.NUMBER_OF_DI_INPUTS); i++ {
				fmt.Printf("DI%02d status = %t\n", i+1, status[i])
			}
			fmt.Printf("\n")

			time.Sleep(time.Second)
		} else {
			fmt.Printf("Error getting status: %s\n", err)
			return
		}
	}
}
