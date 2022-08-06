package enums

type AttributeMask int32

const (
	AttributeMaskStrength     AttributeMask = AttributeMask(0x1)
	AttributeMaskEndurance    AttributeMask = AttributeMask(0x2)
	AttributeMaskQuickness    AttributeMask = AttributeMask(0x4)
	AttributeMaskCoordination AttributeMask = AttributeMask(0x8)
	AttributeMaskFocus        AttributeMask = AttributeMask(0x10)
	AttributeMaskSelf         AttributeMask = AttributeMask(0x20)
	AttributeMaskHealth       AttributeMask = AttributeMask(0x40)
	AttributeMaskStamina      AttributeMask = AttributeMask(0x80)
	AttributeMaskMana         AttributeMask = AttributeMask(0x100)
)
