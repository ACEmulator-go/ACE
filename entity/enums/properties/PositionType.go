package properties

type PositionType uint16

const (
	PositionTypeUndef               PositionType = PositionType(0)
	PositionTypeLocation            PositionType = PositionType(1)
	PositionTypeDestination         PositionType = PositionType(2)
	PositionTypeInstantiation       PositionType = PositionType(3)
	PositionTypeSanctuary           PositionType = PositionType(4)
	PositionTypeHome                PositionType = PositionType(5)
	PositionTypeActivationMove      PositionType = PositionType(6)
	PositionTypeTarget              PositionType = PositionType(7)
	PositionTypeLinkedPortalOne     PositionType = PositionType(8)
	PositionTypeLastPortal          PositionType = PositionType(9)
	PositionTypePortalStorm         PositionType = PositionType(10)
	PositionTypeCrashAndTurn        PositionType = PositionType(11)
	PositionTypePortalSummonLoc     PositionType = PositionType(12)
	PositionTypeHouseBoot           PositionType = PositionType(13)
	PositionTypeLinkedLifestone     PositionType = PositionType(15)
	PositionTypeLinkedPortalTwo     PositionType = PositionType(16)
	PositionTypeRelativeDestination PositionType = PositionType(26)
	PositionTypeTeleportedCharacter PositionType = PositionType(27)
)
