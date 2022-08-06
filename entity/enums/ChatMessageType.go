package enums


type ChatMessageType uint32

const (
    ChatMessageTypeBroadcast ChatMessageType = ChatMessageType(0x00)
    ChatMessageTypeAllChannels ChatMessageType = ChatMessageType(0x01)
    ChatMessageTypeSpeech ChatMessageType = ChatMessageType(0x02)
    ChatMessageTypeTell ChatMessageType = ChatMessageType(0x03)
    ChatMessageTypeOutgoingTell ChatMessageType = ChatMessageType(0x04)
    ChatMessageTypeSystem ChatMessageType = ChatMessageType(0x05)
    ChatMessageTypeCombat ChatMessageType = ChatMessageType(0x06)
    ChatMessageTypeMagic ChatMessageType = ChatMessageType(0x07)
    ChatMessageTypeChannel ChatMessageType = ChatMessageType(0x08)
    ChatMessageTypeChannelSend ChatMessageType = ChatMessageType(0x09)
    ChatMessageTypeSocial ChatMessageType = ChatMessageType(0x0A)
    ChatMessageTypeSocialSend ChatMessageType = ChatMessageType(0x0B)
    ChatMessageTypeEmote ChatMessageType = ChatMessageType(0x0C)
    ChatMessageTypeAdvancement ChatMessageType = ChatMessageType(0x0D)
    ChatMessageTypeAbuse ChatMessageType = ChatMessageType(0x0E)
    ChatMessageTypeHelp ChatMessageType = ChatMessageType(0x0F)
    ChatMessageTypeAppraisal ChatMessageType = ChatMessageType(0x10)
    ChatMessageTypeSpellcasting ChatMessageType = ChatMessageType(0x11)
    ChatMessageTypeAllegiance ChatMessageType = ChatMessageType(0x12)
    ChatMessageTypeFellowship ChatMessageType = ChatMessageType(0x13)
    ChatMessageTypeWorldBroadcast ChatMessageType = ChatMessageType(0x14)
    ChatMessageTypeCombatEnemy ChatMessageType = ChatMessageType(0x15)
    ChatMessageTypeCombatSelf ChatMessageType = ChatMessageType(0x16)
    ChatMessageTypeRecall ChatMessageType = ChatMessageType(0x17)
    ChatMessageTypeCraft ChatMessageType = ChatMessageType(0x18)
    ChatMessageTypeSalvaging ChatMessageType = ChatMessageType(0x19)
    ChatMessageTypex1B ChatMessageType = ChatMessageType(0x1B)
    ChatMessageTypex1C ChatMessageType = ChatMessageType(0x1C)
    ChatMessageTypex1D ChatMessageType = ChatMessageType(0x1D)
    ChatMessageTypeAdminTell ChatMessageType = ChatMessageType(0x1F)
    )
