package enums


type SetupFlags int32

const (
    SetupFlagsHasParent SetupFlags = SetupFlags(0x1)
    SetupFlagsHasDefaultScale SetupFlags = SetupFlags(0x2)
    SetupFlagsAllowFreeHeading SetupFlags = SetupFlags(0x4)
    )
