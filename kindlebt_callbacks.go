package main

/*
#include <kindlebt.h>
#include "goglue.c"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export goOnBleGattcGetGattDbCallback
func goOnBleGattcGetGattDbCallback(connHandle C.bleConnHandle, gattService *C.bleGattsService_t, noSvc C.uint32_t) {
	fmt.Println("Called Go callback! goOnBleGattcGetGattDbCallback")
}

//export goOnBleGattcNotifyCharsCallback
func goOnBleGattcNotifyCharsCallback(connHandle C.bleConnHandle, charsValue C.bleGattCharacteristicsValue_t) {
	fmt.Println("Called Go callback! goOnBleGattcNotifyCharsCallback")

	// Anonymous unions seem like a pain to use from Golang
	base := uintptr(unsafe.Pointer(&charsValue))

	switch charsValue.format {
	case C.BLE_FORMAT_UINT8:
		val := *(*C.uint8_t)(unsafe.Pointer(base))
		fmt.Printf("Received uint8: %d\n", val)
	case C.BLE_FORMAT_UINT16:
		val := *(*C.uint16_t)(unsafe.Pointer(base))
		fmt.Printf("Received uint16: %d\n", val)
	case C.BLE_FORMAT_UINT32:
		val := *(*C.uint32_t)(unsafe.Pointer(base))
		fmt.Printf("Received uint32: %d\n", val)
	case C.BLE_FORMAT_SINT8:
		val := *(*C.int8_t)(unsafe.Pointer(base))
		fmt.Printf("Received int8: %d\n", val)
	case C.BLE_FORMAT_SINT16:
		val := *(*C.int16_t)(unsafe.Pointer(base))
		fmt.Printf("Received int16: %d\n", val)
	case C.BLE_FORMAT_SINT32:
		val := *(*C.int32_t)(unsafe.Pointer(base))
		fmt.Printf("Received int32: %d\n", val)
	case C.BLE_FORMAT_BLOB:
		val := *(*C.bleGattBlobValue_t)(unsafe.Pointer(base))
		dataSlice := C.GoBytes(unsafe.Pointer(val.data), C.int(val.size))

		if isASCIIPrintable(dataSlice) {
			fmt.Println("Received text: ", string(dataSlice))
		} else {
			fmt.Printf("Received binary: %x\n", dataSlice)
		}
	}
}
