package enums

type WeaponMask int32

const (
	WeaponMaskAttackSkill    WeaponMask = WeaponMask(0x1)
	WeaponMaskMeleeDefense   WeaponMask = WeaponMask(0x2)
	WeaponMaskSpeed          WeaponMask = WeaponMask(0x4)
	WeaponMaskDamage         WeaponMask = WeaponMask(0x8)
	WeaponMaskDamageVariance WeaponMask = WeaponMask(0x10)
)
