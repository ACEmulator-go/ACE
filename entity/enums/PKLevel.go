package enums


type PKLevel uint32

const (
    PKLevelNPK PKLevel = PKLevel(0)
    PKLevelPK PKLevel = PKLevel(1)
    PKLevelPKLite PKLevel = PKLevel(2)
    )
