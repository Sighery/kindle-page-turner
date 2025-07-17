package main

/*
#include <stdlib.h>
#include <kindlebt.h>
#include "goglue.c"
*/
import "C"

// import "unsafe"


func setGattBlob(charsValue *C.bleGattCharacteristicsValue_t, data []byte) {
    cData := C.CBytes(data)
    defer C.free(cData)

    C.setGattBlobFromBytes(charsValue, (*C.uint8_t)(cData), C.uint16_t(len(data)))
}

// func createBlobFromHexBytes(b []byte) C.bleGattBlobValue_t {
// 	length := C.uint16_t(len(b))

// 	cBuffer := C.malloc(C.size_t(length))
// 	if cBuffer == nil {
// 		panic("Failed to allocate memory")
// 	}
// 	// defer C.free(cBuffer)

// 	cSlice := (*[1 << 30]byte)(cBuffer)[:length:length]
// 	copy(cSlice, b)

// 	return C.createGattBlobFromBytes((*C.uint8_t)(cBuffer), length)
// }

// func setBlobOnCharacteristic(valuePtr *C.bleGattCharacteristicsValue_t, payload []byte) {
//     if valuePtr == nil {
//         panic("value pointer is nil")
//     }
//     if len(payload) == 0 {
//         // You could choose to reset the blob or bail out
//         return
//     }

//     length := C.uint16_t(len(payload))

//     // Allocate C buffer
//     cBuf := C.malloc(C.size_t(length))
//     if cBuf == nil {
//         panic("malloc failed")
//     }

//     // Copy Go payload into C buffer
//     cSlice := (*[1 << 30]byte)(cBuf)[:length:length]
//     copy(cSlice, payload)

//     // Cast union field to blob pointer
//     blobPtr := (*C.bleGattBlobValue_t)(unsafe.Pointer(valuePtr))

//     // Call your C function to assign and copy data
//     C.setGattBlobFromBytes(blobPtr, (*C.uint8_t)(cBuf), length)

//     // Set format to blob
//     valuePtr.format = C.BLE_FORMAT_BLOB

//     // Optional: free the temporary buffer if your C function copies it
//     C.free(cBuf)
// }
