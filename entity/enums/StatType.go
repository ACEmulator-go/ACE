package enums

type StatType int32

const (
	StatTypeUndef              StatType = StatType(0)
	StatTypeInt                StatType = StatType(1)
	StatTypeFloat              StatType = StatType(2)
	StatTypePosition           StatType = StatType(3)
	StatTypeSkill              StatType = StatType(4)
	StatTypeString             StatType = StatType(5)
	StatTypeDataID             StatType = StatType(6)
	StatTypeInstanceID         StatType = StatType(7)
	StatTypeDID                StatType = StatType(6)
	StatTypeIID                StatType = StatType(7)
	StatTypeAttribute          StatType = StatType(8)
	StatTypeAttribute2nd       StatType = StatType(9)
	StatTypeBodyDamageValue    StatType = StatType(10)
	StatTypeBodyDamageVariance StatType = StatType(11)
	StatTypeBodyArmorValue     StatType = StatType(12)
	StatTypeBool               StatType = StatType(13)
	StatTypeInt64              StatType = StatType(14)
)
