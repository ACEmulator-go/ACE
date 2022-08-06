package enums


type EnchantmentTypeFlags int32

const (
    EnchantmentTypeFlagsUndef EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000000)
    EnchantmentTypeFlagsAttribute EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000001)
    EnchantmentTypeFlagsSecondAtt EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000002)
    EnchantmentTypeFlagsInt EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000004)
    EnchantmentTypeFlagsFloat EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000008)
    EnchantmentTypeFlagsSkill EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000010)
    EnchantmentTypeFlagsBodyDamageValue EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000020)
    EnchantmentTypeFlagsBodyDamageVariance EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000040)
    EnchantmentTypeFlagsBodyArmorValue EnchantmentTypeFlags = EnchantmentTypeFlags(0x0000080)
    EnchantmentTypeFlagsSingleStat EnchantmentTypeFlags = EnchantmentTypeFlags(0x0001000)
    EnchantmentTypeFlagsMultipleStat EnchantmentTypeFlags = EnchantmentTypeFlags(0x0002000)
    EnchantmentTypeFlagsMultiplicative EnchantmentTypeFlags = EnchantmentTypeFlags(0x0004000)
    EnchantmentTypeFlagsAdditive EnchantmentTypeFlags = EnchantmentTypeFlags(0x0008000)
    EnchantmentTypeFlagsAttackSkills EnchantmentTypeFlags = EnchantmentTypeFlags(0x0010000)
    EnchantmentTypeFlagsDefenseSkills EnchantmentTypeFlags = EnchantmentTypeFlags(0x0020000)
    EnchantmentTypeFlagsMultiplicative_Degrade EnchantmentTypeFlags = EnchantmentTypeFlags(0x0100000)
    EnchantmentTypeFlagsAdditive_Degrade EnchantmentTypeFlags = EnchantmentTypeFlags(0x0200000)
    EnchantmentTypeFlagsVitae EnchantmentTypeFlags = EnchantmentTypeFlags(0x0800000)
    EnchantmentTypeFlagsCooldown EnchantmentTypeFlags = EnchantmentTypeFlags(0x1000000)
    EnchantmentTypeFlagsBeneficial EnchantmentTypeFlags = EnchantmentTypeFlags(0x2000000)
    EnchantmentTypeFlagsStatTypes EnchantmentTypeFlags = EnchantmentTypeFlags(0x00000FF)
    )
