package enums


type MovementParams uint32

const (
    MovementParamsCanWalk MovementParams = MovementParams(0x1)
    MovementParamsCanRun MovementParams = MovementParams(0x2)
    MovementParamsCanSideStep MovementParams = MovementParams(0x4)
    MovementParamsCanWalkBackwards MovementParams = MovementParams(0x8)
    MovementParamsCanCharge MovementParams = MovementParams(0x10)
    MovementParamsFailWalk MovementParams = MovementParams(0x20)
    MovementParamsUseFinalHeading MovementParams = MovementParams(0x40)
    MovementParamsSticky MovementParams = MovementParams(0x80)
    MovementParamsMoveAway MovementParams = MovementParams(0x100)
    MovementParamsMoveTowards MovementParams = MovementParams(0x200)
    MovementParamsUseSpheres MovementParams = MovementParams(0x400)
    MovementParamsSetHoldKey MovementParams = MovementParams(0x800)
    MovementParamsAutonomous MovementParams = MovementParams(0x1000)
    MovementParamsModifyRawState MovementParams = MovementParams(0x2000)
    MovementParamsModifyInterpretedState MovementParams = MovementParams(0x4000)
    MovementParamsCancelMoveTo MovementParams = MovementParams(0x8000)
    MovementParamsStopCompletely MovementParams = MovementParams(0x10000)
    )
