package Relay8Module

/*
#cgo LDFLAGS: -lstdc++ -lrt -lemod_controller

#include <Relay8Module.h>
*/
import "C"
import "fmt"

const (
	RELAY1           byte = byte(C.Relay8Module_RELAY1)
	RELAY2           byte = byte(C.Relay8Module_RELAY2)
	RELAY3           byte = byte(C.Relay8Module_RELAY3)
	RELAY4           byte = byte(C.Relay8Module_RELAY4)
	RELAY5           byte = byte(C.Relay8Module_RELAY5)
	RELAY6           byte = byte(C.Relay8Module_RELAY6)
	RELAY7           byte = byte(C.Relay8Module_RELAY7)
	RELAY8           byte = byte(C.Relay8Module_RELAY8)
	ALL_RELAY        byte = byte(C.Relay8Module_ALL_RELAY)
	NUMBER_OF_RELAYS byte = byte(C.Relay8Module_NUMBER_OF_RELAYS)
)

// This struct represents the Signal Relays Module connected via eMOD bus to
// the controller module. The Relays Module has 8 independent signal relays
// that you can activate or deactivate each one separately. Moreover, a timeout
// configuration parameter can be set to generate one shot relay pulse
// automatically. If timeout is set to 0, relay stays in the current state.
type Relay8Module struct {
	ptr *C.Relay8Module_t
}

// Allocates internal memory and returns an instance of the module
func NewRelay8Module() *Relay8Module {
	m := new(Relay8Module)
	m.ptr = C.Relay8Module_create()

	return m
}

// Free the internal memory of the module. The instance can't be used anymore
// after that.
func (r *Relay8Module) Destroy() {
	C.Relay8Module_destroy(r.ptr)
}

// Initialize module. It is important to emphasize that this method is
// mandatory to be called. It can be called several times, each of which the
// module is initialized again.
func (r *Relay8Module) Init() error {
	ret := C.Relay8Module_init(r.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Same as Init(), but allows to initialize a specific module number
// instead of the first one.
func (r *Relay8Module) InitModuleNumber(module_number uint8) error {
	ret := C.Relay8Module_init_v(r.ptr, C.uint8_t(module_number))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Configure pulse width of a relay o list of relays.
// When pulse width is 0 a relay stays in the current state, and there is no
// pulse.
func (r *Relay8Module) ConfigPulseWidth(relay_mask byte, width_ms uint32) error {
	ret := C.Relay8Module_configPulseWidth(r.ptr, C.uint8_t(relay_mask), C.uint32_t(width_ms))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Activates a relay or set of relays.
func (r *Relay8Module) Activate(relay_mask byte) error {
	ret := C.Relay8Module_activate(r.ptr, C.uint8_t(relay_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Deactivates a relay or set of relays.
func (r *Relay8Module) Deactivate(relay_mask byte) error {
	ret := C.Relay8Module_deactivate(r.ptr, C.uint8_t(relay_mask))
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Activates all module relays.
func (r *Relay8Module) ActivateAll() error {
	ret := C.Relay8Module_activateAll(r.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Deactivates all module relays.
func (r *Relay8Module) DeactivateAll() error {
	ret := C.Relay8Module_deactivateAll(r.ptr)
	if ret != 0 {
		return fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return nil
}

// Gets a relay status.
func (r *Relay8Module) GetRelayStatus(relay_mask byte) (byte, error) {
	var status C.uint8_t
	ret := C.Relay8Module_getRelayStatus(r.ptr, C.uint8_t(relay_mask), &status)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return byte(status), nil
}

// Gets all relay status.
func (r *Relay8Module) GetAllRelayStatus() (byte, error) {
	var status C.uint8_t
	ret := C.Relay8Module_getAllRelayStatus(r.ptr, &status)
	if ret != 0 {
		return 0, fmt.Errorf("EmodRet: %d", uint32(ret))
	}

	return byte(status), nil
}
