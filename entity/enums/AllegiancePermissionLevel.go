package enums


type AllegiancePermissionLevel int32

const (
    AllegiancePermissionLevelNone AllegiancePermissionLevel = AllegiancePermissionLevel(0)
    AllegiancePermissionLevelSpeaker AllegiancePermissionLevel = AllegiancePermissionLevel(1)
    AllegiancePermissionLevelSeneschal AllegiancePermissionLevel = AllegiancePermissionLevel(2)
    AllegiancePermissionLevelCastellan AllegiancePermissionLevel = AllegiancePermissionLevel(3)
    )
