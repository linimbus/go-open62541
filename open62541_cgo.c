#include "open62541_cgo.h"

UA_StatusCode UA_VariantToNodeValue(UA_Variant *variant, NodeValue *nodeValue) {
  nodeValue->length = variant->arrayLength;
  nodeValue->data = variant->data;
  if (variant->type == &UA_TYPES[UA_TYPES_BOOLEAN]) {
    nodeValue->type = UA_TYPES_BOOLEAN;
  } else if (variant->type == &UA_TYPES[UA_TYPES_SBYTE]) {
    nodeValue->type = UA_TYPES_SBYTE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_BYTE]) {
    nodeValue->type = UA_TYPES_BYTE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT16]) {
    nodeValue->type = UA_TYPES_INT16;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT16]) {
    nodeValue->type = UA_TYPES_UINT16;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT32]) {
    nodeValue->type = UA_TYPES_INT32;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT32]) {
    nodeValue->type = UA_TYPES_UINT32;
  } else if (variant->type == &UA_TYPES[UA_TYPES_INT64]) {
    nodeValue->type = UA_TYPES_INT64;
  } else if (variant->type == &UA_TYPES[UA_TYPES_UINT64]) {
    nodeValue->type = UA_TYPES_UINT64;
  } else if (variant->type == &UA_TYPES[UA_TYPES_FLOAT]) {
    nodeValue->type = UA_TYPES_FLOAT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DOUBLE]) {
    nodeValue->type = UA_TYPES_DOUBLE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_STRING]) {
    nodeValue->type = UA_TYPES_STRING;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DATETIME]) {
    nodeValue->type = UA_TYPES_DATETIME;
  } else if (variant->type == &UA_TYPES[UA_TYPES_GUID]) {
    nodeValue->type = UA_TYPES_GUID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_BYTESTRING]) {
    nodeValue->type = UA_TYPES_BYTESTRING;
  } else if (variant->type == &UA_TYPES[UA_TYPES_XMLELEMENT]) {
    nodeValue->type = UA_TYPES_XMLELEMENT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_NODEID]) {
    nodeValue->type = UA_TYPES_NODEID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_EXPANDEDNODEID]) {
    nodeValue->type = UA_TYPES_EXPANDEDNODEID;
  } else if (variant->type == &UA_TYPES[UA_TYPES_STATUSCODE]) {
    nodeValue->type = UA_TYPES_STATUSCODE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_QUALIFIEDNAME]) {
    nodeValue->type = UA_TYPES_QUALIFIEDNAME;
  } else if (variant->type == &UA_TYPES[UA_TYPES_LOCALIZEDTEXT]) {
    nodeValue->type = UA_TYPES_LOCALIZEDTEXT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_EXTENSIONOBJECT]) {
    nodeValue->type = UA_TYPES_EXTENSIONOBJECT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DATAVALUE]) {
    nodeValue->type = UA_TYPES_DATAVALUE;
  } else if (variant->type == &UA_TYPES[UA_TYPES_VARIANT]) {
    nodeValue->type = UA_TYPES_VARIANT;
  } else if (variant->type == &UA_TYPES[UA_TYPES_DIAGNOSTICINFO]) {
    nodeValue->type = UA_TYPES_DIAGNOSTICINFO;
  } else {
    return UA_STATUSCODE_BADTYPEMISMATCH;
  }
  return UA_STATUSCODE_GOOD;
}

uint32_t UA_VariantType(NodeValue *nodeValue) { return nodeValue->type; }

UA_Boolean UA_VariantValueBoolean(NodeValue *nodeValue, int index) {
  UA_Boolean *valueData = (UA_Boolean *)nodeValue->data;
  return valueData[index];
}

UA_SByte UA_VariantValueInt8(NodeValue *nodeValue, int index) {
  UA_SByte *valueData = (UA_SByte *)nodeValue->data;
  return valueData[index];
}

UA_Byte UA_VariantValueUint8(NodeValue *nodeValue, int index) {
  UA_Byte *valueData = (UA_Byte *)nodeValue->data;
  return valueData[index];
}

UA_Int16 UA_VariantValueInt16(NodeValue *nodeValue, int index) {
  UA_Int16 *valueData = (UA_Int16 *)nodeValue->data;
  return valueData[index];
}

UA_UInt16 UA_VariantValueUint16(NodeValue *nodeValue, int index) {
  UA_UInt16 *valueData = (UA_UInt16 *)nodeValue->data;
  return valueData[index];
}

UA_Int32 UA_VariantValueInt32(NodeValue *nodeValue, int index) {
  UA_Int32 *valueData = (UA_Int32 *)nodeValue->data;
  return valueData[index];
}

UA_UInt32 UA_VariantValueUint32(NodeValue *nodeValue, int index) {
  UA_UInt32 *valueData = (UA_UInt32 *)nodeValue->data;
  return valueData[index];
}

UA_Int64 UA_VariantValueInt64(NodeValue *nodeValue, int index) {
  UA_Int64 *valueData = (UA_Int64 *)nodeValue->data;
  return valueData[index];
}

UA_UInt64 UA_VariantValueUint64(NodeValue *nodeValue, int index) {
  UA_UInt64 *valueData = (UA_UInt64 *)nodeValue->data;
  return valueData[index];
}

UA_Float UA_VariantValueFloat(NodeValue *nodeValue, int index) {
  UA_Float *valueData = (UA_Float *)nodeValue->data;
  return valueData[index];
}

UA_Double UA_VariantValueDouble(NodeValue *nodeValue, int index) {
  UA_Double *valueData = (UA_Double *)nodeValue->data;
  return valueData[index];
}

UA_String UA_VariantValueString(NodeValue *nodeValue, int index) {
  UA_String *valueData = (UA_String *)nodeValue->data;
  return valueData[index];
}

UA_DateTime UA_VariantValueDateTime(NodeValue *nodeValue, int index) {
  UA_DateTime *valueData = (UA_DateTime *)nodeValue->data;
  return valueData[index];
}

NodeTree *UA_NodeTree_init(NodeTree *parent, uint32_t level, uint32_t index,
                           void *nodeID, size_t length) {
  NodeTree *node = (NodeTree *)malloc(sizeof(NodeTree));
  if (node == NULL) {
    return NULL;
  }
  node->nodeID = (char *)malloc(length + 1);
  if (node->nodeID == NULL) {
    free(node);
    return NULL;
  }

  memset(node->nodeID, '\0', length + 1);
  memcpy(node->nodeID, nodeID, length);

  node->level = level;
  node->index = index;

  node->parent = parent;
  node->next = NULL;
  node->tail = NULL;
  node->head = NULL;

  if (parent != NULL) {
    if (parent->head == NULL) {
      parent->head = node;
    } else {
      parent->tail->next = node;
    }
    parent->tail = node;
  }

  return node;
}

void UA_NodeTree_clear(NodeTree *nodeTree) {
  if (nodeTree->head) {
    for (NodeTree *current = nodeTree->head;;) {
      NodeTree *next = current->next;
      UA_NodeTree_clear(current);
      if (next == NULL) {
        break;
      }
      current = next;
    }
  }

  free(nodeTree->nodeID);
  free(nodeTree);
}

NodeTree *UA_NodeTree_next(NodeTree *nodeTree) { return nodeTree->next; }

NodeTree *UA_NodeTree_head(NodeTree *nodeTree) { return nodeTree->head; }

void UA_BrowseNodeTreeLevel(UA_Client *client, UA_NodeId nodeId,
                            NodeTree *parent, uint32_t level) {
  UA_BrowseRequest bReq;
  UA_BrowseRequest_init(&bReq);
  bReq.requestedMaxReferencesPerNode = 0;
  bReq.nodesToBrowse = UA_BrowseDescription_new();
  bReq.nodesToBrowseSize = 1;

  UA_NodeId_copy(&nodeId, &bReq.nodesToBrowse[0].nodeId);
  bReq.nodesToBrowse[0].resultMask = UA_BROWSERESULTMASK_ALL;

  UA_BrowseResponse bResp = UA_Client_Service_browse(client, bReq);

  if (bResp.responseHeader.serviceResult != UA_STATUSCODE_GOOD) {
    // printf debug log
  }

  for (int i = 0; i < bResp.resultsSize; i++) {
    for (int j = 0; j < bResp.results[i].referencesSize; j++) {

      NodeTree *node = NULL;

      UA_ReferenceDescription *ref = &(bResp.results[i].references[j]);
      if ((ref->nodeClass == UA_NODECLASS_OBJECT ||
           ref->nodeClass == UA_NODECLASS_VARIABLE ||
           ref->nodeClass == UA_NODECLASS_METHOD)) {
        if (ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_NUMERIC) {

          node = UA_NodeTree_init(
              parent, level, ref->nodeId.nodeId.namespaceIndex,
              ref->browseName.name.data, ref->browseName.name.length);

          if (node != NULL) {
            UA_BrowseNodeTreeLevel(
                client,
                UA_NODEID_NUMERIC(ref->nodeId.nodeId.namespaceIndex,
                                  ref->nodeId.nodeId.identifier.numeric),
                node, level + 1);
          }
        } else if (ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_STRING) {

          node =
              UA_NodeTree_init(parent, level, ref->nodeId.nodeId.namespaceIndex,
                               ref->nodeId.nodeId.identifier.string.data,
                               ref->nodeId.nodeId.identifier.string.length);

          if (node != NULL) {
            UA_BrowseNodeTreeLevel(
                client,
                UA_NODEID_STRING(ref->nodeId.nodeId.namespaceIndex,
                                 node->nodeID),
                node, level + 1);
          }
        }
      }
    }
  }

  UA_BrowseResponse_clear(&bResp);
}

NodeTree *UA_BrowseNodeTree(UA_Client *client) {
  NodeTree *root = UA_NodeTree_init(NULL, 0, 0, (void *)"", 0);

  if (root != NULL) {
    UA_BrowseNodeTreeLevel(client, UA_NODEID_NUMERIC(0, UA_NS0ID_OBJECTSFOLDER),
                           root, 1);
  }

  return root;
}