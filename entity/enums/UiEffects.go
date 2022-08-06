package enums


type UiEffects uint32

const (
    UiEffectsUndef UiEffects = UiEffects(0x0000)
    UiEffectsMagical UiEffects = UiEffects(0x0001)
    UiEffectsPoisoned UiEffects = UiEffects(0x0002)
    UiEffectsBoostHealth UiEffects = UiEffects(0x0004)
    UiEffectsBoostMana UiEffects = UiEffects(0x0008)
    UiEffectsBoostStamina UiEffects = UiEffects(0x0010)
    UiEffectsFire UiEffects = UiEffects(0x0020)
    UiEffectsLightning UiEffects = UiEffects(0x0040)
    UiEffectsFrost UiEffects = UiEffects(0x0080)
    UiEffectsAcid UiEffects = UiEffects(0x0100)
    UiEffectsBludgeoning UiEffects = UiEffects(0x0200)
    UiEffectsSlashing UiEffects = UiEffects(0x0400)
    UiEffectsPiercing UiEffects = UiEffects(0x0800)
    )
