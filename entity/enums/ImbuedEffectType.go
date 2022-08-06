package enums


type ImbuedEffectType uint32

const (
    ImbuedEffectTypeUndef ImbuedEffectType = ImbuedEffectType(0)
    ImbuedEffectTypeCriticalStrike ImbuedEffectType = ImbuedEffectType(0x0001)
    ImbuedEffectTypeCripplingBlow ImbuedEffectType = ImbuedEffectType(0x0002)
    ImbuedEffectTypeArmorRending ImbuedEffectType = ImbuedEffectType(0x0004)
    ImbuedEffectTypeSlashRending ImbuedEffectType = ImbuedEffectType(0x0008)
    ImbuedEffectTypePierceRending ImbuedEffectType = ImbuedEffectType(0x0010)
    ImbuedEffectTypeBludgeonRending ImbuedEffectType = ImbuedEffectType(0x0020)
    ImbuedEffectTypeAcidRending ImbuedEffectType = ImbuedEffectType(0x0040)
    ImbuedEffectTypeColdRending ImbuedEffectType = ImbuedEffectType(0x0080)
    ImbuedEffectTypeElectricRending ImbuedEffectType = ImbuedEffectType(0x0100)
    ImbuedEffectTypeFireRending ImbuedEffectType = ImbuedEffectType(0x0200)
    ImbuedEffectTypeMeleeDefense ImbuedEffectType = ImbuedEffectType(0x0400)
    ImbuedEffectTypeMissileDefense ImbuedEffectType = ImbuedEffectType(0x0800)
    ImbuedEffectTypeMagicDefense ImbuedEffectType = ImbuedEffectType(0x1000)
    ImbuedEffectTypeSpellbook ImbuedEffectType = ImbuedEffectType(0x2000)
    ImbuedEffectTypeNetherRending ImbuedEffectType = ImbuedEffectType(0x4000)
    ImbuedEffectTypeIgnoreSomeMagicProjectileDamage ImbuedEffectType = ImbuedEffectType(0x20000000)
    ImbuedEffectTypeAlwaysCritical ImbuedEffectType = ImbuedEffectType(0x40000000)
    )
