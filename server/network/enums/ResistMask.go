package enums

type ResistMask int32

const (
	ResistMaskResistSlash        ResistMask = ResistMask(0x1)
	ResistMaskResistPierce       ResistMask = ResistMask(0x2)
	ResistMaskResistBludgeon     ResistMask = ResistMask(0x4)
	ResistMaskResistFire         ResistMask = ResistMask(0x8)
	ResistMaskResistCold         ResistMask = ResistMask(0x10)
	ResistMaskResistAcid         ResistMask = ResistMask(0x20)
	ResistMaskResistElectric     ResistMask = ResistMask(0x40)
	ResistMaskResistHealthBoost  ResistMask = ResistMask(0x80)
	ResistMaskResistStaminaDrain ResistMask = ResistMask(0x100)
	ResistMaskResistStaminaBoost ResistMask = ResistMask(0x200)
	ResistMaskResistManaDrain    ResistMask = ResistMask(0x400)
	ResistMaskResistManaBoost    ResistMask = ResistMask(0x800)
	ResistMaskManaConversionMod  ResistMask = ResistMask(0x1000)
	ResistMaskElementalDamageMod ResistMask = ResistMask(0x2000)
)
