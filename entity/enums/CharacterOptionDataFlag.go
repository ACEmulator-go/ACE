package enums


type CharacterOptionDataFlag uint32

const (
    CharacterOptionDataFlagShortcut CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000001)
    CharacterOptionDataFlagSquelchList CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000002)
    CharacterOptionDataFlagMultiSpellList CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000004)
    CharacterOptionDataFlagDesiredComps CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000008)
    CharacterOptionDataFlagExtendedMultiSpellLists CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000010)
    CharacterOptionDataFlagSpellbookFilters CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000020)
    CharacterOptionDataFlagCharacterOptions2 CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000040)
    CharacterOptionDataFlagTimestampFormat CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000080)
    CharacterOptionDataFlagGenericQualitiesData CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000100)
    CharacterOptionDataFlagGameplayOptions CharacterOptionDataFlag = CharacterOptionDataFlag(0x00000200)
    )
