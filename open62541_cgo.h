#ifndef OPEN62541_CGO_H_
#define OPEN62541_CGO_H_

#include "open62541.h"
#include <stdio.h>
#include <stdlib.h>

typedef struct nodeValue {
  uint32_t type;
  uint32_t length;
  void *data;
} NodeValue;

typedef struct nodeTree {
  uint32_t level;
  uint32_t index;
  char *nodeID;

  struct nodeTree *parent;
  struct nodeTree *next;

  struct nodeTree *head;
  struct nodeTree *tail;
} NodeTree;

extern uint32_t UA_VariantType(NodeValue *nodeValue);

extern UA_Boolean UA_VariantValueBoolean(NodeValue *nodeValue, int index);

extern UA_SByte UA_VariantValueInt8(NodeValue *nodeValue, int index);

extern UA_Byte UA_VariantValueUint8(NodeValue *nodeValue, int index);

extern UA_Int16 UA_VariantValueInt16(NodeValue *nodeValue, int index);

extern UA_UInt16 UA_VariantValueUint16(NodeValue *nodeValue, int index);

extern UA_Int32 UA_VariantValueInt32(NodeValue *nodeValue, int index);

extern UA_UInt32 UA_VariantValueUint32(NodeValue *nodeValue, int index);

extern UA_Int64 UA_VariantValueInt64(NodeValue *nodeValue, int index);

extern UA_UInt64 UA_VariantValueUint64(NodeValue *nodeValue, int index);

extern UA_Float UA_VariantValueFloat(NodeValue *nodeValue, int index);

extern UA_Double UA_VariantValueDouble(NodeValue *nodeValue, int index);

extern UA_String UA_VariantValueString(NodeValue *nodeValue, int index);

extern UA_DateTime UA_VariantValueDateTime(NodeValue *nodeValue, int index);

extern UA_StatusCode UA_VariantToNodeValue(UA_Variant *variant,
                                           NodeValue *nodeValue);

extern NodeTree *UA_BrowseNodeTree(UA_Client *client);

extern void UA_NodeTree_clear(NodeTree *nodeTree);

extern NodeTree *UA_NodeTree_next(NodeTree *nodeTree);

extern NodeTree *UA_NodeTree_head(NodeTree *nodeTree);

#endif
