package enums

type AllegianceHouseAction uint32

const (
	AllegianceHouseActionUndef       AllegianceHouseAction = AllegianceHouseAction(0)
	AllegianceHouseActionHelp        AllegianceHouseAction = AllegianceHouseAction(1)
	AllegianceHouseActionGuestOpen   AllegianceHouseAction = AllegianceHouseAction(2)
	AllegianceHouseActionGuestClose  AllegianceHouseAction = AllegianceHouseAction(3)
	AllegianceHouseActionStorageOpen AllegianceHouseAction = AllegianceHouseAction(4)
)
