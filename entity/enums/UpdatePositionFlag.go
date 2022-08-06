package enums


type UpdatePositionFlag int32

const (
    UpdatePositionFlagNone UpdatePositionFlag = UpdatePositionFlag(0x00)
    UpdatePositionFlagVelocity UpdatePositionFlag = UpdatePositionFlag(0x01)
    UpdatePositionFlagPlacement UpdatePositionFlag = UpdatePositionFlag(0x02)
    UpdatePositionFlagContact UpdatePositionFlag = UpdatePositionFlag(0x04)
    UpdatePositionFlagZeroQw UpdatePositionFlag = UpdatePositionFlag(0x08)
    UpdatePositionFlagZeroQx UpdatePositionFlag = UpdatePositionFlag(0x10)
    UpdatePositionFlagZeroQy UpdatePositionFlag = UpdatePositionFlag(0x20)
    )
