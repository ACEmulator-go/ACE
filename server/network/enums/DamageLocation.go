package enums


type DamageLocation int32

const (
    DamageLocationHead DamageLocation = DamageLocation(0x0)
    DamageLocationChest DamageLocation = DamageLocation(0x1)
    DamageLocationAbdomen DamageLocation = DamageLocation(0x2)
    DamageLocationUpperArm DamageLocation = DamageLocation(0x3)
    DamageLocationLowerArm DamageLocation = DamageLocation(0x4)
    DamageLocationHand DamageLocation = DamageLocation(0x5)
    DamageLocationUpperLeg DamageLocation = DamageLocation(0x6)
    DamageLocationLowerLeg DamageLocation = DamageLocation(0x7)
    )
