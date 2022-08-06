package enums

type AttackConditions int32

const (
	AttackConditionsNone                           AttackConditions = AttackConditions(0x0)
	AttackConditionsCriticalProtectionAugmentation AttackConditions = AttackConditions(0x1)
	AttackConditionsRecklessness                   AttackConditions = AttackConditions(0x2)
	AttackConditionsSneakAttack                    AttackConditions = AttackConditions(0x4)
)
