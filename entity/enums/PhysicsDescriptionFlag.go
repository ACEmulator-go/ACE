package enums


type PhysicsDescriptionFlag int32

const (
    PhysicsDescriptionFlagNone PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000000)
    PhysicsDescriptionFlagVelocity PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000004)
    PhysicsDescriptionFlagAcceleration PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000008)
    PhysicsDescriptionFlagOmega PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000010)
    PhysicsDescriptionFlagParent PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000020)
    PhysicsDescriptionFlagChildren PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000040)
    PhysicsDescriptionFlagObjScale PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000080)
    PhysicsDescriptionFlagFriction PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000100)
    PhysicsDescriptionFlagElasticity PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000200)
    PhysicsDescriptionFlagTimestamps PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x000400)
    PhysicsDescriptionFlagDefaultScript PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x002000)
    PhysicsDescriptionFlagDefaultScriptIntensity PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x004000)
    PhysicsDescriptionFlagPosition PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x008000)
    PhysicsDescriptionFlagMovement PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x010000)
    PhysicsDescriptionFlagAnimationFrame PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x020000)
    PhysicsDescriptionFlagTranslucency PhysicsDescriptionFlag = PhysicsDescriptionFlag(0x040000)
    )
