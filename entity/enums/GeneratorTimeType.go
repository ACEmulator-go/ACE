package enums


type GeneratorTimeType int32

const (
    GeneratorTimeTypeUndef GeneratorTimeType = GeneratorTimeType(0)
    GeneratorTimeTypeRealTime GeneratorTimeType = GeneratorTimeType(1)
    GeneratorTimeTypeDefined GeneratorTimeType = GeneratorTimeType(2)
    GeneratorTimeTypeEvent GeneratorTimeType = GeneratorTimeType(3)
    GeneratorTimeTypeNight GeneratorTimeType = GeneratorTimeType(4)
    )
