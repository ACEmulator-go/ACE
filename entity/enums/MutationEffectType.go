package enums


type MutationEffectType int32

const (
    MutationEffectTypeAssign MutationEffectType = MutationEffectType(0)
    MutationEffectTypeAdd MutationEffectType = MutationEffectType(1)
    MutationEffectTypeSubtract MutationEffectType = MutationEffectType(2)
    MutationEffectTypeMultiply MutationEffectType = MutationEffectType(3)
    MutationEffectTypeDivide MutationEffectType = MutationEffectType(4)
    MutationEffectTypeAtLeastAdd MutationEffectType = MutationEffectType(5)
    MutationEffectTypeAtMostSubtract MutationEffectType = MutationEffectType(6)
    MutationEffectTypeAddMultiply MutationEffectType = MutationEffectType(7)
    MutationEffectTypeAddDivide MutationEffectType = MutationEffectType(8)
    MutationEffectTypeSubtractMultiply MutationEffectType = MutationEffectType(9)
    MutationEffectTypeSubtractDivide MutationEffectType = MutationEffectType(10)
    MutationEffectTypeAssignAdd MutationEffectType = MutationEffectType(11)
    MutationEffectTypeAssignSubtract MutationEffectType = MutationEffectType(12)
    MutationEffectTypeAssignMultiply MutationEffectType = MutationEffectType(13)
    )
