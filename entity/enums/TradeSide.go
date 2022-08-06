package enums


type TradeSide int32

const (
    TradeSideSelf TradeSide = TradeSide(0x1)
    TradeSidePartner TradeSide = TradeSide(0x2)
    )
