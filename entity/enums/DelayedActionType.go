package enums


type DelayedActionType int32

const (
    DelayedActionTypeStart DelayedActionType = DelayedActionType(0)
    DelayedActionTypeMove DelayedActionType = DelayedActionType(1)
    DelayedActionTypeMovePass DelayedActionType = DelayedActionType(2)
    DelayedActionTypeStalemate DelayedActionType = DelayedActionType(3)
    )
