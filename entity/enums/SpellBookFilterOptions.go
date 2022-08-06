package enums


type SpellBookFilterOptions uint32

const (
    SpellBookFilterOptionsNone SpellBookFilterOptions = SpellBookFilterOptions(0x0000)
    SpellBookFilterOptionsCreature SpellBookFilterOptions = SpellBookFilterOptions(0x0001)
    SpellBookFilterOptionsItem SpellBookFilterOptions = SpellBookFilterOptions(0x0002)
    SpellBookFilterOptionsLife SpellBookFilterOptions = SpellBookFilterOptions(0x0004)
    SpellBookFilterOptionsWar SpellBookFilterOptions = SpellBookFilterOptions(0x0008)
    SpellBookFilterOptionsLevel1 SpellBookFilterOptions = SpellBookFilterOptions(0x0010)
    SpellBookFilterOptionsLevel2 SpellBookFilterOptions = SpellBookFilterOptions(0x0020)
    SpellBookFilterOptionsLevel3 SpellBookFilterOptions = SpellBookFilterOptions(0x0040)
    SpellBookFilterOptionsLevel4 SpellBookFilterOptions = SpellBookFilterOptions(0x0080)
    SpellBookFilterOptionsLevel5 SpellBookFilterOptions = SpellBookFilterOptions(0x0100)
    SpellBookFilterOptionsLevel6 SpellBookFilterOptions = SpellBookFilterOptions(0x0200)
    SpellBookFilterOptionsLevel7 SpellBookFilterOptions = SpellBookFilterOptions(0x0400)
    SpellBookFilterOptionsLevel8 SpellBookFilterOptions = SpellBookFilterOptions(0x0800)
    SpellBookFilterOptionsLevel9 SpellBookFilterOptions = SpellBookFilterOptions(0x1000)
    SpellBookFilterOptionsVoid SpellBookFilterOptions = SpellBookFilterOptions(0x2000)
    )
