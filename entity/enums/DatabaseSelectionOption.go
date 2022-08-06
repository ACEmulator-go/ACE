package enums


type DatabaseSelectionOption int32

const (
    DatabaseSelectionOptionNone DatabaseSelectionOption = DatabaseSelectionOption(0)
    DatabaseSelectionOptionAuthentication DatabaseSelectionOption = DatabaseSelectionOption(1)
    DatabaseSelectionOptionShard DatabaseSelectionOption = DatabaseSelectionOption(2)
    DatabaseSelectionOptionWorld DatabaseSelectionOption = DatabaseSelectionOption(3)
    )
