package enums


type ConfirmationType uint32

const (
    ConfirmationTypeUndefined ConfirmationType = ConfirmationType(0x00)
    ConfirmationTypeSwearAllegiance ConfirmationType = ConfirmationType(0x01)
    ConfirmationTypeAlterSkill ConfirmationType = ConfirmationType(0x02)
    ConfirmationTypeAlterAttribute ConfirmationType = ConfirmationType(0x03)
    ConfirmationTypeFellowship ConfirmationType = ConfirmationType(0x04)
    ConfirmationTypeCraftInteraction ConfirmationType = ConfirmationType(0x05)
    ConfirmationTypeAugmentation ConfirmationType = ConfirmationType(0x06)
    )
