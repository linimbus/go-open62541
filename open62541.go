package go_open62541

/*
#cgo CFLAGS: -I. -std=c99
#cgo LDFLAGS: -lws2_32 -lIphlpapi

#include "open62541.h"
*/
import "C"

import (
	"fmt"
)

func UAClientStartup() {
	fmt.Println("Hello from Go!")
	client := C.UA_Client_new()
	C.UA_Client_delete(client)
}
