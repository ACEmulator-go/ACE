package enums

type Quadrant int32

const (
	QuadrantNone   Quadrant = Quadrant(0x0)
	QuadrantHigh   Quadrant = Quadrant(0x1)
	QuadrantMedium Quadrant = Quadrant(0x2)
	QuadrantLow    Quadrant = Quadrant(0x4)
	QuadrantLeft   Quadrant = Quadrant(0x8)
	QuadrantRight  Quadrant = Quadrant(0x10)
	QuadrantFront  Quadrant = Quadrant(0x20)
)
