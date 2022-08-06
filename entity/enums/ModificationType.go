package enums


type ModificationType int32

const (
    ModificationTypeSuccessTarget ModificationType = ModificationType(0)
    ModificationTypeSuccessSource ModificationType = ModificationType(1)
    ModificationTypeSuccessPlayer ModificationType = ModificationType(2)
    ModificationTypeSuccessResult ModificationType = ModificationType(3)
    ModificationTypeFailureTarget ModificationType = ModificationType(4)
    ModificationTypeFailureSource ModificationType = ModificationType(5)
    ModificationTypeFailurePlayer ModificationType = ModificationType(6)
    )
