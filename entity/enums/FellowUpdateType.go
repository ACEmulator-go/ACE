package enums


type FellowUpdateType int32

const (
    FellowUpdateTypeUndef FellowUpdateType = FellowUpdateType(0)
    FellowUpdateTypeFull FellowUpdateType = FellowUpdateType(1)
    FellowUpdateTypeStats FellowUpdateType = FellowUpdateType(2)
    )
