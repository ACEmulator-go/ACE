package enums

type AttributeCache uint32

const (
	AttributeCacheUndef        AttributeCache = AttributeCache(0x00000000)
	AttributeCacheStrength     AttributeCache = AttributeCache(0x00000001)
	AttributeCacheEndurance    AttributeCache = AttributeCache(0x00000002)
	AttributeCacheQuickness    AttributeCache = AttributeCache(0x00000004)
	AttributeCacheCoordination AttributeCache = AttributeCache(0x00000008)
	AttributeCacheFocus        AttributeCache = AttributeCache(0x00000010)
	AttributeCacheSelf         AttributeCache = AttributeCache(0x00000020)
	AttributeCacheHealth       AttributeCache = AttributeCache(0x00000040)
	AttributeCacheStamina      AttributeCache = AttributeCache(0x00000080)
	AttributeCacheMana         AttributeCache = AttributeCache(0x00000100)
)
