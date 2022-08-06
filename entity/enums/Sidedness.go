package enums


type Sidedness int32

const (
    SidednessPositive Sidedness = Sidedness(0x0)
    SidednessNegative Sidedness = Sidedness(0x1)
    SidednessInPlane Sidedness = Sidedness(0x2)
    )
