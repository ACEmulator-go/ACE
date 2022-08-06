package enums

type PortalBitmask int32

const (
	PortalBitmaskUndef         PortalBitmask = PortalBitmask(0x00)
	PortalBitmaskUnrestricted  PortalBitmask = PortalBitmask(0x01)
	PortalBitmaskNoPk          PortalBitmask = PortalBitmask(0x02)
	PortalBitmaskNoPKLite      PortalBitmask = PortalBitmask(0x04)
	PortalBitmaskNoNPK         PortalBitmask = PortalBitmask(0x08)
	PortalBitmaskNoSummon      PortalBitmask = PortalBitmask(0x10)
	PortalBitmaskNoRecall      PortalBitmask = PortalBitmask(0x20)
	PortalBitmaskOnlyOlthoiPCs PortalBitmask = PortalBitmask(0x40)
	PortalBitmaskNoOlthoiPCs   PortalBitmask = PortalBitmask(0x80)
	PortalBitmaskNoVitae       PortalBitmask = PortalBitmask(0x100)
)
