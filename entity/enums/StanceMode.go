package enums


type StanceMode uint16

const (
    StanceModeInvalid StanceMode = StanceMode(0x0)
    StanceModeHandCombat StanceMode = StanceMode(0x3c)
    StanceModeNonCombat StanceMode = StanceMode(0x3d)
    StanceModeSwordCombat StanceMode = StanceMode(0x3e)
    StanceModeBowCombat StanceMode = StanceMode(0x3f)
    StanceModeSwordShieldCombat StanceMode = StanceMode(0x40)
    StanceModeCrossbowCombat StanceMode = StanceMode(0x41)
    StanceModeUnusedCombat StanceMode = StanceMode(0x42)
    StanceModeSlingCombat StanceMode = StanceMode(0x43)
    StanceModeDualWieldCombat StanceMode = StanceMode(0x46)
    StanceModeThrownWeaponCombat StanceMode = StanceMode(0x47)
    StanceModeMagic StanceMode = StanceMode(0x49)
    StanceModeBowNoAmmo StanceMode = StanceMode(0xe8)
    StanceModeCrossBowNoAmmo StanceMode = StanceMode(0xe9)
    )
