package properties

type PropertyType int32

const (
	PropertyTypePropertyAttribute    PropertyType = PropertyType(0)
	PropertyTypePropertyAttribute2nd PropertyType = PropertyType(1)
	PropertyTypePropertyBook         PropertyType = PropertyType(2)
	PropertyTypePropertyBool         PropertyType = PropertyType(3)
	PropertyTypePropertyDataId       PropertyType = PropertyType(4)
	PropertyTypePropertyDouble       PropertyType = PropertyType(5)
	PropertyTypePropertyInstanceId   PropertyType = PropertyType(6)
	PropertyTypePropertyInt          PropertyType = PropertyType(7)
	PropertyTypePropertyInt64        PropertyType = PropertyType(8)
	PropertyTypePropertyString       PropertyType = PropertyType(9)
)
