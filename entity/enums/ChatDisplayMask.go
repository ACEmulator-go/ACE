package enums


type ChatDisplayMask int64

const (
    ChatDisplayMaskGameplay ChatDisplayMask = ChatDisplayMask(0x0000000003912021)
    ChatDisplayMaskMandatory ChatDisplayMask = ChatDisplayMask(0x000000000000c302)
    ChatDisplayMaskAreaChat ChatDisplayMask = ChatDisplayMask(0x0000000000001004)
    ChatDisplayMaskTells ChatDisplayMask = ChatDisplayMask(0x0000000000000018)
    ChatDisplayMaskCombat ChatDisplayMask = ChatDisplayMask(0x0000000000600040)
    ChatDisplayMaskMagic ChatDisplayMask = ChatDisplayMask(0x0000000000020080)
    ChatDisplayMaskAllegiance ChatDisplayMask = ChatDisplayMask(0x0000000000040c00)
    ChatDisplayMaskFellowship ChatDisplayMask = ChatDisplayMask(0x0000000000080000)
    ChatDisplayMaskErrors ChatDisplayMask = ChatDisplayMask(0x0000000004000000)
    ChatDisplayMaskGeneralChannel ChatDisplayMask = ChatDisplayMask(0x0000000008000000)
    ChatDisplayMaskTradeChannel ChatDisplayMask = ChatDisplayMask(0x0000000010000000)
    ChatDisplayMaskLFGChannel ChatDisplayMask = ChatDisplayMask(0x0000000020000000)
    ChatDisplayMaskRoleplayChannel ChatDisplayMask = ChatDisplayMask(0x0000000040000000)
    )
