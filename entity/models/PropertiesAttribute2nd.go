package models

type PropertiesAttribute2nd struct {
	InitLevel    uint
	LevelFromCP  uint
	CPSpent      uint
	CurrentLevel uint
}

func (prop *PropertiesAttribute2nd) Clone() PropertiesAttribute2nd {
	return PropertiesAttribute2nd{
		InitLevel:    prop.InitLevel,
		LevelFromCP:  prop.LevelFromCP,
		CPSpent:      prop.CPSpent,
		CurrentLevel: prop.CurrentLevel,
	}
}
