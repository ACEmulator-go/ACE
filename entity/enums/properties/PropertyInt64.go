package properties

type PropertyInt64 uint16

const (
	PropertyInt64Undef                 PropertyInt64 = PropertyInt64(0)
	PropertyInt64TotalExperience       PropertyInt64 = PropertyInt64(1)
	PropertyInt64AvailableExperience   PropertyInt64 = PropertyInt64(2)
	PropertyInt64AugmentationCost      PropertyInt64 = PropertyInt64(3)
	PropertyInt64ItemTotalXp           PropertyInt64 = PropertyInt64(4)
	PropertyInt64ItemBaseXp            PropertyInt64 = PropertyInt64(5)
	PropertyInt64AvailableLuminance    PropertyInt64 = PropertyInt64(6)
	PropertyInt64MaximumLuminance      PropertyInt64 = PropertyInt64(7)
	PropertyInt64InteractionReqs       PropertyInt64 = PropertyInt64(8)
	PropertyInt64AllegianceXPCached    PropertyInt64 = PropertyInt64(9000)
	PropertyInt64AllegianceXPGenerated PropertyInt64 = PropertyInt64(9001)
	PropertyInt64AllegianceXPReceived  PropertyInt64 = PropertyInt64(9002)
)
