package main

import (
	"os"
	"strconv"
	"syscall"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func UseBluetoothPrivileges() error {
	groupId, err := strconv.Atoi(GetEnv("BLUETOOTH_GROUP_ID", "1003"))
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(GetEnv("BLUETOOTH_USER_ID", "1003"))
	if err != nil {
		return err
	}

	err = syscall.Setgid(groupId)
	if err != nil {
		return err
	}
	err = syscall.Setuid(userId)
	if err != nil {
		return err
	}

	return nil
}

func isASCIIPrintable(data []byte) bool {
	for _, b := range data {
		if b < 0x20 || b > 0x7E {
			return false
		}
	}
	return true
}
