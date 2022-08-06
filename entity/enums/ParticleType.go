package enums


type ParticleType int32

const (
    ParticleTypeUnknown ParticleType = ParticleType(0)
    ParticleTypeStill ParticleType = ParticleType(1)
    ParticleTypeLocalVelocity ParticleType = ParticleType(2)
    ParticleTypeParabolicLVGA ParticleType = ParticleType(3)
    ParticleTypeParabolicLVGAGR ParticleType = ParticleType(4)
    ParticleTypeSwarm ParticleType = ParticleType(5)
    ParticleTypeExplode ParticleType = ParticleType(6)
    ParticleTypeImplode ParticleType = ParticleType(7)
    ParticleTypeParabolicLVLA ParticleType = ParticleType(8)
    ParticleTypeParabolicLVLALR ParticleType = ParticleType(9)
    ParticleTypeParabolicGVGA ParticleType = ParticleType(10)
    ParticleTypeParabolicGVGAGR ParticleType = ParticleType(11)
    ParticleTypeGlobalVelocity ParticleType = ParticleType(12)
    ParticleTypeNumParticleType ParticleType = ParticleType(13)
    )
