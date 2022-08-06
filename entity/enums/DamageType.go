package enums

type DamageType int32

const (
	DamageTypeUndef     DamageType = DamageType(0x0)
	DamageTypeSlash     DamageType = DamageType(0x1)
	DamageTypePierce    DamageType = DamageType(0x2)
	DamageTypeBludgeon  DamageType = DamageType(0x4)
	DamageTypeCold      DamageType = DamageType(0x8)
	DamageTypeFire      DamageType = DamageType(0x10)
	DamageTypeAcid      DamageType = DamageType(0x20)
	DamageTypeElectric  DamageType = DamageType(0x40)
	DamageTypeHealth    DamageType = DamageType(0x80)
	DamageTypeStamina   DamageType = DamageType(0x100)
	DamageTypeMana      DamageType = DamageType(0x200)
	DamageTypeNether    DamageType = DamageType(0x400)
	DamageTypeBase      DamageType = DamageType(0x10000000)
	DamageTypePhysical  DamageType = DamageType(DamageTypeSlash | DamageTypePierce | DamageTypeBludgeon)
	DamageTypeElemental DamageType = DamageType(DamageTypeCold | DamageTypeFire | DamageTypeAcid | DamageTypeElectric)
)
