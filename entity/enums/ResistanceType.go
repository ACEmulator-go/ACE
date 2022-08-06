package enums


type ResistanceType int32

const (
    ResistanceTypeUndef ResistanceType = ResistanceType(0)
    ResistanceTypeSlash ResistanceType = ResistanceType(1)
    ResistanceTypePierce ResistanceType = ResistanceType(2)
    ResistanceTypeBludgeon ResistanceType = ResistanceType(3)
    ResistanceTypeFire ResistanceType = ResistanceType(4)
    ResistanceTypeCold ResistanceType = ResistanceType(5)
    ResistanceTypeAcid ResistanceType = ResistanceType(6)
    ResistanceTypeElectric ResistanceType = ResistanceType(7)
    ResistanceTypeNether ResistanceType = ResistanceType(8)
    ResistanceTypeHealthDrain ResistanceType = ResistanceType(9)
    ResistanceTypeHealthBoost ResistanceType = ResistanceType(10)
    ResistanceTypeStaminaDrain ResistanceType = ResistanceType(11)
    ResistanceTypeStaminaBoost ResistanceType = ResistanceType(12)
    ResistanceTypeManaDrain ResistanceType = ResistanceType(13)
    )
