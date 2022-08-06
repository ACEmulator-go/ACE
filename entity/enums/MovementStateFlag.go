package enums


type MovementStateFlag uint32

const (
    MovementStateFlagInvalid MovementStateFlag = MovementStateFlag(0x0)
    MovementStateFlagCurrentStyle MovementStateFlag = MovementStateFlag(0x1)
    MovementStateFlagForwardCommand MovementStateFlag = MovementStateFlag(0x2)
    MovementStateFlagForwardSpeed MovementStateFlag = MovementStateFlag(0x4)
    MovementStateFlagSideStepCommand MovementStateFlag = MovementStateFlag(0x8)
    MovementStateFlagSideStepSpeed MovementStateFlag = MovementStateFlag(0x10)
    MovementStateFlagTurnCommand MovementStateFlag = MovementStateFlag(0x20)
    )
