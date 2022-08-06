package enums


type RadarBehavior byte

const (
    RadarBehaviorUndefined RadarBehavior = RadarBehavior(0)
    RadarBehaviorShowNever RadarBehavior = RadarBehavior(1)
    RadarBehaviorShowMovement RadarBehavior = RadarBehavior(2)
    RadarBehaviorShowAttacking RadarBehavior = RadarBehavior(3)
    )
