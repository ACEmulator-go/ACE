package enums


type StipplingType int32

const (
    StipplingTypeNone StipplingType = StipplingType(0x0)
    StipplingTypePositive StipplingType = StipplingType(0x1)
    StipplingTypeNegative StipplingType = StipplingType(0x2)
    StipplingTypeBoth StipplingType = StipplingType(0x3)
    StipplingTypeNoPos StipplingType = StipplingType(0x4)
    StipplingTypeNoNeg StipplingType = StipplingType(0x8)
    )
