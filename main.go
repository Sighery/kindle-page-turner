package main

// #include <stdlib.h>
// #include "kindlebt.h"
import "C"

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"strconv"
	"time"
)

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

func main() {
	fmt.Println("Hello World from Kindle and Golang!")

	isBLE := C.isBLESupported()
	fmt.Printf("Is BLE supported? %t\n", isBLE)

	supportedSession := C.getSupportedSession()
	fmt.Printf("Supported session %d\n", supportedSession)

	// Dropping privileges
	fmt.Println("Dropping privileges")

	groupId, err := strconv.Atoi(getEnv("BLUETOOTH_GROUP_ID", "1003"))
	if err != nil {
		panic(err)
	}
	userId, err := strconv.Atoi(getEnv("BLUETOOTH_USER_ID", "1003"))
	if err != nil {
		panic(err)
	}

	err = syscall.Setgid(groupId)
	if err != nil {
		panic(err)
	}
	err = syscall.Setuid(userId)
	if err != nil {
		panic(err)
	}

	var bt_session C.sessionHandle
	fmt.Println(reflect.TypeOf(bt_session))

	sessionResult := C.openSession(C.ACEBT_SESSION_TYPE_DUAL_MODE, &bt_session)
	fmt.Printf("Opening session result %d\n", sessionResult)
	fmt.Printf("bt_session: %+v\n", bt_session)

	time.Sleep(8 * time.Second)

	fmt.Println("Finishing program")

	defer C.closeSession(bt_session)
}
