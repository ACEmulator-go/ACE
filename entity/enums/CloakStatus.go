package enums


type CloakStatus int32

const (
    CloakStatusUndef CloakStatus = CloakStatus(0)
    CloakStatusOff CloakStatus = CloakStatus(1)
    CloakStatusOn CloakStatus = CloakStatus(2)
    CloakStatusPlayer CloakStatus = CloakStatus(3)
    )
