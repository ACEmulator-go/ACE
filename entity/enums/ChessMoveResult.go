package enums


type ChessMoveResult int32

const (
    ChessMoveResultNoMoveResult ChessMoveResult = ChessMoveResult(0x0000)
    ChessMoveResultOKMoveToEmptySquare ChessMoveResult = ChessMoveResult(0x0001)
    ChessMoveResultOKMoveToOccupiedSquare ChessMoveResult = ChessMoveResult(0x0002)
    ChessMoveResultOKMoveCheck ChessMoveResult = ChessMoveResult(0x0400)
    ChessMoveResultOKMoveCheckmate ChessMoveResult = ChessMoveResult(0x0800)
    ChessMoveResultOKMovePromotion ChessMoveResult = ChessMoveResult(0x1000)
    ChessMoveResultBadMoveNotPlaying ChessMoveResult = ChessMoveResult(-2)
    )
