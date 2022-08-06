package models

type PropertiesTextureMap struct {
	PartIndex  byte
	OldTexture uint
	NewTexture uint
}

func (prop PropertiesTextureMap) Clone() PropertiesTextureMap {
	return PropertiesTextureMap{
		PartIndex:  prop.PartIndex,
		OldTexture: prop.OldTexture,
		NewTexture: prop.NewTexture,
	}
}
