package go_open62541

/*
#cgo CFLAGS: -I. -std=c99
#cgo LDFLAGS: -lm -lrt -lpthread

#include <stdlib.h>
#include <stdio.h>
#include "open62541.h"
#include "open62541_cgo.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type Client struct {
	cli uintptr
}

type NodeInfo struct {
	Index  uint32
	NodeID string
}

type NodeTree struct {
	Level    uint32
	Node     NodeInfo
	SubNodes []*NodeTree
}

type UAType int

const (
	UA_BOOLEAN UAType = iota
	UA_INT8
	UA_UINT8
	UA_INT16
	UA_UINT16
	UA_INT32
	UA_UINT32
	UA_INT64
	UA_UINT64
	UA_FLOAT
	UA_DOUBLE
	UA_STRING
	UA_DATETIME
	// UA_GUID
	// UA_BYTESTRING
	// UA_XMLELEMENT
	// UA_NODEID
	// UA_EXPANDEDNODEID
	// UA_STATUSCODE
	// UA_QUALIFIEDNAME
	// UA_LOCALIZEDTEXT
	// UA_EXTENSIONOBJECT
	// UA_DATAVALUE
	// UA_VARIANT
	// UA_DIAGNOSTICINFO
)

type NodeValue struct {
	Type   UAType
	Length uint32
	Value  interface{} // []bool, []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []float, []double, []string, []time.Time
}

func (v *NodeValue) ToString() string {
	switch v.Type {
	case UA_BOOLEAN:
		return fmt.Sprintf("%v", v.Value.([]bool))
	case UA_INT8:
		return fmt.Sprintf("%v", v.Value.([]int8))
	case UA_UINT8:
		return fmt.Sprintf("%v", v.Value.([]uint8))
	case UA_INT16:
		return fmt.Sprintf("%v", v.Value.([]int16))
	case UA_UINT16:
		return fmt.Sprintf("%v", v.Value.([]uint16))
	case UA_INT32:
		return fmt.Sprintf("%v", v.Value.([]int32))
	case UA_UINT32:
		return fmt.Sprintf("%v", v.Value.([]uint32))
	case UA_INT64:
		return fmt.Sprintf("%v", v.Value.([]int64))
	case UA_UINT64:
		return fmt.Sprintf("%v", v.Value.([]uint64))
	case UA_FLOAT:
		return fmt.Sprintf("%v", v.Value.([]float32))
	case UA_DOUBLE:
		return fmt.Sprintf("%v", v.Value.([]float64))
	case UA_STRING:
		return fmt.Sprintf("%v", v.Value.([]string))
	default:
		return ""
	}
}

func NewNodeValue(value *C.NodeValue) *NodeValue {
	arrayLength := uint32(value.length)

	if arrayLength == 0 {
		data := uintptr(unsafe.Pointer(value.data))
		if data == 0 {
			return nil
		}
		if data > uintptr(unsafe.Pointer(C.UA_EMPTY_ARRAY_SENTINEL)) {
			arrayLength = 1
		}
	}

	switch C.UA_VariantType(value) {
	case C.UA_TYPES_BOOLEAN:
		arrayList := make([]bool, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = bool(C.UA_VariantValueBoolean(value, C.int(i)))
		}
		return &NodeValue{Type: UA_BOOLEAN, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_SBYTE:
		arrayList := make([]int8, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = int8(C.UA_VariantValueInt8(value, C.int(i)))
		}
		return &NodeValue{Type: UA_INT8, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_BYTE:
		arrayList := make([]uint8, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = uint8(C.UA_VariantValueUint8(value, C.int(i)))
		}
		return &NodeValue{Type: UA_UINT8, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_INT16:
		arrayList := make([]int16, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = int16(C.UA_VariantValueInt16(value, C.int(i)))
		}
		return &NodeValue{Type: UA_INT16, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_UINT16:
		arrayList := make([]uint16, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = uint16(C.UA_VariantValueUint16(value, C.int(i)))
		}
		return &NodeValue{Type: UA_UINT16, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_INT32:
		arrayList := make([]int32, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = int32(C.UA_VariantValueInt32(value, C.int(i)))
		}
		return &NodeValue{Type: UA_INT32, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_UINT32:
		arrayList := make([]uint32, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = uint32(C.UA_VariantValueUint32(value, C.int(i)))
		}
		return &NodeValue{Type: UA_UINT32, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_INT64:
		arrayList := make([]int64, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = int64(C.UA_VariantValueInt64(value, C.int(i)))
		}
		return &NodeValue{Type: UA_INT64, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_DATETIME:
		fallthrough
	case C.UA_TYPES_UINT64:
		arrayList := make([]uint64, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = uint64(C.UA_VariantValueUint64(value, C.int(i)))
		}
		return &NodeValue{Type: UA_UINT64, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_FLOAT:
		arrayList := make([]float32, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = float32(C.UA_VariantValueFloat(value, C.int(i)))
		}
		return &NodeValue{Type: UA_FLOAT, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_DOUBLE:
		arrayList := make([]float64, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			arrayList[i] = float64(C.UA_VariantValueDouble(value, C.int(i)))
		}
		return &NodeValue{Type: UA_DOUBLE, Length: arrayLength, Value: arrayList}
	case C.UA_TYPES_STRING:
		arrayList := make([]string, arrayLength)
		for i := 0; i < int(arrayLength); i++ {
			cString := C.UA_VariantValueString(value, C.int(i))
			goBytes := C.GoBytes(unsafe.Pointer(cString.data), C.int(cString.length))
			arrayList[i] = string(goBytes)
		}
		return &NodeValue{Type: UA_STRING, Length: arrayLength, Value: arrayList}
	default:
		return nil
	}
}

func NewClient(add string) (*Client, error) {
	client := C.UA_Client_new()
	if client == nil {
		return nil, errors.New("ua client create failed")
	}
	C.UA_ClientConfig_setDefault(C.UA_Client_getConfig(client))

	cStr := C.CString(add)
	defer C.free(unsafe.Pointer(cStr))

	retval := C.UA_Client_connect(client, cStr)
	if retval != C.UA_STATUSCODE_GOOD {
		C.UA_Client_delete(client)
		return nil, fmt.Errorf("ua client connect failed, retval = %d", uint32(retval))
	}

	return &Client{cli: uintptr(unsafe.Pointer(client))}, nil
}

func (c *Client) Close() {
	if c.cli != 0 {
		client := (*C.UA_Client)(unsafe.Pointer(c.cli))
		C.UA_Client_disconnect(client)
		C.UA_Client_delete(client)
		c.cli = 0
	}
}

func (c *Client) ReadNode(node NodeInfo) (*NodeValue, error) {
	cID := C.CString(node.NodeID)
	defer C.free(unsafe.Pointer(cID))

	client := (*C.UA_Client)(unsafe.Pointer(c.cli))

	var variant C.UA_Variant
	retval := C.UA_Client_readValueAttribute(client, C.UA_NODEID_STRING(C.UA_UInt16(node.Index), cID), &variant)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client read value failed, retval = %d", uint32(retval))
	}

	defer C.UA_Variant_clear(&variant)

	var cNodeValue C.NodeValue
	cNodeValue.length = 0

	retval = C.UA_VariantToNodeValue(&variant, &cNodeValue)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client parse node value failed, retval = %d", uint32(retval))
	}

	nodeValue := NewNodeValue(&cNodeValue)
	if nodeValue == nil {
		return nil, fmt.Errorf("ua client parse node data type not support, type = %d", uint32(C.UA_VariantType(&cNodeValue)))
	}

	return nodeValue, nil
}

func UA_NodeTree(cNodeTree *C.NodeTree) *NodeTree {
	subNodes := make([]*NodeTree, 0)
	node := C.UA_NodeTree_head(cNodeTree)
	for {
		if unsafe.Pointer(node) == nil {
			break
		}
		subNodes = append(subNodes, UA_NodeTree(node))
		node = C.UA_NodeTree_next(node)
	}
	return &NodeTree{Level: uint32(cNodeTree.level), Node: NodeInfo{Index: uint32(cNodeTree.index), NodeID: C.GoString(cNodeTree.nodeID)}, SubNodes: subNodes}
}

func (c *Client) BrowseNode() (*NodeTree, error) {
	client := (*C.UA_Client)(unsafe.Pointer(c.cli))
	cNodeTree := C.UA_BrowseNodeTree(client)
	defer C.UA_NodeTree_clear(cNodeTree)

	return UA_NodeTree(cNodeTree), nil
}

// func (c *Client) WriteNode(node NodeInfo, value NodeValue) error {

// 	cID := C.CString(node.ID)
// 	defer C.free(unsafe.Pointer(cID))

// 	client := (*C.UA_Client)(unsafe.Pointer(c.cli))

// 	var vts C.UAVariantToString
// 	vts.buildValue(value)

// 	var variant C.UA_Variant
// 	C.UA_Variant_copy(&vts.value, &variant)

// 	retval := C.UA_Client_writeValueAttribute(client, C.UA_NODEID_STRING(C.uint32(node.Index), cID), &variant)
// 	if retval != C.UA_STATUSCODE_GOOD {
// 		return fmt.Errorf("ua client write value failed, retval = %d", uint32(retval))
// 	}

// 	C.UA_Variant_clear(&variant)

// 	return nil
// }
