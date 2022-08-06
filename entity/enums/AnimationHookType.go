package enums


type AnimationHookType int32

const (
    AnimationHookTypeUnknown AnimationHookType = AnimationHookType(-1)
    AnimationHookTypeNoOp AnimationHookType = AnimationHookType(0)
    AnimationHookTypeSound AnimationHookType = AnimationHookType(1)
    AnimationHookTypeSoundTable AnimationHookType = AnimationHookType(2)
    AnimationHookTypeAttack AnimationHookType = AnimationHookType(3)
    AnimationHookTypeAnimationDone AnimationHookType = AnimationHookType(4)
    AnimationHookTypeReplaceObject AnimationHookType = AnimationHookType(5)
    AnimationHookTypeEthereal AnimationHookType = AnimationHookType(6)
    AnimationHookTypeTransparentPart AnimationHookType = AnimationHookType(7)
    AnimationHookTypeLuminous AnimationHookType = AnimationHookType(8)
    AnimationHookTypeLuminousPart AnimationHookType = AnimationHookType(9)
    AnimationHookTypeDiffuse AnimationHookType = AnimationHookType(10)
    AnimationHookTypeDiffusePart AnimationHookType = AnimationHookType(11)
    AnimationHookTypeScale AnimationHookType = AnimationHookType(12)
    AnimationHookTypeCreateParticle AnimationHookType = AnimationHookType(13)
    AnimationHookTypeDestroyParticle AnimationHookType = AnimationHookType(14)
    AnimationHookTypeStopParticle AnimationHookType = AnimationHookType(15)
    AnimationHookTypeNoDraw AnimationHookType = AnimationHookType(16)
    AnimationHookTypeDefaultScript AnimationHookType = AnimationHookType(17)
    AnimationHookTypeDefaultScriptPart AnimationHookType = AnimationHookType(18)
    AnimationHookTypeTransparent AnimationHookType = AnimationHookType(20)
    AnimationHookTypeSoundTweaked AnimationHookType = AnimationHookType(21)
    AnimationHookTypeSetOmega AnimationHookType = AnimationHookType(22)
    AnimationHookTypeTextureVelocity AnimationHookType = AnimationHookType(23)
    AnimationHookTypeTextureVelocityPart AnimationHookType = AnimationHookType(24)
    AnimationHookTypeSetLight AnimationHookType = AnimationHookType(25)
    AnimationHookTypeCreateBlockingParticle AnimationHookType = AnimationHookType(26)
    )
