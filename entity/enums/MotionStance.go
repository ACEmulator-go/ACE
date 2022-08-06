package enums


type MotionStance uint32

const (
    MotionStanceInvalid MotionStance = MotionStance(0x80000000)
    MotionStanceHandCombat MotionStance = MotionStance(0x8000003c)
    MotionStanceNonCombat MotionStance = MotionStance(0x8000003d)
    MotionStanceSwordCombat MotionStance = MotionStance(0x8000003e)
    MotionStanceBowCombat MotionStance = MotionStance(0x8000003f)
    MotionStanceSwordShieldCombat MotionStance = MotionStance(0x80000040)
    MotionStanceCrossbowCombat MotionStance = MotionStance(0x80000041)
    MotionStanceUnusedCombat MotionStance = MotionStance(0x80000042)
    MotionStanceSlingCombat MotionStance = MotionStance(0x80000043)
    MotionStanceDualWieldCombat MotionStance = MotionStance(0x80000046)
    MotionStanceThrownWeaponCombat MotionStance = MotionStance(0x80000047)
    MotionStanceGraze MotionStance = MotionStance(0x80000048)
    MotionStanceMagic MotionStance = MotionStance(0x80000049)
    MotionStanceBowNoAmmo MotionStance = MotionStance(0x800000e8)
    MotionStanceCrossBowNoAmmo MotionStance = MotionStance(0x800000e9)
    )
