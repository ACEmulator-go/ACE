package enums


type MotionFlags byte

const (
    MotionFlagsNone MotionFlags = MotionFlags(0x0)
    MotionFlagsStickToObject MotionFlags = MotionFlags(0x1)
    )
