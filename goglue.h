#ifndef GOGLUE_H
#define GOGLUE_H

#include <kindlebt.h>

// Declaration of Go callbacks. Will be defined in Go code
extern void goOnBleGattcGetGattDbCallback(bleConnHandle, bleGattsService_t*, uint32_t);
extern void goOnBleGattcNotifyCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t);
extern void goOnBleGattcReadCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t, status_t);
extern void goOnBleGattcWriteCharsCallback(bleConnHandle, bleGattCharacteristicsValue_t, status_t);

// Callback helpers. Easier to do in C and just call from Go
__attribute__((unused)) static bleGattClientCallbacks_t applicationGattcCallbacks();

// Util helpers
__attribute__((unused)) static void setUUIDType(uuid_t*, UUIDType_t);

#endif // GOGLUE_H
