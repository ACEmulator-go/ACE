package enums


type MovementTypes int32

const (
    MovementTypesRawCommand MovementTypes = MovementTypes(0x1)
    MovementTypesInterpretedCommand MovementTypes = MovementTypes(0x2)
    MovementTypesStopRawCommand MovementTypes = MovementTypes(0x3)
    MovementTypesStopInterpretedCommand MovementTypes = MovementTypes(0x4)
    MovementTypesStopCompletely MovementTypes = MovementTypes(0x5)
    MovementTypesMoveToObject MovementTypes = MovementTypes(0x6)
    MovementTypesMoveToPosition MovementTypes = MovementTypes(0x7)
    MovementTypesTurnToObject MovementTypes = MovementTypes(0x8)
    MovementTypesTurnToHeading MovementTypes = MovementTypes(0x9)
    )
