#include <stdint.h>

#include "_cgo_export.h"

void dInput10Module_internal_c_cb(const uint8_t *data, const uint16_t data_len, const uint8_t id_function, void *ctx)
{
    dInput10Module_internal_go_cb((uint8_t *)data, (uint16_t)data_len, (uint8_t)id_function, ctx);
}
