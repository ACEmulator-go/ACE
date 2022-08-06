package enums

type MutateFilter int32

const (
	MutateFilterUndef               MutateFilter = MutateFilter(0x0)
	MutateFilterArmorModVsAcid      MutateFilter = MutateFilter(0x1)
	MutateFilterArmorModVsCold      MutateFilter = MutateFilter(0x2)
	MutateFilterArmorModVsElectric  MutateFilter = MutateFilter(0x4)
	MutateFilterArmorModVsFire      MutateFilter = MutateFilter(0x8)
	MutateFilterEncumbranceVal      MutateFilter = MutateFilter(0x10)
	MutateFilterIcon                MutateFilter = MutateFilter(0x20)
	MutateFilterItemWorkmanship     MutateFilter = MutateFilter(0x40)
	MutateFilterLongDesc            MutateFilter = MutateFilter(0x80)
	MutateFilterName                MutateFilter = MutateFilter(0x100)
	MutateFilterResistItemAppraisal MutateFilter = MutateFilter(0x200)
	MutateFilterSetup               MutateFilter = MutateFilter(0x400)
	MutateFilterShieldValue         MutateFilter = MutateFilter(0x800)
	MutateFilterShortDesc           MutateFilter = MutateFilter(0x1000)
	MutateFilterValue               MutateFilter = MutateFilter(0x2000)
	MutateFilterWeaponTime          MutateFilter = MutateFilter(0x4000)
	MutateFilterArmorModVsType      MutateFilter = MutateFilter(MutateFilterArmorModVsAcid | MutateFilterArmorModVsCold | MutateFilterArmorModVsElectric | MutateFilterArmorModVsFire)
	MutateFilterBase                MutateFilter = MutateFilter(MutateFilterIcon | MutateFilterItemWorkmanship | MutateFilterLongDesc | MutateFilterName | MutateFilterResistItemAppraisal | MutateFilterSetup | MutateFilterShortDesc | MutateFilterValue)
)
