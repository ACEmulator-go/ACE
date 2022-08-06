package enums


type EnvCellFlags int32

const (
    EnvCellFlagsSeenOutside EnvCellFlags = EnvCellFlags(0x1)
    EnvCellFlagsHasStaticObjs EnvCellFlags = EnvCellFlags(0x2)
    )
