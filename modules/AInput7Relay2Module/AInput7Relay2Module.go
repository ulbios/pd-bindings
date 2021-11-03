package AInput7Relay2Module

/*
#cgo LDFLAGS: -lstdc++ -lrt -lemod_controller

#include <AInput7Relay2Module.h>

// This callback function is exported from go code, and calls the actual user registered callback
extern void aInput7Relay2Module_internal_go_cb(uint8_t* data, uint16_t data_len, uint8_t id_function, void* ctx);

// This callback function is registered for the C/C++ eMod API, and simply calls the previous function
extern void aInput7Relay2Module_internal_c_cb(const uint8_t *data, const uint16_t data_len, const uint8_t id_function, void *ctx);

// It is required to to it this way because you can't directly register an exported Go function in C/C++.
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

const (
	INPUT01   uint16 = uint16(C.AInput7Relay2Module_INPUT01)
	INPUT02   uint16 = uint16(C.AInput7Relay2Module_INPUT02)
	INPUT03   uint16 = uint16(C.AInput7Relay2Module_INPUT03)
	INPUT04   uint16 = uint16(C.AInput7Relay2Module_INPUT04)
	INPUT05   uint16 = uint16(C.AInput7Relay2Module_INPUT05)
	INPUT06   uint16 = uint16(C.AInput7Relay2Module_INPUT06)
	INPUT07   uint16 = uint16(C.AInput7Relay2Module_INPUT07)
	ALL_INPUT uint16 = uint16(C.AInput7Relay2Module_ALL_INPUT)
)

const (
	RELAY1    byte = byte(C.AInput7Relay2Module_RELAY1)
	RELAY2    byte = byte(C.AInput7Relay2Module_RELAY2)
	ALL_RELAY byte = byte(C.AInput7Relay2Module_ALL_RELAY)
)

const (
	NUMBER_OF_AI_INPUTS uint8 = uint8(C.AInput7Relay2Module_NUMBER_OF_AI_INPUTS)
	NUMBER_OF_RELAYS    uint8 = uint8(C.AInput7Relay2Module_NUMBER_OF_RELAYS)
)

const (
	IdFunctionINPUT01 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT01)
	IdFunctionINPUT02 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT02)
	IdFunctionINPUT03 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT03)
	IdFunctionINPUT04 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT04)
	IdFunctionINPUT05 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT05)
	IdFunctionINPUT06 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT06)
	IdFunctionINPUT07 uint8 = uint8(C.AInput7Relay2Module_idFunctionINPUT07)
)

const (
	MAX_ADC     uint16  = uint16(C.AInput7Relay2Module_MAX_ADC)
	MAX_VOLTAGE float32 = float32(C.AInput7Relay2Module_MAX_VOLTAGE)
	MAX_CURRENT float32 = float32(C.AInput7Relay2Module_MAX_CURRENT)
)

const (
	CFG_VOLTAGE byte = byte(C.AInput7Relay2Module_CFG_VOLTAGE)
	CFG_CURRENT byte = byte(C.AInput7Relay2Module_CFG_CURRENT)
)

type AInput7Relay2Module_callback_func func([]byte, uint8, interface{})

// This class represents the Analogue Inputs and Power Relays Module
// connected via eMOD bus to controller module. On the one hand, this module has
// 7 independent analogue inputs, the value of each one can be obtained
// separately. Its configuration for measuring voltage or current can be set
// also separately. And on the other hand, has 2 independent power relays
// than can be activated or deactivated separately. If polling is to be
// performed more than once per second, it's recommended set a callback instead.
type AInput7Relay2Module struct {
	ptr    *C.AInput7Relay2Module_t
	cb     AInput7Relay2Module_callback_func
	cb_ctx interface{}
}

//export aInput7Relay2Module_internal_go_cb
func aInput7Relay2Module_internal_go_cb(data *C.uint8_t, data_len C.uint16_t, id_function C.uint8_t, ctx unsafe.Pointer) {
	air := (*AInput7Relay2Module)(ctx)
	if air.cb != nil {
		air.cb(C.GoBytes(unsafe.Pointer(data), C.int(data_len)), uint8(id_function), air.cb_ctx)
	}
}

// Allocates internal memory and returns an instance of the module
func NewAInput7Relay2Module() *AInput7Relay2Module {
	m := new(AInput7Relay2Module)
	m.ptr = C.AInput7Relay2Module_create()

	return m
}

// Free the internal memory of the module. The instance can't be used anymore
// after that.
func (air *AInput7Relay2Module) Destroy() {
	C.AInput7Relay2Module_destroy(air.ptr)
}

// Initialize module. It is important to emphasize that this method is
// mandatory to be called and must be called always after instantiating the
// module. It can be called several times, each of which the module is
// initialized again.
// The callback function only will be called with the context if the module
// is later configured in callback mode.
func (air *AInput7Relay2Module) Init(cb AInput7Relay2Module_callback_func, cb_ctx interface{}) error {
	ret := C.AInput7Relay2Module_init(air.ptr, C.AInput7Relay2ModuleCallback_Type(C.aInput7Relay2Module_internal_c_cb), unsafe.Pointer(air))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	air.cb = cb
	air.cb_ctx = cb_ctx

	return nil
}

// Same as Init(), but allows to initialize a specific module number
// instead of the first one.
func (air *AInput7Relay2Module) InitModuleNumber(cb AInput7Relay2Module_callback_func, cb_ctx interface{}, module_number uint8) error {
	ret := C.AInput7Relay2Module_init_v(air.ptr, C.AInput7Relay2ModuleCallback_Type(C.aInput7Relay2Module_internal_c_cb), unsafe.Pointer(air), C.uint8_t(module_number))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	air.cb = cb
	air.cb_ctx = cb_ctx

	return nil
}

// Configure sample rate at which you want to check the analog input data.
func (air *AInput7Relay2Module) ConfigSampleRate(ms_period uint32) error {
	ret := C.AInput7Relay2Module_configSampleRate(air.ptr, C.uint32_t(ms_period))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configure an input or set of inputs to select them for acquiring current
// (4 to 20 mA) or voltage (0 to 10 V).
func (air *AInput7Relay2Module) ConfigInput(input_mask uint16) error {
	ret := C.AInput7Relay2Module_configInput(air.ptr, C.uint16_t(input_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Gets configured inputs mask. Each input indicates if it’s configured to
// acquire voltage (0) or current (1).
func (air *AInput7Relay2Module) GetInputConfig() (uint16, error) {
	var input_mask C.uint16_t
	ret := C.AInput7Relay2Module_getInputConfig(air.ptr, &input_mask)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return uint16(input_mask), nil
}

// Gets last measured value of a specified analog input.
func (air *AInput7Relay2Module) GetAnalogInput(input_mask uint16) (float32, error) {
	var data C.float
	ret := C.AInput7Relay2Module_getAnalogInput(air.ptr, C.uint16_t(input_mask), &data)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return float32(data), nil
}

// Gets last measured values for all analog inputs.
func (air *AInput7Relay2Module) GetAllAnalogInput() ([NUMBER_OF_AI_INPUTS]float32, error) {
	var data [NUMBER_OF_AI_INPUTS]float32
	var internal_data [NUMBER_OF_AI_INPUTS]C.float
	ret := C.AInput7Relay2Module_getAllAnalogInput(air.ptr, &internal_data[0])
	if ret != 0 {
		return data, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	for i := 0; i < int(NUMBER_OF_AI_INPUTS); i++ {
		data[i] = float32(internal_data[i])
	}

	return data, nil
}

// Configure pulse width of a relay o list of relays. When pulse width is 0
// a relay stays in the current state, and there is no pulse.
func (air *AInput7Relay2Module) ConfigPulseWidth(relay_mask byte, width_ms uint32) error {
	ret := C.AInput7Relay2Module_configPulseWidth(air.ptr, C.uint8_t(relay_mask), C.uint32_t(width_ms))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Activates a relay or set of relays.
func (air *AInput7Relay2Module) Activate(relay_mask byte) error {
	ret := C.AInput7Relay2Module_activate(air.ptr, C.uint8_t(relay_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Deactivates a relay or set of relays.
func (air *AInput7Relay2Module) Deactivate(relay_mask byte) error {
	ret := C.AInput7Relay2Module_deactivate(air.ptr, C.uint8_t(relay_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Activates all module relays.
func (air *AInput7Relay2Module) ActivateAll() error {
	ret := C.AInput7Relay2Module_activateAll(air.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Deactivates all module relays.
func (air *AInput7Relay2Module) DeactivateAll() error {
	ret := C.AInput7Relay2Module_deactivateAll(air.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Gets a relay status.
func (air *AInput7Relay2Module) GetRelayStatus(relay_mask byte) (byte, error) {
	var status C.uint8_t
	ret := C.AInput7Relay2Module_getRelayStatus(air.ptr, C.uint8_t(relay_mask), &status)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return byte(status), nil
}

// Gets all relay status.
func (air *AInput7Relay2Module) GetAllRelayStatus() (byte, error) {
	var status C.uint8_t
	ret := C.AInput7Relay2Module_getAllRelayStatus(air.ptr, &status)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return byte(status), nil
}

// Configures module for calling init callback_func at specified period.
func (air *AInput7Relay2Module) ConfigEventAtTimeInterval(time_interval time.Duration, event_mask uint16) error {
	ret := C.AInput7Relay2Module_configEventAtTimeInterval(air.ptr, C.uint16_t(time_interval.Milliseconds()), C.uint32_t(event_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configures module for calling Init callback func when input changes its
// value by a specified amount.
func (air *AInput7Relay2Module) ConfigEventOnValueChange(threshold uint32, event_mask uint16) error {
	ret := C.AInput7Relay2Module_configEventOnValueChange(air.ptr, C.uint32_t(threshold), C.uint32_t(event_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configures module for calling Init callback func when input value enters
// the specified range. The high and low limits of the range are in samples,
// so use the VoltageToSample or CurrentToSample methods.
func (air *AInput7Relay2Module) ConfigEventWithinRange(low_limit uint32, high_limit uint32, event_mask uint16) error {
	ret := C.AInput7Relay2Module_configEventWithinRange(air.ptr, C.uint32_t(low_limit), C.uint32_t(high_limit), C.uint32_t(event_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configures module for calling init callback_func when input value gets
// out of the specified range. The high and low limits of the range are in samples,
// so use the VoltageToSample or CurrentToSample methods.
func (air *AInput7Relay2Module) ConfigEventOutOfRange(low_limit uint32, high_limit uint32, event_mask uint16) error {
	ret := C.AInput7Relay2Module_configEventOutOfRange(air.ptr, C.uint32_t(low_limit), C.uint32_t(high_limit), C.uint32_t(event_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Resets all previously configured events.
func (air *AInput7Relay2Module) ResetEventConfiguration() error {
	ret := C.AInput7Relay2Module_resetEventConfiguration(air.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Converts voltage to discrete samples.
func (air *AInput7Relay2Module) VoltageToSamples(voltage float32) uint16 {
	return uint16(C.AInput7Relay2Module_voltageToSamples(air.ptr, C.float(voltage)))
}

// Converts current to discrete samples.
func (air *AInput7Relay2Module) CurrentToSamples(current float32) uint16 {
	return uint16(C.AInput7Relay2Module_currentToSamples(air.ptr, C.float(current)))
}

// Converts discrete samples to voltage.
func (air *AInput7Relay2Module) SamplesToVoltage(samples uint16) float32 {
	return float32(C.AInput7Relay2Module_samplesToVoltage(air.ptr, C.uint16_t(samples)))
}

// Converts discrete samples to current.
func (air *AInput7Relay2Module) SamplesToCurrent(samples uint16) float32 {
	return float32(C.AInput7Relay2Module_samplesToCurrent(air.ptr, C.uint16_t(samples)))
}
