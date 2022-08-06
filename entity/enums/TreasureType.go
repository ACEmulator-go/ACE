package enums


type TreasureType int32

const (
    TreasureTypeUndef TreasureType = TreasureType(0)
    TreasureTypeItem TreasureType = TreasureType(1)
    TreasureTypeMagicItem TreasureType = TreasureType(2)
    TreasureTypeMundaneItem TreasureType = TreasureType(3)
    )
