package enums


type Vital uint32

const (
    VitalUndefined Vital = Vital(0)
    VitalMaxHealth Vital = Vital(1)
    VitalHealth Vital = Vital(2)
    VitalMaxStamina Vital = Vital(3)
    VitalStamina Vital = Vital(4)
    VitalMaxMana Vital = Vital(5)
    )
