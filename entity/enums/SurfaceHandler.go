package enums


type SurfaceHandler int32

const (
    SurfaceHandlerInvalid SurfaceHandler = SurfaceHandler(0x0)
    SurfaceHandlerDatabase SurfaceHandler = SurfaceHandler(0x1)
    SurfaceHandlerPalShift SurfaceHandler = SurfaceHandler(0x2)
    SurfaceHandlerTexMerge SurfaceHandler = SurfaceHandler(0x3)
    )
