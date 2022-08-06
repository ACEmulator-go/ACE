package enums


type IdentifyResponseFlags int32

const (
    IdentifyResponseFlagsNone IdentifyResponseFlags = IdentifyResponseFlags(0x0000)
    IdentifyResponseFlagsIntStatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x0001)
    IdentifyResponseFlagsBoolStatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x0002)
    IdentifyResponseFlagsFloatStatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x0004)
    IdentifyResponseFlagsStringStatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x0008)
    IdentifyResponseFlagsSpellBook IdentifyResponseFlags = IdentifyResponseFlags(0x0010)
    IdentifyResponseFlagsWeaponProfile IdentifyResponseFlags = IdentifyResponseFlags(0x0020)
    IdentifyResponseFlagsHookProfile IdentifyResponseFlags = IdentifyResponseFlags(0x0040)
    IdentifyResponseFlagsArmorProfile IdentifyResponseFlags = IdentifyResponseFlags(0x0080)
    IdentifyResponseFlagsCreatureProfile IdentifyResponseFlags = IdentifyResponseFlags(0x0100)
    IdentifyResponseFlagsArmorEnchantmentBitfield IdentifyResponseFlags = IdentifyResponseFlags(0x0200)
    IdentifyResponseFlagsResistEnchantmentBitfield IdentifyResponseFlags = IdentifyResponseFlags(0x0400)
    IdentifyResponseFlagsWeaponEnchantmentBitfield IdentifyResponseFlags = IdentifyResponseFlags(0x0800)
    IdentifyResponseFlagsDidStatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x1000)
    IdentifyResponseFlagsInt64StatsTable IdentifyResponseFlags = IdentifyResponseFlags(0x2000)
    IdentifyResponseFlagsArmorLevels IdentifyResponseFlags = IdentifyResponseFlags(0x4000)
    )
