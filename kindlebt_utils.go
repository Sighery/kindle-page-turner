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
