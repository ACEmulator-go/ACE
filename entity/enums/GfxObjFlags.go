package enums


type GfxObjFlags uint32

const (
    GfxObjFlagsHasPhysics GfxObjFlags = GfxObjFlags(0x1)
    GfxObjFlagsHasDrawing GfxObjFlags = GfxObjFlags(0x2)
    GfxObjFlagsUnknown GfxObjFlags = GfxObjFlags(0x4)
    )
