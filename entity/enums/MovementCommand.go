package enums


type MovementCommand uint16

const (
    MovementCommandInvalid MovementCommand = MovementCommand(0x0)
    MovementCommandHoldRun MovementCommand = MovementCommand(0x1)
    MovementCommandHoldSideStep MovementCommand = MovementCommand(0x2)
    MovementCommandReady MovementCommand = MovementCommand(0x3)
    MovementCommandWalkForward MovementCommand = MovementCommand(0x5)
    MovementCommandWalkBackwards MovementCommand = MovementCommand(0x6)
    MovementCommandRunForward MovementCommand = MovementCommand(0x7)
    MovementCommandTurnRight MovementCommand = MovementCommand(0x0D)
    MovementCommandTurnLeft MovementCommand = MovementCommand(0x0E)
    MovementCommandSideStepRight MovementCommand = MovementCommand(0x0F)
    )
