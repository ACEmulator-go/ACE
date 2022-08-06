package enums


type EnchantmentMask int32

const (
    EnchantmentMaskMultiplicative EnchantmentMask = EnchantmentMask(0x1)
    EnchantmentMaskAdditive EnchantmentMask = EnchantmentMask(0x2)
    EnchantmentMaskVitae EnchantmentMask = EnchantmentMask(0x4)
    EnchantmentMaskCooldown EnchantmentMask = EnchantmentMask(0x8)
    )
