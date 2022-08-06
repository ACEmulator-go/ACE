package models

import "github.com/ACEmulator-go/ACE/entity/enums"

type PropertiesSkill struct {
	LevelFromPP           uint16
	SAC                   enums.SkillAdvancementClass
	PP                    uint
	InitLevel             uint
	ResistanceAtLastCheck uint
	LastUsedTime          float64
}

func (prop PropertiesSkill) Clone() PropertiesSkill {
	return PropertiesSkill{
		LevelFromPP:           prop.LevelFromPP,
		SAC:                   prop.SAC,
		PP:                    prop.PP,
		InitLevel:             prop.InitLevel,
		ResistanceAtLastCheck: prop.ResistanceAtLastCheck,
		LastUsedTime:          prop.LastUsedTime,
	}
}
