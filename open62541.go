package go_open62541

/*
#cgo CFLAGS: -I. -std=c99
#cgo LDFLAGS: -lws2_32 -lIphlpapi

#include "open62541.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Client struct {
	cli uintptr
}

func NewClient() (*Client, error) {
	client := C.UA_Client_new()
	if client == nil {
		return nil, errors.New("ua client create failed")
	}
	C.UA_ClientConfig_setDefault(C.UA_Client_getConfig(client))

	return &Client{cli: uintptr(unsafe.Pointer(client))}, nil
}

func (c *Client) Connect(add string) error {
	cStr := C.CString(add)
	defer C.free(unsafe.Pointer(cStr))

	client := (*C.UA_Client)(unsafe.Pointer(c.cli))
	retval := C.UA_Client_connect(client, cStr)
	if retval != C.UA_STATUSCODE_GOOD {
		return errors.New("ua client connect failed")
	}
	return nil
}

func (c *Client) Close() {
	if c.cli != 0 {
		client := (*C.UA_Client)(unsafe.Pointer(c.cli))
		C.UA_Client_disconnect(client)
		C.UA_Client_delete(client)
		c.cli = 0
	}
}
