package enums

type GenericQualitiesPackHeader int32

const (
	GenericQualitiesPackHeaderPacked_None        GenericQualitiesPackHeader = GenericQualitiesPackHeader(0)
	GenericQualitiesPackHeaderPacked_IntStats    GenericQualitiesPackHeader = GenericQualitiesPackHeader(1)
	GenericQualitiesPackHeaderPacked_BoolStats   GenericQualitiesPackHeader = GenericQualitiesPackHeader(2)
	GenericQualitiesPackHeaderPacked_FloatStats  GenericQualitiesPackHeader = GenericQualitiesPackHeader(4)
	GenericQualitiesPackHeaderPacked_StringStats GenericQualitiesPackHeader = GenericQualitiesPackHeader(8)
)
