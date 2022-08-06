package enums

type AllegianceIndex int32

const (
	AllegianceIndexUndefined        AllegianceIndex = AllegianceIndex(0x0)
	AllegianceIndexLoggedIn         AllegianceIndex = AllegianceIndex(0x1)
	AllegianceIndexUpdate           AllegianceIndex = AllegianceIndex(0x2)
	AllegianceIndexHasAllegianceAge AllegianceIndex = AllegianceIndex(0x4)
	AllegianceIndexHasPackedLevel   AllegianceIndex = AllegianceIndex(0x8)
)
