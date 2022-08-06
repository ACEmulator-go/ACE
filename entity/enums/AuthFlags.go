package enums


type AuthFlags int32

const (
    AuthFlagsNone AuthFlags = AuthFlags(0x0)
    AuthFlagsEnableCrypto AuthFlags = AuthFlags(0x1)
    AuthFlagsExtraData AuthFlags = AuthFlags(0x4)
    )
