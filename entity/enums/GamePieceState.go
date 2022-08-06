package enums


type GamePieceState int32

const (
    GamePieceStateNone GamePieceState = GamePieceState(0)
    GamePieceStateMoveToSquare GamePieceState = GamePieceState(1)
    GamePieceStateWaitingForMoveToSquare GamePieceState = GamePieceState(2)
    GamePieceStateWaitingForMoveToSquareAnimComplete GamePieceState = GamePieceState(3)
    GamePieceStateMoveToAttack GamePieceState = GamePieceState(4)
    GamePieceStateWaitingForMoveToAttack GamePieceState = GamePieceState(5)
    )
