package enums


type MovementType byte

const (
    MovementTypeInvalid MovementType = MovementType(0x0)
    MovementTypeRawCommand MovementType = MovementType(0x1)
    MovementTypeInterpretedCommand MovementType = MovementType(0x2)
    MovementTypeStopRawCommand MovementType = MovementType(0x3)
    MovementTypeStopInterpretedCommand MovementType = MovementType(0x4)
    MovementTypeStopCompletely MovementType = MovementType(0x5)
    MovementTypeMoveToObject MovementType = MovementType(0x6)
    MovementTypeMoveToPosition MovementType = MovementType(0x7)
    MovementTypeTurnToObject MovementType = MovementType(0x8)
    MovementTypeTurnToHeading MovementType = MovementType(0x9)
    )
