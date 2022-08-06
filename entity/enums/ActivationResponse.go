package enums

type ActivationResponse int32

const (
	ActivationResponseUndef     ActivationResponse = ActivationResponse(0)
	ActivationResponseUse       ActivationResponse = ActivationResponse(0x2)
	ActivationResponseAnimate   ActivationResponse = ActivationResponse(0x4)
	ActivationResponseTalk      ActivationResponse = ActivationResponse(0x10)
	ActivationResponseEmote     ActivationResponse = ActivationResponse(0x800)
	ActivationResponseCastSpell ActivationResponse = ActivationResponse(0x1000)
)
