package enums


type SurfaceType uint32

const (
    SurfaceTypeBase1Solid SurfaceType = SurfaceType(0x1)
    SurfaceTypeBase1Image SurfaceType = SurfaceType(0x2)
    SurfaceTypeBase1ClipMap SurfaceType = SurfaceType(0x4)
    SurfaceTypeTranslucent SurfaceType = SurfaceType(0x10)
    SurfaceTypeDiffuse SurfaceType = SurfaceType(0x20)
    SurfaceTypeLuminous SurfaceType = SurfaceType(0x40)
    SurfaceTypeAlpha SurfaceType = SurfaceType(0x100)
    SurfaceTypeInvAlpha SurfaceType = SurfaceType(0x200)
    SurfaceTypeAdditive SurfaceType = SurfaceType(0x10000)
    SurfaceTypeDetail SurfaceType = SurfaceType(0x20000)
    SurfaceTypeGouraud SurfaceType = SurfaceType(0x10000000)
    SurfaceTypeStippled SurfaceType = SurfaceType(0x40000000)
    )
