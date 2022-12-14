package enums


type ObjectDescriptionFlag int32

const (
    ObjectDescriptionFlagNone ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000000)
    ObjectDescriptionFlagOpenable ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000001)
    ObjectDescriptionFlagInscribable ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000002)
    ObjectDescriptionFlagStuck ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000004)
    ObjectDescriptionFlagPlayer ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000008)
    ObjectDescriptionFlagAttackable ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000010)
    ObjectDescriptionFlagPlayerKiller ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000020)
    ObjectDescriptionFlagHiddenAdmin ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000040)
    ObjectDescriptionFlagUiHidden ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000080)
    ObjectDescriptionFlagBook ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000100)
    ObjectDescriptionFlagVendor ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000200)
    ObjectDescriptionFlagPkSwitch ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000400)
    ObjectDescriptionFlagNpkSwitch ObjectDescriptionFlag = ObjectDescriptionFlag(0x00000800)
    ObjectDescriptionFlagDoor ObjectDescriptionFlag = ObjectDescriptionFlag(0x00001000)
    ObjectDescriptionFlagCorpse ObjectDescriptionFlag = ObjectDescriptionFlag(0x00002000)
    ObjectDescriptionFlagLifeStone ObjectDescriptionFlag = ObjectDescriptionFlag(0x00004000)
    ObjectDescriptionFlagFood ObjectDescriptionFlag = ObjectDescriptionFlag(0x00008000)
    ObjectDescriptionFlagHealer ObjectDescriptionFlag = ObjectDescriptionFlag(0x00010000)
    ObjectDescriptionFlagLockpick ObjectDescriptionFlag = ObjectDescriptionFlag(0x00020000)
    ObjectDescriptionFlagPortal ObjectDescriptionFlag = ObjectDescriptionFlag(0x00040000)
    ObjectDescriptionFlagAdmin ObjectDescriptionFlag = ObjectDescriptionFlag(0x00100000)
    ObjectDescriptionFlagFreePkStatus ObjectDescriptionFlag = ObjectDescriptionFlag(0x00200000)
    ObjectDescriptionFlagImmuneCellRestrictions ObjectDescriptionFlag = ObjectDescriptionFlag(0x00400000)
    ObjectDescriptionFlagRequiresPackSlot ObjectDescriptionFlag = ObjectDescriptionFlag(0x00800000)
    ObjectDescriptionFlagRetained ObjectDescriptionFlag = ObjectDescriptionFlag(0x01000000)
    ObjectDescriptionFlagPkLiteStatus ObjectDescriptionFlag = ObjectDescriptionFlag(0x02000000)
    ObjectDescriptionFlagIncludesSecondHeader ObjectDescriptionFlag = ObjectDescriptionFlag(0x04000000)
    ObjectDescriptionFlagBindStone ObjectDescriptionFlag = ObjectDescriptionFlag(0x08000000)
    ObjectDescriptionFlagVolatileRare ObjectDescriptionFlag = ObjectDescriptionFlag(0x10000000)
    ObjectDescriptionFlagWieldOnUse ObjectDescriptionFlag = ObjectDescriptionFlag(0x20000000)
    ObjectDescriptionFlagWieldLeft ObjectDescriptionFlag = ObjectDescriptionFlag(0x40000000)
    )
