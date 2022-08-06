package enums

type NetAuthType uint32

const (
	NetAuthTypeUndef           NetAuthType = NetAuthType(0x00000000)
	NetAuthTypeAccount         NetAuthType = NetAuthType(0x00000001)
	NetAuthTypeAccountPassword NetAuthType = NetAuthType(0x00000002)
)
