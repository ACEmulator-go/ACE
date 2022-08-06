package models

type PropertiesAnimPart struct {
	Index       byte
	AnimationID uint
}

func (prop *PropertiesAnimPart) Clone() PropertiesAnimPart {
	return PropertiesAnimPart{
		Index:       prop.Index,
		AnimationID: prop.AnimationID,
	}
}
