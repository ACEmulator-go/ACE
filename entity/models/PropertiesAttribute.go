package models

type PropertiesAttribute struct {
	InitLevel   uint
	LevelFromCP uint
	CPSpent     uint
}

func (prop *PropertiesAttribute) Clone() PropertiesAttribute {
	return PropertiesAttribute{
		InitLevel:   prop.InitLevel,
		LevelFromCP: prop.LevelFromCP,
		CPSpent:     prop.CPSpent,
	}
}
