package enums


type EmitterType int32

const (
    EmitterTypeUnknown EmitterType = EmitterType(0)
    EmitterTypeBirthratePerSec EmitterType = EmitterType(1)
    EmitterTypeBirthratePerMeter EmitterType = EmitterType(2)
    )
