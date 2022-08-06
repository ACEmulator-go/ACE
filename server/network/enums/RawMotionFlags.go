package enums

type RawMotionFlags uint32

const (
	RawMotionFlagsInvalid         RawMotionFlags = RawMotionFlags(0x0)
	RawMotionFlagsCurrentHoldKey  RawMotionFlags = RawMotionFlags(0x1)
	RawMotionFlagsCurrentStyle    RawMotionFlags = RawMotionFlags(0x2)
	RawMotionFlagsForwardCommand  RawMotionFlags = RawMotionFlags(0x4)
	RawMotionFlagsForwardHoldKey  RawMotionFlags = RawMotionFlags(0x8)
	RawMotionFlagsForwardSpeed    RawMotionFlags = RawMotionFlags(0x10)
	RawMotionFlagsSideStepCommand RawMotionFlags = RawMotionFlags(0x20)
	RawMotionFlagsSideStepHoldKey RawMotionFlags = RawMotionFlags(0x40)
	RawMotionFlagsSideStepSpeed   RawMotionFlags = RawMotionFlags(0x80)
	RawMotionFlagsTurnCommand     RawMotionFlags = RawMotionFlags(0x100)
	RawMotionFlagsTurnHoldKey     RawMotionFlags = RawMotionFlags(0x200)
)
