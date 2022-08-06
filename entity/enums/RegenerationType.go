package enums


type RegenerationType uint32

const (
    RegenerationTypeUndef RegenerationType = RegenerationType(0x0)
    RegenerationTypeDestruction RegenerationType = RegenerationType(0x1)
    RegenerationTypePickUp RegenerationType = RegenerationType(0x2)
    )
