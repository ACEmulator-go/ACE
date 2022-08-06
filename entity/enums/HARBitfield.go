package enums


type HARBitfield int32

const (
    HARBitfieldUndef HARBitfield = HARBitfield(0x0)
    HARBitfieldOpenHouse HARBitfield = HARBitfield(0x1)
    HARBitfieldAllegianceGuests HARBitfield = HARBitfield(0x2)
    )
