package enums


type ChatFilterMask int32

const (
    ChatFilterMaskSpeech ChatFilterMask = ChatFilterMask(0x00000004)
    ChatFilterMaskTell ChatFilterMask = ChatFilterMask(0x00000008)
    ChatFilterMaskCombat ChatFilterMask = ChatFilterMask(0x00000040)
    ChatFilterMaskMagic ChatFilterMask = ChatFilterMask(0x00000080)
    ChatFilterMaskEmote ChatFilterMask = ChatFilterMask(0x00001000)
    ChatFilterMaskAppraisal ChatFilterMask = ChatFilterMask(0x00010000)
    ChatFilterMaskSpellcasting ChatFilterMask = ChatFilterMask(0x00020000)
    ChatFilterMaskAllegiance ChatFilterMask = ChatFilterMask(0x00040000)
    ChatFilterMaskFellowship ChatFilterMask = ChatFilterMask(0x00080000)
    ChatFilterMaskCombat_Enemy ChatFilterMask = ChatFilterMask(0x00200000)
    ChatFilterMaskCombat_Self ChatFilterMask = ChatFilterMask(0x00400000)
    ChatFilterMaskRecall ChatFilterMask = ChatFilterMask(0x00800000)
    ChatFilterMaskCraft ChatFilterMask = ChatFilterMask(0x01000000)
    ChatFilterMaskSalvaging ChatFilterMask = ChatFilterMask(0x02000000)
    )
