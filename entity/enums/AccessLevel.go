package enums

type AccessLevel int32

const (
	AccessLevelPlayer    AccessLevel = AccessLevel(0)
	AccessLevelAdvocate  AccessLevel = AccessLevel(1)
	AccessLevelSentinel  AccessLevel = AccessLevel(2)
	AccessLevelEnvoy     AccessLevel = AccessLevel(3)
	AccessLevelDeveloper AccessLevel = AccessLevel(4)
)
