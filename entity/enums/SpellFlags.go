package enums


type SpellFlags int32

const (
    SpellFlagsResistable SpellFlags = SpellFlags(0x1)
    SpellFlagsPKSensitive SpellFlags = SpellFlags(0x2)
    SpellFlagsBeneficial SpellFlags = SpellFlags(0x4)
    SpellFlagsSelfTargeted SpellFlags = SpellFlags(0x8)
    SpellFlagsReversed SpellFlags = SpellFlags(0x10)
    SpellFlagsNotIndoor SpellFlags = SpellFlags(0x20)
    SpellFlagsNotOutdoor SpellFlags = SpellFlags(0x40)
    SpellFlagsNotResearchable SpellFlags = SpellFlags(0x80)
    SpellFlagsProjectile SpellFlags = SpellFlags(0x100)
    SpellFlagsCreatureSpell SpellFlags = SpellFlags(0x200)
    SpellFlagsExcludedFromItemDescriptions SpellFlags = SpellFlags(0x400)
    SpellFlagsIgnoresManaConversion SpellFlags = SpellFlags(0x800)
    SpellFlagsNonTrackingProjectile SpellFlags = SpellFlags(0x1000)
    SpellFlagsFellowshipSpell SpellFlags = SpellFlags(0x2000)
    SpellFlagsFastCast SpellFlags = SpellFlags(0x4000)
    SpellFlagsIndoorLongRange SpellFlags = SpellFlags(0x8000)
    SpellFlagsDamageOverTime SpellFlags = SpellFlags(0x10000)
    )
