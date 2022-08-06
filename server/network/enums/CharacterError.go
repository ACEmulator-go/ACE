package enums

type CharacterError int32

const (
	CharacterErrorLogon                           CharacterError = CharacterError(0x00000001)
	CharacterErrorAccountLogin                    CharacterError = CharacterError(0x00000003)
	CharacterErrorServerCrash1                    CharacterError = CharacterError(0x00000004)
	CharacterErrorLogoff                          CharacterError = CharacterError(0x00000005)
	CharacterErrorDelete                          CharacterError = CharacterError(0x00000006)
	CharacterErrorServerCrash2                    CharacterError = CharacterError(0x00000008)
	CharacterErrorAccountInvalid                  CharacterError = CharacterError(0x00000009)
	CharacterErrorAccountDoesntExist              CharacterError = CharacterError(0x0000000A)
	CharacterErrorEnterGameGeneric                CharacterError = CharacterError(0x0000000B)
	CharacterErrorEnterGameStressAccount          CharacterError = CharacterError(0x0000000C)
	CharacterErrorEnterGameCharacterInWorld       CharacterError = CharacterError(0x0000000D)
	CharacterErrorEnterGamePlayerAccountMissing   CharacterError = CharacterError(0x0000000E)
	CharacterErrorEnterGameCharacterNotOwned      CharacterError = CharacterError(0x0000000F)
	CharacterErrorEnterGameCharacterInWorldServer CharacterError = CharacterError(0x00000010)
	CharacterErrorEnterGameOldCharacter           CharacterError = CharacterError(0x00000011)
	CharacterErrorEnterGameCorruptCharacter       CharacterError = CharacterError(0x00000012)
	CharacterErrorEnterGameStartServerDown        CharacterError = CharacterError(0x00000013)
	CharacterErrorEnterGameCouldntPlaceCharacter  CharacterError = CharacterError(0x00000014)
	CharacterErrorLogonServerFull                 CharacterError = CharacterError(0x00000015)
	CharacterErrorEnterGameCharacterLocked        CharacterError = CharacterError(0x00000017)
)
