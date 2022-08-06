package enums


type ModificationOperation int32

const (
    ModificationOperationNone ModificationOperation = ModificationOperation(0)
    ModificationOperationSetValue ModificationOperation = ModificationOperation(1)
    ModificationOperationAdd ModificationOperation = ModificationOperation(2)
    ModificationOperationCopyFromSourceToTarget ModificationOperation = ModificationOperation(3)
    ModificationOperationCopyFromSourceToResult ModificationOperation = ModificationOperation(4)
    ModificationOperationUnknown1 ModificationOperation = ModificationOperation(5)
    ModificationOperationUnknown2 ModificationOperation = ModificationOperation(6)
    ModificationOperationAddSpell ModificationOperation = ModificationOperation(7)
    ModificationOperationSetBitsOn ModificationOperation = ModificationOperation(8)
    )
