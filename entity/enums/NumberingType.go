package enums


type NumberingType byte

const (
    NumberingTypeUndefined NumberingType = NumberingType(0x0)
    NumberingTypeNormal NumberingType = NumberingType(0x1)
    NumberingTypeSequential NumberingType = NumberingType(0x1)
    NumberingTypeBitfield NumberingType = NumberingType(0x2)
    NumberingTypeBitfield32 NumberingType = NumberingType(0x3)
    )
