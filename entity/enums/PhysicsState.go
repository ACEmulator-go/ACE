package enums


type PhysicsState int32

const (
    PhysicsStateStatic PhysicsState = PhysicsState(0x00000001)
    PhysicsStateUnused1 PhysicsState = PhysicsState(0x00000002)
    PhysicsStateEthereal PhysicsState = PhysicsState(0x00000004)
    PhysicsStateReportCollisions PhysicsState = PhysicsState(0x00000008)
    PhysicsStateIgnoreCollisions PhysicsState = PhysicsState(0x00000010)
    PhysicsStateNoDraw PhysicsState = PhysicsState(0x00000020)
    PhysicsStateMissile PhysicsState = PhysicsState(0x00000040)
    PhysicsStatePushable PhysicsState = PhysicsState(0x00000080)
    PhysicsStateAlignPath PhysicsState = PhysicsState(0x00000100)
    PhysicsStatePathClipped PhysicsState = PhysicsState(0x00000200)
    PhysicsStateGravity PhysicsState = PhysicsState(0x00000400)
    PhysicsStateLightingOn PhysicsState = PhysicsState(0x00000800)
    PhysicsStateParticleEmitter PhysicsState = PhysicsState(0x00001000)
    PhysicsStateUnused2 PhysicsState = PhysicsState(0x00002000)
    PhysicsStateHidden PhysicsState = PhysicsState(0x00004000)
    PhysicsStateScriptedCollision PhysicsState = PhysicsState(0x00008000)
    PhysicsStateHasPhysicsBSP PhysicsState = PhysicsState(0x00010000)
    PhysicsStateInelastic PhysicsState = PhysicsState(0x00020000)
    PhysicsStateHasDefaultAnim PhysicsState = PhysicsState(0x00040000)
    PhysicsStateHasDefaultScript PhysicsState = PhysicsState(0x00080000)
    PhysicsStateCloaked PhysicsState = PhysicsState(0x00100000)
    PhysicsStateReportCollisionsAsEnvironment PhysicsState = PhysicsState(0x00200000)
    PhysicsStateEdgeSlide PhysicsState = PhysicsState(0x00400000)
    PhysicsStateSledding PhysicsState = PhysicsState(0x00800000)
    PhysicsStateFrozen PhysicsState = PhysicsState(0x01000000)
    )
