#include <kindlebt_application.h>

#include "goglue.h"


__attribute__((unused)) static bleGattClientCallbacks_t applicationGattcCallbacks() {
    application_gatt_client_callbacks.on_ble_gattc_get_gatt_db_cb =
        goOnBleGattcGetGattDbCallback;
    return application_gatt_client_callbacks;
}
