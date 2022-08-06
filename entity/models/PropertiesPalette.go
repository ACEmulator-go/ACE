package models

type PropertiesPalette struct {
	SubPaletteID uint
	Offset       uint16
	Length       uint16
}

func (prop PropertiesPalette) Clone() PropertiesPalette {
	return PropertiesPalette{
		SubPaletteID: prop.SubPaletteID,
		Offset:       prop.Offset,
		Length:       prop.Length,
	}
}
