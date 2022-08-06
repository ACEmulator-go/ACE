package enums


type ChessMoveFlag int32

const (
    ChessMoveFlagNone ChessMoveFlag = ChessMoveFlag(0x0)
    ChessMoveFlagNormal ChessMoveFlag = ChessMoveFlag(0x1)
    ChessMoveFlagCapture ChessMoveFlag = ChessMoveFlag(0x2)
    ChessMoveFlagBigPawn ChessMoveFlag = ChessMoveFlag(0x4)
    ChessMoveFlagEnPassantCapture ChessMoveFlag = ChessMoveFlag(0x8)
    ChessMoveFlagPromotion ChessMoveFlag = ChessMoveFlag(0x10)
    ChessMoveFlagKingSideCastle ChessMoveFlag = ChessMoveFlag(0x20)
    )
