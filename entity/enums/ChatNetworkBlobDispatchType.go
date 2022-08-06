package enums


type ChatNetworkBlobDispatchType int32

const (
    ChatNetworkBlobDispatchTypeASYNCMETHOD_UNKNOWN ChatNetworkBlobDispatchType = ChatNetworkBlobDispatchType(0)
    ChatNetworkBlobDispatchTypeASYNCMETHOD_SENDTOROOMBYNAME ChatNetworkBlobDispatchType = ChatNetworkBlobDispatchType(1)
    ChatNetworkBlobDispatchTypeASYNCMETHOD_SENDTOROOMBYID ChatNetworkBlobDispatchType = ChatNetworkBlobDispatchType(2)
    ChatNetworkBlobDispatchTypeASYNCMETHOD_CREATEROOM ChatNetworkBlobDispatchType = ChatNetworkBlobDispatchType(3)
    ChatNetworkBlobDispatchTypeASYNCMETHOD_INVITECLIENTTOROOMBYID ChatNetworkBlobDispatchType = ChatNetworkBlobDispatchType(4)
    )
