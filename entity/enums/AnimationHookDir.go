package enums


type AnimationHookDir int32

const (
    AnimationHookDirUnknown AnimationHookDir = AnimationHookDir(-2)
    AnimationHookDirBackward AnimationHookDir = AnimationHookDir(-1)
    AnimationHookDirBoth AnimationHookDir = AnimationHookDir(0)
    AnimationHookDirForward AnimationHookDir = AnimationHookDir(1)
    )
