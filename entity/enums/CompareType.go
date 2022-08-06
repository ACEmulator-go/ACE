package enums


type CompareType int32

const (
    CompareTypeGreaterThan CompareType = CompareType(0)
    CompareTypeLessThanEqual CompareType = CompareType(1)
    CompareTypeLessThan CompareType = CompareType(2)
    CompareTypeGreaterThanEqual CompareType = CompareType(3)
    CompareTypeNotEqual CompareType = CompareType(4)
    CompareTypeNotEqualNotExist CompareType = CompareType(5)
    CompareTypeEqual CompareType = CompareType(6)
    CompareTypeNotExist CompareType = CompareType(7)
    CompareTypeExist CompareType = CompareType(8)
    CompareTypeNotHasBits CompareType = CompareType(9)
    )
