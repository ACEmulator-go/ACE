package models

import (
	"github.com/ACEmulator-go/ACE/entity/enums"
)

type PropertiesBodyPart struct {
	DType           enums.DamageType
	DVal            int
	DVar            float32
	BaseArmor       int
	ArmorVsSlash    int
	ArmorVsPierce   int
	ArmorVsBludgeon int
	ArmorVsCold     int
	ArmorVsFire     int
	ArmorVsAcid     int
	ArmorVsElectric int
	ArmorVsNether   int
	BH              int
	HLF             int
	MLF             int
	LLF             int
	HRF             int
	MRF             int
	LRF             int
	HLB             int
	MLB             int
	LLB             int
	HRB             int
	MRB             int
	LRB             int
}

func (prop PropertiesBodyPart) Clone() PropertiesBodyPart {
	return PropertiesBodyPart{
		DType:           prop.DType,
		DVal:            prop.DVal,
		DVar:            prop.DVar,
		BaseArmor:       prop.BaseArmor,
		ArmorVsSlash:    prop.ArmorVsSlash,
		ArmorVsPierce:   prop.ArmorVsPierce,
		ArmorVsBludgeon: prop.ArmorVsBludgeon,
		ArmorVsCold:     prop.ArmorVsCold,
		ArmorVsFire:     prop.ArmorVsFire,
		ArmorVsAcid:     prop.ArmorVsAcid,
		ArmorVsElectric: prop.ArmorVsElectric,
		ArmorVsNether:   prop.ArmorVsNether,
		BH:              prop.BH,
		HLF:             prop.HLF,
		MLF:             prop.MLF,
		LLF:             prop.LLF,
		HRF:             prop.HRF,
		MRF:             prop.MRF,
		LRF:             prop.LRF,
		HLB:             prop.HLB,
		MLB:             prop.MLB,
		LLB:             prop.LLB,
		HRB:             prop.HRB,
		MRB:             prop.MRB,
		LRB:             prop.LRB,
	}
}
