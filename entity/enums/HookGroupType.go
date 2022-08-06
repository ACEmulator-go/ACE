package enums


type HookGroupType int32

const (
    HookGroupTypeUndef HookGroupType = HookGroupType(0x0)
    HookGroupTypeNoisemakingItems HookGroupType = HookGroupType(0x1)
    HookGroupTypeTestItems HookGroupType = HookGroupType(0x2)
    HookGroupTypePortalItems HookGroupType = HookGroupType(0x4)
    HookGroupTypeWritableItems HookGroupType = HookGroupType(0x8)
    HookGroupTypeSpellCastingItems HookGroupType = HookGroupType(0x10)
    )
