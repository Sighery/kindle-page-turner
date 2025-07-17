#include <kindlebt_application.h>

#include "goglue.h"


__attribute__((unused)) static bleGattClientCallbacks_t applicationGattcCallbacks() {
    application_gatt_client_callbacks.on_ble_gattc_get_gatt_db_cb =
        goOnBleGattcGetGattDbCallback;
    application_gatt_client_callbacks.notify_characteristics_cb =
        goOnBleGattcNotifyCharsCallback;
    application_gatt_client_callbacks.on_ble_gattc_read_characteristics_cb =
        goOnBleGattcReadCharsCallback;
    application_gatt_client_callbacks.on_ble_gattc_write_characteristics_cb =
        goOnBleGattcWriteCharsCallback;
    return application_gatt_client_callbacks;
}

__attribute__((unused)) static void setUUIDType(uuid_t* uuid, UUIDType_t newType) {
    uuid->type = newType;
}
