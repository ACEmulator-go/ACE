package enums


type ChessState int32

const (
    ChessStateWaitingForPlayers ChessState = ChessState(0)
    ChessStateInProgress ChessState = ChessState(1)
    )
