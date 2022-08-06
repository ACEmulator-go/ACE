package enums

type PlayerKillerStatus uint32

const (
	PlayerKillerStatusUndef       PlayerKillerStatus = PlayerKillerStatus(0x00)
	PlayerKillerStatusProtected   PlayerKillerStatus = PlayerKillerStatus(0x01)
	PlayerKillerStatusNPK         PlayerKillerStatus = PlayerKillerStatus(0x02)
	PlayerKillerStatusPK          PlayerKillerStatus = PlayerKillerStatus(0x04)
	PlayerKillerStatusUnprotected PlayerKillerStatus = PlayerKillerStatus(0x08)
	PlayerKillerStatusRubberGlue  PlayerKillerStatus = PlayerKillerStatus(0x10)
	PlayerKillerStatusFree        PlayerKillerStatus = PlayerKillerStatus(0x20)
	PlayerKillerStatusPKLite      PlayerKillerStatus = PlayerKillerStatus(0x40)
	PlayerKillerStatusCreature    PlayerKillerStatus = PlayerKillerStatus(PlayerKillerStatusUnprotected)
	PlayerKillerStatusTrap        PlayerKillerStatus = PlayerKillerStatus(PlayerKillerStatusUnprotected)
	PlayerKillerStatusNPC         PlayerKillerStatus = PlayerKillerStatus(PlayerKillerStatusProtected)
	PlayerKillerStatusVendor      PlayerKillerStatus = PlayerKillerStatus(PlayerKillerStatusRubberGlue)
)
