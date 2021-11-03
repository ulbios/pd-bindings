package DInput10Module

/*
#cgo LDFLAGS: -lstdc++ -lrt -lemod_controller

#include <DInput10Module.h>

// This callback function is exported from go code, and calls the actual user registered callback
extern void dInput10Module_internal_go_cb(uint8_t* data, uint16_t data_len, uint8_t id_function, void* ctx);

// This callback function is registered for the C/C++ eMod API, and simply calls the previous function
extern void dInput10Module_internal_c_cb(const uint8_t *data, const uint16_t data_len, const uint8_t id_function, void *ctx);

// It is required to to it this way because you can't directly register an exported Go function in C/C++.
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

const (
	INPUT01    uint16 = uint16(C.DInput10Module_INPUT01)
	INPUT02    uint16 = uint16(C.DInput10Module_INPUT02)
	INPUT03    uint16 = uint16(C.DInput10Module_INPUT03)
	INPUT04    uint16 = uint16(C.DInput10Module_INPUT04)
	INPUT05    uint16 = uint16(C.DInput10Module_INPUT05)
	INPUT06    uint16 = uint16(C.DInput10Module_INPUT06)
	INPUT07    uint16 = uint16(C.DInput10Module_INPUT07)
	INPUT08    uint16 = uint16(C.DInput10Module_INPUT08)
	INPUT09    uint16 = uint16(C.DInput10Module_INPUT09)
	INPUT10    uint16 = uint16(C.DInput10Module_INPUT10)
	ALL_INPUTS uint16 = uint16(C.DInput10Module_ALL_INPUT)

	NUMBER_OF_DI_INPUTS  uint8  = uint8(C.DInput10Module_NUMBER_OF_DI_INPUTS)
	MODE_PRETRIGGER_TIME uint32 = uint32(C.DInput10Module_MODE_PRETRIGGER_TIME)
	MODE_PULSE_COUNTER   uint32 = uint32(C.DInput10Module_MODE_PULSE_COUNTER)
	MODE_WIDTH_COUNTER   uint32 = uint32(C.DInput10Module_MODE_WIDTH_COUNTER)

	IdFunctionINPUTS    uint8 = uint8(C.DInput10Module_idFunctionINPUTS)
	IdFunctionCOUNTER01 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER01)
	IdFunctionCOUNTER02 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER02)
	IdFunctionCOUNTER03 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER03)
	IdFunctionCOUNTER04 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER04)
	IdFunctionCOUNTER05 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER05)
	IdFunctionCOUNTER06 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER06)
	IdFunctionCOUNTER07 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER07)
	IdFunctionCOUNTER08 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER08)
	IdFunctionCOUNTER09 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER09)
	IdFunctionCOUNTER10 uint8 = uint8(C.DInput10Module_idFunctionCOUNTER10)
)

type DInput10Module_callback_func func([]byte, uint8, interface{})

// This struct represents the 10 Digital Inputs Module connected via eMOD bus
// to controller module. This module has 10 independent digital inputs and
// the status of each one can be obtained separately. If a polling is to be
// performed more than once per second, it's recommended to set a callback instead.
type DInput10Module struct {
	ptr    *C.DInput10Module_t
	cb     DInput10Module_callback_func
	cb_ctx interface{}
}

//export dInput10Module_internal_go_cb
func dInput10Module_internal_go_cb(data *C.uint8_t, data_len C.uint16_t, id_function C.uint8_t, ctx unsafe.Pointer) {
	air := (*DInput10Module)(ctx)
	if air.cb != nil {
		air.cb(C.GoBytes(unsafe.Pointer(data), C.int(data_len)), uint8(id_function), air.cb_ctx)
	}
}

// Allocates internal memory and returns an instance of the module
func NewDInput10Module() *DInput10Module {
	m := new(DInput10Module)
	m.ptr = C.DInput10Module_create()

	return m
}

// Free the internal memory of the module. The instance can't be used anymore
// after that.
func (din *DInput10Module) Destroy() {
	C.DInput10Module_destroy(din.ptr)
}

// Initialize module. It is important to emphasize that this method is
// mandatory to be called and must be called always after instantiating the
// module. It can be called several times, each of which the module is
// initialized again.
// The callback function only will be called with the context if the module
// is later configured in callback mode.
func (din *DInput10Module) Init(cb DInput10Module_callback_func, cb_ctx interface{}) error {
	ret := C.DInput10Module_init(din.ptr, C.DInput10ModuleCallback_Type(C.dInput10Module_internal_c_cb), unsafe.Pointer(din))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	din.cb = cb
	din.cb_ctx = cb_ctx

	return nil
}

// Same as Init(), but allows to initialize a specific module number
// instead of the first one.
func (din *DInput10Module) InitModuleNumber(cb DInput10Module_callback_func, cb_ctx interface{}, module_number uint8) error {
	ret := C.DInput10Module_init_v(din.ptr, C.DInput10ModuleCallback_Type(C.dInput10Module_internal_c_cb), unsafe.Pointer(din), C.uint8_t(module_number))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	din.cb = cb
	din.cb_ctx = cb_ctx

	return nil
}

// Gets a digital input status.
func (din *DInput10Module) GetStatus(input_mask uint16) (uint8, error) {
	var status C.uint8_t
	ret := C.DInput10Module_getStatus(din.ptr, C.uint16_t(input_mask), &status)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return uint8(status), nil
}

// Gets all digital input status.
func (din *DInput10Module) GetAllStatus() ([NUMBER_OF_DI_INPUTS]bool, error) {
	var status [NUMBER_OF_DI_INPUTS]bool
	var internal_status [NUMBER_OF_DI_INPUTS]C.uint8_t
	ret := C.DInput10Module_getAllStatus(din.ptr, &internal_status[0])
	if ret != 0 {
		return status, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	for i := 0; i < int(NUMBER_OF_DI_INPUTS); i++ {
		status[i] = internal_status[i] != 0
	}

	return status, nil
}

// Gets number of pulses occurred in an input after a reset.
func (din *DInput10Module) GetPulseCount(input_number uint16) (uint32, error) {
	var count C.uint32_t
	ret := C.DInput10Module_getPulseCount(din.ptr, C.uint16_t(input_number), &count)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return uint32(count), nil
}

// Gets number of pulses occurred in all inputs.
func (din *DInput10Module) GetAllPulseCount() ([NUMBER_OF_DI_INPUTS]uint32, error) {
	var counts [NUMBER_OF_DI_INPUTS]uint32
	var internal_counts [NUMBER_OF_DI_INPUTS]C.uint32_t
	ret := C.DInput10Module_getAllPulseCount(din.ptr, &internal_counts[0])
	if ret != 0 {
		return counts, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	for i := 0; i < int(NUMBER_OF_DI_INPUTS); i++ {
		counts[i] = uint32(internal_counts[i])
	}

	return counts, nil
}

// Resets the number of pulses in an input.
func (din *DInput10Module) ResetPulseCount(input_mask uint16) error {
	ret := C.DInput10Module_resetPulseCount(din.ptr, C.uint16_t(input_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Resets the number of pulses in all inputs.
func (din *DInput10Module) ResetAllPulseCount() error {
	ret := C.DInput10Module_resetAllPulseCount(din.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Gets last pulse width occurred in an input.
func (din *DInput10Module) GetPulseWidth(input_mask uint16) (uint32, error) {
	var width C.uint32_t
	ret := C.DInput10Module_getPulseWidth(din.ptr, C.uint16_t(input_mask), &width)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return uint32(width), nil
}

// Gets last pulse width occurred in all inputs.
func (din *DInput10Module) GetAllPulseWidth() ([NUMBER_OF_DI_INPUTS]uint32, error) {
	var widths [NUMBER_OF_DI_INPUTS]uint32
	var internal_widths [NUMBER_OF_DI_INPUTS]C.uint32_t
	ret := C.DInput10Module_getAllPulseWidth(din.ptr, &internal_widths[0])
	if ret != 0 {
		return widths, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	for i := 0; i < int(NUMBER_OF_DI_INPUTS); i++ {
		widths[i] = uint32(internal_widths[i])
	}

	return widths, nil
}

// Configures module for calling Init callback func at specified period.
func (din *DInput10Module) ConfigEventAtTimeInterval(time_interval time.Duration) error {
	ret := C.DInput10Module_configEventAtTimeInterval(din.ptr, C.uint16_t(time_interval.Milliseconds()))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Usually, this function only needs to be called at the beginning in the
// configuration stage. It allows to change any input mode from status to
// pulse counter.
func (din *DInput10Module) SwitchToMode(event_mask uint16, mode uint32) error {
	ret := C.DInput10Module_switchToMode(din.ptr, C.uint16_t(event_mask), C.uint32_t(mode))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configures module for calling Init callback func when a input value changes.
func (din *DInput10Module) ConfigEventOnNewData() error {
	ret := C.DInput10Module_configEventOnNewData(din.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configures module for calling Init callback func when input changes its
// value by a specified amount, accordingly to the configured mode (see
// SwitchToMode).
func (din *DInput10Module) ConfigEventOnValueChange(threshold uint32, event_mask uint16) error {
	ret := C.DInput10Module_configEventOnValueChange(din.ptr, C.uint32_t(threshold), C.uint32_t(event_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Resets all previously configured events.
func (din *DInput10Module) ResetEventConfig() error {
	ret := C.DInput10Module_resetEventConfig(din.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}
