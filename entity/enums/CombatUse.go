package enums


type CombatUse byte

const (
    CombatUseNone CombatUse = CombatUse(0x00)
    CombatUseMelee CombatUse = CombatUse(0x01)
    CombatUseMissile CombatUse = CombatUse(0x02)
    CombatUseAmmo CombatUse = CombatUse(0x03)
    CombatUseShield CombatUse = CombatUse(0x04)
    )
