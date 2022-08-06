package enums

type AllegianceLockAction uint32

const (
	AllegianceLockActionUndef         AllegianceLockAction = AllegianceLockAction(0)
	AllegianceLockActionOff           AllegianceLockAction = AllegianceLockAction(1)
	AllegianceLockActionOn            AllegianceLockAction = AllegianceLockAction(2)
	AllegianceLockActionToggle        AllegianceLockAction = AllegianceLockAction(3)
	AllegianceLockActionCheck         AllegianceLockAction = AllegianceLockAction(4)
	AllegianceLockActionCheckApproved AllegianceLockAction = AllegianceLockAction(5)
)
