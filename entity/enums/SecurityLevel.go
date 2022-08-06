package enums

type SecurityLevel int32

const (
	SecurityLevelUndef       SecurityLevel = SecurityLevel(0)
	SecurityLevelPlayer      SecurityLevel = SecurityLevel(SecurityLevelUndef)
	SecurityLevelAdvocate1   SecurityLevel = SecurityLevel(1)
	SecurityLevelAdvocate2   SecurityLevel = SecurityLevel(2)
	SecurityLevelAdvocate3   SecurityLevel = SecurityLevel(3)
	SecurityLevelAdvocate4   SecurityLevel = SecurityLevel(4)
	SecurityLevelAdvocate5   SecurityLevel = SecurityLevel(5)
	SecurityLevelMaxAdvocate SecurityLevel = SecurityLevel(SecurityLevelAdvocate5)
	SecurityLevelSentinel1   SecurityLevel = SecurityLevel(6)
	SecurityLevelSentinel2   SecurityLevel = SecurityLevel(7)
	SecurityLevelSentinel3   SecurityLevel = SecurityLevel(8)
	SecurityLevelMaxSentinel SecurityLevel = SecurityLevel(SecurityLevelSentinel3)
	SecurityLevelTurbine     SecurityLevel = SecurityLevel(9)
	SecurityLevelArch        SecurityLevel = SecurityLevel(10)
	SecurityLevelAdmin       SecurityLevel = SecurityLevel(11)
)
