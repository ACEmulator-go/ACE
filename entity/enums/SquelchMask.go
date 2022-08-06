package enums

type SquelchMask uint32

const (
	SquelchMaskNone         SquelchMask = SquelchMask(0x0)
	SquelchMaskSpeech       SquelchMask = SquelchMask(0x4)
	SquelchMaskTell         SquelchMask = SquelchMask(0x8)
	SquelchMaskCombat       SquelchMask = SquelchMask(0x40)
	SquelchMaskMagic        SquelchMask = SquelchMask(0x80)
	SquelchMaskEmote        SquelchMask = SquelchMask(0x1000)
	SquelchMaskAppraisal    SquelchMask = SquelchMask(0x10000)
	SquelchMaskSpellcasting SquelchMask = SquelchMask(0x20000)
	SquelchMaskAllegiance   SquelchMask = SquelchMask(0x40000)
	SquelchMaskFellowship   SquelchMask = SquelchMask(0x80000)
	SquelchMaskCombatEnemy  SquelchMask = SquelchMask(0x200000)
	SquelchMaskCombatSelf   SquelchMask = SquelchMask(0x400000)
	SquelchMaskRecall       SquelchMask = SquelchMask(0x800000)
	SquelchMaskCraft        SquelchMask = SquelchMask(0x1000000)
	SquelchMaskSalvaging    SquelchMask = SquelchMask(0x2000000)
	SquelchMaskCombined     SquelchMask = SquelchMask(SquelchMaskSpeech | SquelchMaskTell | SquelchMaskCombat | SquelchMaskMagic | SquelchMaskEmote | SquelchMaskAppraisal | SquelchMaskSpellcasting | SquelchMaskAllegiance | SquelchMaskFellowship | SquelchMaskCombatEnemy | SquelchMaskCombatSelf | SquelchMaskRecall | SquelchMaskCraft | SquelchMaskSalvaging)
)
