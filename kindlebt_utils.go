package main

/*
#include <stdlib.h>
#include <kindlebt.h>
#include "goglue.c"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)


func setGattBlob(charsValue *C.bleGattCharacteristicsValue_t, data []byte) {
	cData := C.CBytes(data)
	defer C.free(cData)

	C.setGattBlobFromBytes(charsValue, (*C.uint8_t)(cData), C.uint16_t(len(data)))
}

func uuidToString(uuid C.uuid_t) string {
	bytes := C.GoBytes(unsafe.Pointer(&uuid.uu[0]), C.int(16))
	return fmt.Sprintf("%x", bytes)
}

func uuidStrToUuidC(uuidStr string) (C.uuid_t, int, error) {
	var characUuid C.uuid_t
	uuidLen := 0

	if len(uuidStr) % 2 != 0 {
		return characUuid, uuidLen, errors.New("UUID hex string must have an even length")
	}

	uuidLen = len(uuidStr) / 2

	characUuidStr := C.CString(uuidStr)
	defer C.free(unsafe.Pointer(characUuidStr))

	if C.utilsConvertHexStrToByteArray(characUuidStr, (*C.uint8_t)(unsafe.Pointer(&characUuid.uu[0]))) == 0 {
		return characUuid, uuidLen, errors.New("Failed to convert Characteristic UUID string")
	}

	switch uuidLen {
	case 2:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_16)
	case 4:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_32)
	case 16:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_128)
	default:
		return characUuid, uuidLen, errors.New(fmt.Sprintf("Unsupported UUID length: %d", uuidLen))
	}

	return characUuid, uuidLen, nil
}
