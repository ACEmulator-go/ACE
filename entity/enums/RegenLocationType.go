package enums


type RegenLocationType uint32

const (
    RegenLocationTypeUndef RegenLocationType = RegenLocationType(0x00)
    RegenLocationTypeOnTop RegenLocationType = RegenLocationType(0x01)
    RegenLocationTypeScatter RegenLocationType = RegenLocationType(0x02)
    RegenLocationTypeSpecific RegenLocationType = RegenLocationType(0x04)
    RegenLocationTypeContain RegenLocationType = RegenLocationType(0x08)
    RegenLocationTypeWield RegenLocationType = RegenLocationType(0x10)
    RegenLocationTypeShop RegenLocationType = RegenLocationType(0x20)
    RegenLocationTypeTreasure RegenLocationType = RegenLocationType(0x40)
    )
