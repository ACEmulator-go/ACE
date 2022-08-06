package properties

type PropertyAttribute uint16

const (
	PropertyAttributeUndef        PropertyAttribute = PropertyAttribute(0)
	PropertyAttributeStrength     PropertyAttribute = PropertyAttribute(1)
	PropertyAttributeEndurance    PropertyAttribute = PropertyAttribute(2)
	PropertyAttributeQuickness    PropertyAttribute = PropertyAttribute(3)
	PropertyAttributeCoordination PropertyAttribute = PropertyAttribute(4)
	PropertyAttributeFocus        PropertyAttribute = PropertyAttribute(5)
)
