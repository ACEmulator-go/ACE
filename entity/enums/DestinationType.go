package enums

type DestinationType int32

const (
	DestinationTypeUndef           DestinationType = DestinationType(0)
	DestinationTypeContain         DestinationType = DestinationType(1)
	DestinationTypeWield           DestinationType = DestinationType(2)
	DestinationTypeShop            DestinationType = DestinationType(4)
	DestinationTypeTreasure        DestinationType = DestinationType(8)
	DestinationTypeHouseBuy        DestinationType = DestinationType(16)
	DestinationTypeHouseRent       DestinationType = DestinationType(32)
	DestinationTypeCheckpoint      DestinationType = DestinationType(DestinationTypeContain | DestinationTypeWield | DestinationTypeShop)
	DestinationTypeContainTreasure DestinationType = DestinationType(DestinationTypeContain | DestinationTypeTreasure)
	DestinationTypeWieldTreasure   DestinationType = DestinationType(DestinationTypeWield | DestinationTypeTreasure)
)
