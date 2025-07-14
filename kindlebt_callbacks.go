package main

/*
#include <kindlebt.h>
#include "goglue.c"
*/
import "C"

import (
	"fmt"
)

//export goOnBleGattcGetGattDbCallback
func goOnBleGattcGetGattDbCallback(connHandle C.bleConnHandle, gattService *C.bleGattsService_t, noSvc C.uint32_t) {
	fmt.Println("Called Go callback! goOnBleGattcGetGattDbCallback")
}
