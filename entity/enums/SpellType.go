package enums


type SpellType int32

const (
    SpellTypeUndef SpellType = SpellType(0)
    SpellTypeEnchantment SpellType = SpellType(1)
    SpellTypeProjectile SpellType = SpellType(2)
    SpellTypeBoost SpellType = SpellType(3)
    SpellTypeTransfer SpellType = SpellType(4)
    SpellTypePortalLink SpellType = SpellType(5)
    SpellTypePortalRecall SpellType = SpellType(6)
    SpellTypePortalSummon SpellType = SpellType(7)
    SpellTypePortalSending SpellType = SpellType(8)
    SpellTypeDispel SpellType = SpellType(9)
    SpellTypeLifeProjectile SpellType = SpellType(10)
    SpellTypeFellowBoost SpellType = SpellType(11)
    SpellTypeFellowEnchantment SpellType = SpellType(12)
    SpellTypeFellowPortalSending SpellType = SpellType(13)
    SpellTypeFellowDispel SpellType = SpellType(14)
    )
