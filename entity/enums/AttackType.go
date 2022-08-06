package enums

type AttackType int32

const (
	AttackTypeUndef               AttackType = AttackType(0x0)
	AttackTypePunch               AttackType = AttackType(0x0001)
	AttackTypeThrust              AttackType = AttackType(0x0002)
	AttackTypeSlash               AttackType = AttackType(0x0004)
	AttackTypeKick                AttackType = AttackType(0x0008)
	AttackTypeOffhandPunch        AttackType = AttackType(0x0010)
	AttackTypeDoubleSlash         AttackType = AttackType(0x0020)
	AttackTypeTripleSlash         AttackType = AttackType(0x0040)
	AttackTypeDoubleThrust        AttackType = AttackType(0x0080)
	AttackTypeTripleThrust        AttackType = AttackType(0x0100)
	AttackTypeOffhandThrust       AttackType = AttackType(0x0200)
	AttackTypeOffhandSlash        AttackType = AttackType(0x0400)
	AttackTypeOffhandDoubleSlash  AttackType = AttackType(0x0800)
	AttackTypeOffhandTripleSlash  AttackType = AttackType(0x1000)
	AttackTypeOffhandDoubleThrust AttackType = AttackType(0x2000)
	AttackTypeOffhandTripleThrust AttackType = AttackType(0x4000)
	AttackTypeUnarmed             AttackType = AttackType(AttackTypePunch | AttackTypeKick | AttackTypeOffhandPunch)
	AttackTypeDoubleStrike        AttackType = AttackType(AttackTypeDoubleSlash | AttackTypeDoubleThrust | AttackTypeOffhandDoubleSlash | AttackTypeOffhandDoubleThrust)
	AttackTypeTripleStrike        AttackType = AttackType(AttackTypeTripleSlash | AttackTypeTripleThrust | AttackTypeOffhandTripleSlash | AttackTypeOffhandTripleThrust)
	AttackTypeOffhand             AttackType = AttackType(AttackTypeOffhandThrust | AttackTypeOffhandSlash | AttackTypeOffhandDoubleSlash | AttackTypeOffhandTripleSlash | AttackTypeOffhandDoubleThrust | AttackTypeOffhandTripleThrust)
	AttackTypeThrusts             AttackType = AttackType(AttackTypeThrust | AttackTypeDoubleThrust | AttackTypeTripleThrust | AttackTypeOffhandThrust | AttackTypeOffhandDoubleThrust | AttackTypeOffhandTripleThrust)
	AttackTypeSlashes             AttackType = AttackType(AttackTypeSlash | AttackTypeDoubleSlash | AttackTypeTripleSlash | AttackTypeOffhandSlash | AttackTypeOffhandDoubleSlash | AttackTypeOffhandTripleSlash)
	AttackTypePunches             AttackType = AttackType(AttackTypePunch | AttackTypeOffhandPunch)
)
