package enums

type ArmorMask int32

const (
	ArmorMaskArmorLevel            ArmorMask = ArmorMask(0x1)
	ArmorMaskSlashingProtection    ArmorMask = ArmorMask(0x2)
	ArmorMaskPiercingProtection    ArmorMask = ArmorMask(0x4)
	ArmorMaskBludgeoningProtection ArmorMask = ArmorMask(0x8)
	ArmorMaskColdProtection        ArmorMask = ArmorMask(0x10)
	ArmorMaskFireProtection        ArmorMask = ArmorMask(0x20)
	ArmorMaskAcidProtection        ArmorMask = ArmorMask(0x40)
)
