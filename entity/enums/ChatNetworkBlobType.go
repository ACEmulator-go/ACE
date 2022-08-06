package enums


type ChatNetworkBlobType int32

const (
    ChatNetworkBlobTypeNETBLOB_UNKNOWN ChatNetworkBlobType = ChatNetworkBlobType(0)
    ChatNetworkBlobTypeNETBLOB_EVENT_BINARY ChatNetworkBlobType = ChatNetworkBlobType(1)
    ChatNetworkBlobTypeNETBLOB_EVENT_XMLRPC ChatNetworkBlobType = ChatNetworkBlobType(2)
    ChatNetworkBlobTypeNETBLOB_REQUEST_BINARY ChatNetworkBlobType = ChatNetworkBlobType(3)
    ChatNetworkBlobTypeNETBLOB_REQUEST_XMLRPC ChatNetworkBlobType = ChatNetworkBlobType(4)
    ChatNetworkBlobTypeNETBLOB_RESPONSE_BINARY ChatNetworkBlobType = ChatNetworkBlobType(5)
    ChatNetworkBlobTypeNETBLOB_RESPONSE_XMLRPC ChatNetworkBlobType = ChatNetworkBlobType(6)
    )
