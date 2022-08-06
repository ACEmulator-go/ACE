package enums


type TransferFlags int32

const (
    TransferFlagsCasterSource TransferFlags = TransferFlags(0x1)
    TransferFlagsTargetSource TransferFlags = TransferFlags(0x2)
    TransferFlagsCasterDestination TransferFlags = TransferFlags(0x4)
    )
