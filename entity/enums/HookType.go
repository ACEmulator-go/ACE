package enums


type HookType int32

const (
    HookTypeUndef HookType = HookType(0x00)
    HookTypeFloor HookType = HookType(0x01)
    HookTypeWall HookType = HookType(0x02)
    HookTypeCeiling HookType = HookType(0x04)
    HookTypeYard HookType = HookType(0x08)
    )
