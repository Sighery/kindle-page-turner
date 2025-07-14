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

	var gattDb C.bleGattsService_t;
	dbStatus := C.bleGetDatabase(connHandle, &gattDb)
	fmt.Printf("GATT DB status %d\n", dbStatus)
	fmt.Printf("%+v\n", gattDb)

	time.Sleep(8 * time.Second)

	fmt.Println("Finishing program")
}
