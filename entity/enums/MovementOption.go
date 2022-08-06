package enums


type MovementOption byte

const (
    MovementOptionNone MovementOption = MovementOption(0x0)
    MovementOptionStickToObject MovementOption = MovementOption(0x1)
    )
