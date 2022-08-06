package enums

type SimplePolygonType int32

const (
	SimplePolygonTypeSimplePolygon SimplePolygonType = SimplePolygonType(0x0)
	SimplePolygonTypePathPolygon   SimplePolygonType = SimplePolygonType(0x1)
)
