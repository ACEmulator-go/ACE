package enums


type EmoteType int32

const (
    EmoteTypeInvalid EmoteType = EmoteType(0)
    EmoteTypeInvalidVendor EmoteType = EmoteType(0)
    EmoteTypeAct EmoteType = EmoteType(1)
    EmoteTypeAwardXP EmoteType = EmoteType(2)
    EmoteTypeGive EmoteType = EmoteType(3)
    EmoteTypeMoveHome EmoteType = EmoteType(4)
    EmoteTypeMotion EmoteType = EmoteType(5)
    EmoteTypeMove EmoteType = EmoteType(6)
    EmoteTypePhysScript EmoteType = EmoteType(7)
    EmoteTypeSay EmoteType = EmoteType(8)
    EmoteTypeSound EmoteType = EmoteType(9)
    EmoteTypeTell EmoteType = EmoteType(10)
    EmoteTypeTurn EmoteType = EmoteType(11)
    EmoteTypeTurnToTarget EmoteType = EmoteType(12)
    EmoteTypeTextDirect EmoteType = EmoteType(13)
    EmoteTypeCastSpell EmoteType = EmoteType(14)
    EmoteTypeActivate EmoteType = EmoteType(15)
    EmoteTypeWorldBroadcast EmoteType = EmoteType(16)
    EmoteTypeLocalBroadcast EmoteType = EmoteType(17)
    EmoteTypeDirectBroadcast EmoteType = EmoteType(18)
    EmoteTypeCastSpellInstant EmoteType = EmoteType(19)
    EmoteTypeUpdateQuest EmoteType = EmoteType(20)
    EmoteTypeInqQuest EmoteType = EmoteType(21)
    EmoteTypeStampQuest EmoteType = EmoteType(22)
    EmoteTypeStartEvent EmoteType = EmoteType(23)
    EmoteTypeStopEvent EmoteType = EmoteType(24)
    EmoteTypeBLog EmoteType = EmoteType(25)
    EmoteTypeAdminSpam EmoteType = EmoteType(26)
    EmoteTypeTeachSpell EmoteType = EmoteType(27)
    EmoteTypeAwardSkillXP EmoteType = EmoteType(28)
    EmoteTypeAwardSkillPoints EmoteType = EmoteType(29)
    EmoteTypeInqQuestSolves EmoteType = EmoteType(30)
    EmoteTypeEraseQuest EmoteType = EmoteType(31)
    EmoteTypeDecrementQuest EmoteType = EmoteType(32)
    EmoteTypeIncrementQuest EmoteType = EmoteType(33)
    EmoteTypeAddCharacterTitle EmoteType = EmoteType(34)
    EmoteTypeInqBoolStat EmoteType = EmoteType(35)
    EmoteTypeInqIntStat EmoteType = EmoteType(36)
    EmoteTypeInqFloatStat EmoteType = EmoteType(37)
    EmoteTypeInqStringStat EmoteType = EmoteType(38)
    EmoteTypeInqAttributeStat EmoteType = EmoteType(39)
    EmoteTypeInqRawAttributeStat EmoteType = EmoteType(40)
    EmoteTypeInqSecondaryAttributeStat EmoteType = EmoteType(41)
    EmoteTypeInqRawSecondaryAttributeStat EmoteType = EmoteType(42)
    EmoteTypeInqSkillStat EmoteType = EmoteType(43)
    EmoteTypeInqRawSkillStat EmoteType = EmoteType(44)
    EmoteTypeInqSkillTrained EmoteType = EmoteType(45)
    EmoteTypeInqSkillSpecialized EmoteType = EmoteType(46)
    EmoteTypeAwardTrainingCredits EmoteType = EmoteType(47)
    EmoteTypeInflictVitaePenalty EmoteType = EmoteType(48)
    EmoteTypeAwardLevelProportionalXP EmoteType = EmoteType(49)
    EmoteTypeAwardLevelProportionalSkillXP EmoteType = EmoteType(50)
    EmoteTypeInqEvent EmoteType = EmoteType(51)
    EmoteTypeForceMotion EmoteType = EmoteType(52)
    EmoteTypeSetIntStat EmoteType = EmoteType(53)
    EmoteTypeIncrementIntStat EmoteType = EmoteType(54)
    EmoteTypeDecrementIntStat EmoteType = EmoteType(55)
    EmoteTypeCreateTreasure EmoteType = EmoteType(56)
    EmoteTypeResetHomePosition EmoteType = EmoteType(57)
    EmoteTypeInqFellowQuest EmoteType = EmoteType(58)
    EmoteTypeInqFellowNum EmoteType = EmoteType(59)
    EmoteTypeUpdateFellowQuest EmoteType = EmoteType(60)
    EmoteTypeStampFellowQuest EmoteType = EmoteType(61)
    EmoteTypeAwardNoShareXP EmoteType = EmoteType(62)
    EmoteTypeSetSanctuaryPosition EmoteType = EmoteType(63)
    EmoteTypeTellFellow EmoteType = EmoteType(64)
    EmoteTypeFellowBroadcast EmoteType = EmoteType(65)
    EmoteTypeLockFellow EmoteType = EmoteType(66)
    EmoteTypeGoto EmoteType = EmoteType(67)
    EmoteTypePopUp EmoteType = EmoteType(68)
    EmoteTypeSetBoolStat EmoteType = EmoteType(69)
    EmoteTypeSetQuestCompletions EmoteType = EmoteType(70)
    EmoteTypeInqNumCharacterTitles EmoteType = EmoteType(71)
    EmoteTypeGenerate EmoteType = EmoteType(72)
    EmoteTypePetCastSpellOnOwner EmoteType = EmoteType(73)
    EmoteTypeTakeItems EmoteType = EmoteType(74)
    EmoteTypeInqYesNo EmoteType = EmoteType(75)
    EmoteTypeInqOwnsItems EmoteType = EmoteType(76)
    EmoteTypeDeleteSelf EmoteType = EmoteType(77)
    EmoteTypeKillSelf EmoteType = EmoteType(78)
    EmoteTypeUpdateMyQuest EmoteType = EmoteType(79)
    EmoteTypeInqMyQuest EmoteType = EmoteType(80)
    EmoteTypeStampMyQuest EmoteType = EmoteType(81)
    EmoteTypeInqMyQuestSolves EmoteType = EmoteType(82)
    EmoteTypeEraseMyQuest EmoteType = EmoteType(83)
    EmoteTypeDecrementMyQuest EmoteType = EmoteType(84)
    EmoteTypeIncrementMyQuest EmoteType = EmoteType(85)
    EmoteTypeSetMyQuestCompletions EmoteType = EmoteType(86)
    EmoteTypeMoveToPos EmoteType = EmoteType(87)
    EmoteTypeLocalSignal EmoteType = EmoteType(88)
    EmoteTypeInqPackSpace EmoteType = EmoteType(89)
    EmoteTypeRemoveVitaePenalty EmoteType = EmoteType(90)
    EmoteTypeSetEyeTexture EmoteType = EmoteType(91)
    EmoteTypeSetEyePalette EmoteType = EmoteType(92)
    EmoteTypeSetNoseTexture EmoteType = EmoteType(93)
    EmoteTypeSetNosePalette EmoteType = EmoteType(94)
    EmoteTypeSetMouthTexture EmoteType = EmoteType(95)
    EmoteTypeSetMouthPalette EmoteType = EmoteType(96)
    EmoteTypeSetHeadObject EmoteType = EmoteType(97)
    EmoteTypeSetHeadPalette EmoteType = EmoteType(98)
    EmoteTypeTeleportTarget EmoteType = EmoteType(99)
    EmoteTypeTeleportSelf EmoteType = EmoteType(100)
    EmoteTypeStartBarber EmoteType = EmoteType(101)
    EmoteTypeInqQuestBitsOn EmoteType = EmoteType(102)
    EmoteTypeInqQuestBitsOff EmoteType = EmoteType(103)
    EmoteTypeInqMyQuestBitsOn EmoteType = EmoteType(104)
    EmoteTypeInqMyQuestBitsOff EmoteType = EmoteType(105)
    EmoteTypeSetQuestBitsOn EmoteType = EmoteType(106)
    EmoteTypeSetQuestBitsOff EmoteType = EmoteType(107)
    EmoteTypeSetMyQuestBitsOn EmoteType = EmoteType(108)
    EmoteTypeSetMyQuestBitsOff EmoteType = EmoteType(109)
    EmoteTypeUntrainSkill EmoteType = EmoteType(110)
    EmoteTypeSetAltRacialSkills EmoteType = EmoteType(111)
    EmoteTypeSpendLuminance EmoteType = EmoteType(112)
    EmoteTypeAwardLuminance EmoteType = EmoteType(113)
    EmoteTypeInqInt64Stat EmoteType = EmoteType(114)
    EmoteTypeSetInt64Stat EmoteType = EmoteType(115)
    EmoteTypeOpenMe EmoteType = EmoteType(116)
    EmoteTypeCloseMe EmoteType = EmoteType(117)
    EmoteTypeSetFloatStat EmoteType = EmoteType(118)
    EmoteTypeAddContract EmoteType = EmoteType(119)
    EmoteTypeRemoveContract EmoteType = EmoteType(120)
    EmoteTypeInqContractsFull EmoteType = EmoteType(121)
    )