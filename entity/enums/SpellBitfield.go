package enums


type SpellBitfield int32

const (
    SpellBitfieldResistable SpellBitfield = SpellBitfield(0x1)
    SpellBitfieldPKSensitive SpellBitfield = SpellBitfield(0x2)
    SpellBitfieldBeneficial SpellBitfield = SpellBitfield(0x4)
    SpellBitfieldSelfTargeted SpellBitfield = SpellBitfield(0x8)
    SpellBitfieldReversed SpellBitfield = SpellBitfield(0x10)
    SpellBitfieldNotIndoor SpellBitfield = SpellBitfield(0x20)
    SpellBitfieldNotOutdoor SpellBitfield = SpellBitfield(0x40)
    SpellBitfieldNotResearchable SpellBitfield = SpellBitfield(0x80)
    SpellBitfieldProjectile SpellBitfield = SpellBitfield(0x100)
    SpellBitfieldCreatureSpell SpellBitfield = SpellBitfield(0x200)
    SpellBitfieldExcludedFromItemDescriptions SpellBitfield = SpellBitfield(0x400)
    SpellBitfieldIgnoresManaConversion SpellBitfield = SpellBitfield(0x800)
    SpellBitfieldNonTrackingProjectile SpellBitfield = SpellBitfield(0x1000)
    SpellBitfieldFellowshipSpell SpellBitfield = SpellBitfield(0x2000)
    SpellBitfieldFastCast SpellBitfield = SpellBitfield(0x4000)
    SpellBitfieldIndoorLongRange SpellBitfield = SpellBitfield(0x8000)
    SpellBitfieldDamageOverTime SpellBitfield = SpellBitfield(0x10000)
    )
