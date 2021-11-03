package main

import (
	"fmt"
	"time"

	"bitbucket.org/pickdata-fw/emod_controller_binding_go/modules/AInput7Relay2Module"
)

func displayInputs(data [AInput7Relay2Module.NUMBER_OF_AI_INPUTS]float32) {
	fmt.Printf("AI01 = %2.2f V //", data[0])
	fmt.Printf("AI02 = %2.2f V //", data[1])
	fmt.Printf("AI03 = %2.2f V //", data[2])
	fmt.Printf("AI04 = %2.2f V //", data[3])
	fmt.Printf("AI05 = %2.2f mA //", data[4])
	fmt.Printf("AI06 = %2.2f mA //", data[5])
	fmt.Printf("AI07 = %2.2f mA", data[6])
	fmt.Printf("\n")
}

func testAll(module *AInput7Relay2Module.AInput7Relay2Module) {
	type config_def struct {
		input_mask uint16
		config     bool
	}

	var err error

	// Analog input configuration as voltage or current
	input_config := [AInput7Relay2Module.NUMBER_OF_AI_INPUTS]config_def{
		{AInput7Relay2Module.INPUT01, false}, //voltage
		{AInput7Relay2Module.INPUT02, false}, //voltage
		{AInput7Relay2Module.INPUT03, false}, //voltage
		{AInput7Relay2Module.INPUT04, false}, //voltage
		{AInput7Relay2Module.INPUT05, true},  //current
		{AInput7Relay2Module.INPUT06, true},  //current
		{AInput7Relay2Module.INPUT07, true},  //current
	}

	// Convert previous configuration to mask
	var cmask uint16
	for i := 0; i < int(AInput7Relay2Module.NUMBER_OF_AI_INPUTS); i++ {
		if input_config[i].config {
			cmask |= (1 << i)
		}
	}

	// Apply input configuration mask
	if err = module.ConfigInput(cmask); err != nil {
		fmt.Printf("Error setting analog inputs configuration: %s\n", err)
		return
	}

	time.Sleep(100 * time.Millisecond)

	// Check if the current input configuration mask is the expected one
	if crmask, err := module.GetInputConfig(); err != nil {
		fmt.Printf("Error getting alalog inputs configuration: %s\n", err)
		return
	} else if crmask != cmask {
		fmt.Printf("Bad input configuration\n")
		return
	}

	time.Sleep(20 * time.Millisecond)

	// Get all analog inputs value array
	aInputs, err := module.GetAllAnalogInput()
	if err != nil {
		fmt.Printf("Error getting all analog input values: %s\n", err)
		return
	}

	// Print the array
	displayInputs(aInputs)

	// Activate relay 1
	if err = module.Activate(AInput7Relay2Module.RELAY1); err != nil {
		fmt.Printf("Error activating relay 1: %s\n", err)
		return
	}

	// Activate relay 2
	if err = module.Activate(AInput7Relay2Module.RELAY2); err != nil {
		fmt.Printf("Error activating relay 2: %s\n", err)
		return
	}

	// Wait for 100 ms
	time.Sleep(100 * time.Millisecond)

	// Configure sample rate of the analog inputs to 200 ms
	if err = module.ConfigSampleRate(200); err != nil {
		fmt.Printf("Error configuring sample rate: %s\n", err)
		return
	}

	for {
		// Read each analog input individually instead of all at the same time

		aInputs[0], err = module.GetAnalogInput(AInput7Relay2Module.INPUT01)
		if err != nil {
			fmt.Printf("Error getting analog input 1 value: %s\n", err)
			return
		}

		aInputs[1], err = module.GetAnalogInput(AInput7Relay2Module.INPUT02)
		if err != nil {
			fmt.Printf("Error getting analog input 2 value: %s\n", err)
			return
		}

		aInputs[2], err = module.GetAnalogInput(AInput7Relay2Module.INPUT03)
		if err != nil {
			fmt.Printf("Error getting analog input 3 value: %s\n", err)
			return
		}

		aInputs[3], err = module.GetAnalogInput(AInput7Relay2Module.INPUT04)
		if err != nil {
			fmt.Printf("Error getting analog input 4 value: %s\n", err)
			return
		}

		aInputs[4], err = module.GetAnalogInput(AInput7Relay2Module.INPUT05)
		if err != nil {
			fmt.Printf("Error getting analog input 5 value: %s\n", err)
			return
		}

		aInputs[5], err = module.GetAnalogInput(AInput7Relay2Module.INPUT06)
		if err != nil {
			fmt.Printf("Error getting analog input 6 value: %s\n", err)
			return
		}

		aInputs[6], err = module.GetAnalogInput(AInput7Relay2Module.INPUT07)
		if err != nil {
			fmt.Printf("Error getting analog input 7 value: %s\n", err)
			return
		}

		// Print the array
		displayInputs(aInputs)

		time.Sleep(time.Second)
	}
}

func main() {
	module := AInput7Relay2Module.NewAInput7Relay2Module()
	defer module.Destroy()

	// Intialize the module without any callback
	if err := module.Init(nil, nil); err != nil {
		fmt.Printf("Error initializing module: %s\n", err)
		return
	}

	// Reset the event configuration just in case, as we want to poll the inputs
	if err := module.ResetEventConfiguration(); err != nil {
		fmt.Printf("Error resetting module event configuration: %s\n", err)
		return
	}

	testAll(module)
}
