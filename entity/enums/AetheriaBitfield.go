package enums

type AetheriaBitfield int32

const (
	AetheriaBitfieldNone   AetheriaBitfield = AetheriaBitfield(0x0)
	AetheriaBitfieldBlue   AetheriaBitfield = AetheriaBitfield(0x1)
	AetheriaBitfieldYellow AetheriaBitfield = AetheriaBitfield(0x2)
	AetheriaBitfieldRed    AetheriaBitfield = AetheriaBitfield(0x4)
)
