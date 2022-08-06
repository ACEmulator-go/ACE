package enums


type ChessDelayedActionType int32

const (
    ChessDelayedActionTypeStart ChessDelayedActionType = ChessDelayedActionType(0)
    ChessDelayedActionTypeMove ChessDelayedActionType = ChessDelayedActionType(1)
    ChessDelayedActionTypeMovePass ChessDelayedActionType = ChessDelayedActionType(2)
    ChessDelayedActionTypeStalemate ChessDelayedActionType = ChessDelayedActionType(3)
    )
