#ifndef GOGLUE_H
#define GOGLUE_H

#include <kindlebt.h>

// Declaration of Go callbacks. Will be defined in Go code
extern void goOnBleGattcGetGattDbCallback(bleConnHandle, bleGattsService_t*, uint32_t);

// Callback helpers. Easier to do in C and just call from Go
__attribute__((unused)) static bleGattClientCallbacks_t applicationGattcCallbacks();

#endif // GOGLUE_H
