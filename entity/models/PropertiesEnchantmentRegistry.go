package models

import "github.com/ACEmulator-go/ACE/entity/enums"

type PropertiesEnchantmentRegistry struct {
	EnchantmentCategory uint
	SpellID             int
	LayerID             uint16
	HasSpellSetId       bool
	SpellCategory       enums.SpellCategory
	PowerLevel          uint
	StartTime           float64
	Duration            float64
	CasterObjectID      uint
	DegradeModifier     float32
	DegradeLimit        float32
	LastTimeDegraded    float64
	StatModType         enums.EnchantmentTypeFlags
	StatModKey          uint
	StatModValue        float32
	SpellSetId          enums.EquipmentSet
}
