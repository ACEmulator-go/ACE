package enums

type CombatMode int32

const (
	CombatModeUndef       CombatMode = CombatMode(0x00)
	CombatModeNonCombat   CombatMode = CombatMode(0x01)
	CombatModeMelee       CombatMode = CombatMode(0x02)
	CombatModeMissile     CombatMode = CombatMode(0x04)
	CombatModeMagic       CombatMode = CombatMode(0x08)
	CombatModeValidCombat CombatMode = CombatMode(CombatModeNonCombat | CombatModeMelee | CombatModeMissile | CombatModeMagic)
)
