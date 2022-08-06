package enums

type PortalType uint32

const (
	PortalTypePurple PortalType = PortalType(0x020001B3)
	PortalTypeBlue   PortalType = PortalType(0x020005D2)
	PortalTypeGreen  PortalType = PortalType(0x020005D3)
	PortalTypeOrange PortalType = PortalType(0x020005D4)
	PortalTypeRed    PortalType = PortalType(0x020005D5)
	PortalTypeYellow PortalType = PortalType(0x020005D6)
	PortalTypeWhite  PortalType = PortalType(0x020006F4)
	PortalTypeShadow PortalType = PortalType(0x020008FD)
	PortalTypeBroken PortalType = PortalType(0x02000F2E)
)
