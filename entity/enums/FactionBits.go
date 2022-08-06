package enums

type FactionBits int32

const (
	FactionBitsNone          FactionBits = FactionBits(0x0)
	FactionBitsCelestialHand FactionBits = FactionBits(0x1)
	FactionBitsEldrytchWeb   FactionBits = FactionBits(0x2)
	FactionBitsRadiantBlood  FactionBits = FactionBits(0x4)
)
