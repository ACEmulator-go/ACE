package enums


type CommandMask uint32

const (
    CommandMaskStyle CommandMask = CommandMask(0x80000000)
    CommandMaskSubState CommandMask = CommandMask(0x40000000)
    CommandMaskModifier CommandMask = CommandMask(0x20000000)
    CommandMaskAction CommandMask = CommandMask(0x10000000)
    CommandMaskUI CommandMask = CommandMask(0x08000000)
    CommandMaskToggle CommandMask = CommandMask(0x04000000)
    CommandMaskChatEmote CommandMask = CommandMask(0x02000000)
    CommandMaskMappable CommandMask = CommandMask(0x01000000)
    )
