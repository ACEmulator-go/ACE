package enums


type WeenieHeaderFlag uint32

const (
    WeenieHeaderFlagNone WeenieHeaderFlag = WeenieHeaderFlag(0x00000000)
    WeenieHeaderFlagPluralName WeenieHeaderFlag = WeenieHeaderFlag(0x00000001)
    WeenieHeaderFlagItemsCapacity WeenieHeaderFlag = WeenieHeaderFlag(0x00000002)
    WeenieHeaderFlagContainersCapacity WeenieHeaderFlag = WeenieHeaderFlag(0x00000004)
    WeenieHeaderFlagValue WeenieHeaderFlag = WeenieHeaderFlag(0x00000008)
    WeenieHeaderFlagUseRadius WeenieHeaderFlag = WeenieHeaderFlag(0x00000020)
    WeenieHeaderFlagMonarch WeenieHeaderFlag = WeenieHeaderFlag(0x00000040)
    WeenieHeaderFlagUiEffects WeenieHeaderFlag = WeenieHeaderFlag(0x00000080)
    WeenieHeaderFlagAmmoType WeenieHeaderFlag = WeenieHeaderFlag(0x00000100)
    WeenieHeaderFlagCombatUse WeenieHeaderFlag = WeenieHeaderFlag(0x00000200)
    WeenieHeaderFlagStructure WeenieHeaderFlag = WeenieHeaderFlag(0x00000400)
    WeenieHeaderFlagMaxStructure WeenieHeaderFlag = WeenieHeaderFlag(0x00000800)
    WeenieHeaderFlagStackSize WeenieHeaderFlag = WeenieHeaderFlag(0x00001000)
    WeenieHeaderFlagMaxStackSize WeenieHeaderFlag = WeenieHeaderFlag(0x00002000)
    WeenieHeaderFlagContainer WeenieHeaderFlag = WeenieHeaderFlag(0x00004000)
    WeenieHeaderFlagWielder WeenieHeaderFlag = WeenieHeaderFlag(0x00008000)
    WeenieHeaderFlagValidLocations WeenieHeaderFlag = WeenieHeaderFlag(0x00010000)
    WeenieHeaderFlagPriority WeenieHeaderFlag = WeenieHeaderFlag(0x00040000)
    WeenieHeaderFlagTargetType WeenieHeaderFlag = WeenieHeaderFlag(0x00080000)
    WeenieHeaderFlagRadarBlipColor WeenieHeaderFlag = WeenieHeaderFlag(0x00100000)
    WeenieHeaderFlagBurden WeenieHeaderFlag = WeenieHeaderFlag(0x00200000)
    WeenieHeaderFlagSpell WeenieHeaderFlag = WeenieHeaderFlag(0x00400000)
    WeenieHeaderFlagRadarBehavior WeenieHeaderFlag = WeenieHeaderFlag(0x00800000)
    WeenieHeaderFlagWorkmanship WeenieHeaderFlag = WeenieHeaderFlag(0x01000000)
    WeenieHeaderFlagHouseOwner WeenieHeaderFlag = WeenieHeaderFlag(0x02000000)
    WeenieHeaderFlagHouseRestrictions WeenieHeaderFlag = WeenieHeaderFlag(0x04000000)
    WeenieHeaderFlagPScript WeenieHeaderFlag = WeenieHeaderFlag(0x08000000)
    WeenieHeaderFlagHookType WeenieHeaderFlag = WeenieHeaderFlag(0x10000000)
    WeenieHeaderFlagHookItemTypes WeenieHeaderFlag = WeenieHeaderFlag(0x20000000)
    WeenieHeaderFlagIconOverlay WeenieHeaderFlag = WeenieHeaderFlag(0x40000000)
    )

type WeenieHeaderFlag2 uint32

const (
    WeenieHeaderFlag2None WeenieHeaderFlag2 = WeenieHeaderFlag2(0x00)
    WeenieHeaderFlag2IconUnderlay WeenieHeaderFlag2 = WeenieHeaderFlag2(0x01)
    WeenieHeaderFlag2Cooldown WeenieHeaderFlag2 = WeenieHeaderFlag2(0x02)
    WeenieHeaderFlag2CooldownDuration WeenieHeaderFlag2 = WeenieHeaderFlag2(0x04)
    )
