package enums


type GameEventState int32

const (
    GameEventStateUndef GameEventState = GameEventState(0x0)
    GameEventStateEnabled GameEventState = GameEventState(0x1)
    GameEventStateDisabled GameEventState = GameEventState(0x2)
    GameEventStateOff GameEventState = GameEventState(0x3)
    )
