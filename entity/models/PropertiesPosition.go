package models

type PropertiesPosition struct {
	ObjCellID uint
	PositionX float32
	PositionY float32
	PositionZ float32
	RotationW float32
	RotationX float32
	RotationY float32
	RotationZ float32
}

func (prop PropertiesPosition) Clone() PropertiesPosition {
	return PropertiesPosition{
		ObjCellID: prop.ObjCellID,
		PositionX: prop.PositionX,
		PositionY: prop.PositionY,
		PositionZ: prop.PositionZ,
		RotationW: prop.RotationW,
		RotationX: prop.RotationX,
		RotationY: prop.RotationY,
		RotationZ: prop.RotationZ,
	}
}
