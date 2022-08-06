package enums


type EndTradeReason int32

const (
    EndTradeReasonNormal EndTradeReason = EndTradeReason(0x1)
    EndTradeReasonEnteredCombat EndTradeReason = EndTradeReason(0x2)
    )
