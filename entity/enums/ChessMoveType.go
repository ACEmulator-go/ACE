package enums


type ChessMoveType int32

const (
    ChessMoveTypeInvalid ChessMoveType = ChessMoveType(0)
    ChessMoveTypePass ChessMoveType = ChessMoveType(1)
    ChessMoveTypeResign ChessMoveType = ChessMoveType(2)
    ChessMoveTypeStalemate ChessMoveType = ChessMoveType(3)
    ChessMoveTypeGrid ChessMoveType = ChessMoveType(4)
    ChessMoveTypeFromTo ChessMoveType = ChessMoveType(5)
    )
