package enums

type CombatStyle int32

const (
	CombatStyleUndef              CombatStyle = CombatStyle(0x00000)
	CombatStyleUnarmed            CombatStyle = CombatStyle(0x00001)
	CombatStyleOneHanded          CombatStyle = CombatStyle(0x00002)
	CombatStyleOneHandedAndShield CombatStyle = CombatStyle(0x00004)
	CombatStyleTwoHanded          CombatStyle = CombatStyle(0x00008)
	CombatStyleBow                CombatStyle = CombatStyle(0x00010)
	CombatStyleCrossbow           CombatStyle = CombatStyle(0x00020)
	CombatStyleSling              CombatStyle = CombatStyle(0x00040)
	CombatStyleThrownWeapon       CombatStyle = CombatStyle(0x00080)
	CombatStyleDualWield          CombatStyle = CombatStyle(0x00100)
	CombatStyleMagic              CombatStyle = CombatStyle(0x00200)
	CombatStyleAtlatl             CombatStyle = CombatStyle(0x00400)
	CombatStyleThrownShield       CombatStyle = CombatStyle(0x00800)
	CombatStyleReserved1          CombatStyle = CombatStyle(0x01000)
	CombatStyleReserved2          CombatStyle = CombatStyle(0x02000)
	CombatStyleReserved3          CombatStyle = CombatStyle(0x04000)
	CombatStyleReserved4          CombatStyle = CombatStyle(0x08000)
	CombatStyleStubbornMagic      CombatStyle = CombatStyle(0x10000)
	CombatStyleStubbornProjectile CombatStyle = CombatStyle(0x20000)
	CombatStyleStubbornMelee      CombatStyle = CombatStyle(0x40000)
	CombatStyleStubbornMissile    CombatStyle = CombatStyle(0x80000)
	CombatStyleMelee              CombatStyle = CombatStyle(CombatStyleUnarmed | CombatStyleOneHanded | CombatStyleOneHandedAndShield | CombatStyleTwoHanded | CombatStyleDualWield)
	CombatStyleMissile            CombatStyle = CombatStyle(CombatStyleBow | CombatStyleCrossbow | CombatStyleSling | CombatStyleThrownWeapon | CombatStyleAtlatl | CombatStyleThrownShield)
)
