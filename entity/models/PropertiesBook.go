package models

type PropertiesBook struct {
	MaxNumPages        int
	MaxNumCharsPerPage int
}

func (prop PropertiesBook) Clone() PropertiesBook {
	return PropertiesBook{
		MaxNumPages:        prop.MaxNumPages,
		MaxNumCharsPerPage: prop.MaxNumCharsPerPage,
	}
}
