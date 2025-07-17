package main

/*
#include <stdlib.h>
#include <kindlebt.h>
#include "goglue.c"
*/
import "C"

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

var ledStatus bool

func main() {
	fmt.Println("Hello World from Kindle and Golang!")

	isBLE := C.isBLESupported()
	fmt.Printf("Is BLE supported? %t\n", isBLE)

	supportedSession := C.getSupportedSession()
	fmt.Printf("Supported session %d\n", supportedSession)

	// Dropping privileges
	fmt.Println("Dropping privileges")
	err := UseBluetoothPrivileges()
	if err != nil {
		panic(err)
	}

	var btSession C.sessionHandle
	fmt.Println(reflect.TypeOf(btSession))

	sessionResult := C.openSession(C.ACEBT_SESSION_TYPE_DUAL_MODE, &btSession)
	defer C.closeSession(btSession)

	fmt.Printf("Opening session result %d\n", sessionResult)
	fmt.Printf("btSession: %+v\n", btSession)

	bleResult := C.bleRegister(btSession)
	defer C.bleDeregister(btSession)
	fmt.Printf("Opening BLE result %d\n", bleResult)

	if bleResult != 0 {
		panic("Couldn't open BLE")
	}

	gattcCb := C.applicationGattcCallbacks()

	gattcResult := C.bleRegisterGattClient(btSession, &gattcCb)
	defer C.bleDeregisterGattClient(btSession)
	fmt.Printf("Register GATT Client result %d\n", gattcResult)

	bleAddrStr := C.CString("2C:CF:67:B8:DC:3F")
	defer C.free(unsafe.Pointer(bleAddrStr))
	var bleAddr C.bdAddr_t
	bdaddrconvStatus := C.utilsConvertStrToBdAddr(bleAddrStr, &bleAddr)
	fmt.Printf("Converted str BT addr %s, result %d\n", bleAddrStr, bdaddrconvStatus)

	var connHandle C.bleConnHandle
	connStatus := C.bleConnect(
		btSession, &connHandle, &bleAddr, C.ACE_BT_BLE_CONN_PARAM_BALANCED,
		C.ACEBT_BLE_GATT_CLIENT_ROLE, C.ACE_BT_BLE_CONN_PRIO_MEDIUM,
	)
	defer C.bleDisconnect(connHandle)

	fmt.Printf("Connection status %d\n", connStatus)

	var gattDb C.bleGattsService_t
	dbStatus := C.bleGetDatabase(connHandle, &gattDb)
	fmt.Printf("GATT DB status %d\n", dbStatus)
	fmt.Printf("%+v\n", gattDb)

	// Pico LED characteristic
	uuidStr := "ff120000000000000000000000000000"
	if len(uuidStr) % 2 != 0 {
		panic("UUID hex string must have an even length")
	}
	uuidLen := len(uuidStr) / 2

	characUuidStr := C.CString(uuidStr)
	defer C.free(unsafe.Pointer(characUuidStr))

	var characUuid C.uuid_t

	if C.utilsConvertHexStrToByteArray(characUuidStr, (*C.uint8_t)(unsafe.Pointer(&characUuid.uu[0]))) == 0 {
		panic("Failed to convert Characteristic UUID string")
	}

	switch uuidLen {
	case 2:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_16)
	case 4:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_32)
	case 16:
		C.setUUIDType(&characUuid, C.ACEBT_UUID_TYPE_128)
	default:
		panic(fmt.Sprintf("Unsupported UUID length: %d", uuidLen))
	}

	characRec := C.utilsFindCharRec(characUuid, C.uint8_t(uuidLen))
	if characRec == nil {
		panic("Couldn't find characteristic. Did you not get the GATT Database?")
	}

	fmt.Println("Enabling notification on PICO LED Characteristic")
	notificationStatus := C.bleSetNotification(
		btSession, connHandle, characRec.value, C.bool(true),
	)
	fmt.Printf("Enabled notification: %d\n", notificationStatus)
	time.Sleep(5 * time.Second)
	// Disable notification
	notificationStatus = C.bleSetNotification(
		btSession, connHandle, characRec.value, C.bool(false),
	)
	fmt.Printf("Disabled notification: %d\n", notificationStatus)

	time.Sleep(2 * time.Second)

	// C.bleWriteCharacteristic(
	// 	btSession, connHandle, &characRec.value,
	// 	C.ACEBT_BLE_WRITE_TYPE_RESP_REQUIRED,
	// 	// C.ACEBT_BLE_WRITE_TYPE_RESP_NO,
	// )
	// return

	for true {
		fmt.Println("Reading PICO LED Characteristic")
		readStatus := C.bleReadCharacteristic(btSession, connHandle, characRec.value)
		fmt.Printf("Read status: %d\n", readStatus)

		// Characteristic becomes busy. Need to implement some kinda semaphor or such for characRec
		time.Sleep(1 * time.Second)

		fmt.Println("Writing to LED Characteristic")
		// Reset the shared blob before writes
		C.freeGattBlob(&characRec.value)
		// C.freeGattBlob((*C.bleGattBlobValue_t)(unsafe.Pointer(&characRec.value)))

		var writeVal string
		if ledStatus == true {
			writeVal = "OFF"
		} else {
			writeVal = "ON"
		}

		fmt.Println("Setting to value", writeVal)

		fmt.Printf("Byte array in hex %x\n", []byte(writeVal))

		setGattBlob(&characRec.value, []byte(writeVal))

		// setBlobOnCharacteristic(&characRec.value, []byte(writeVal))
		// blob := createBlobFromHexBytes([]byte(writeVal))
		// cBlobPtr := (*C.bleGattBlobValue_t)(unsafe.Pointer(&characRec.value))
		// *cBlobPtr = blob
		// characRec.value.format = C.BLE_FORMAT_BLOB

		writeStatus := C.bleWriteCharacteristic(
			btSession, connHandle, &characRec.value,
			C.ACEBT_BLE_WRITE_TYPE_RESP_REQUIRED,
			// C.ACEBT_BLE_WRITE_TYPE_RESP_NO,
		)
		fmt.Printf("Write result %d\n", writeStatus)
		// C.freeGattBlob((*C.bleGattBlobValue_t)(unsafe.Pointer(&characRec.value)))

		time.Sleep(2 * time.Second)

		C.freeGattBlob(&characRec.value)
	}


	time.Sleep(5 * time.Second)
	fmt.Println("Finishing program")
}
