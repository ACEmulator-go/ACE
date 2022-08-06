package enums

type CharacterGenerationVerificationResponse int32

const (
	CharacterGenerationVerificationResponseUndef                CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(0)
	CharacterGenerationVerificationResponseOk                   CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(1)
	CharacterGenerationVerificationResponsePending              CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(2)
	CharacterGenerationVerificationResponseNameInUse            CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(3)
	CharacterGenerationVerificationResponseNameBanned           CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(4)
	CharacterGenerationVerificationResponseCorrupt              CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(5)
	CharacterGenerationVerificationResponseDatabaseDown         CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(6)
	CharacterGenerationVerificationResponseAdminPrivilegeDenied CharacterGenerationVerificationResponse = CharacterGenerationVerificationResponse(7)
)
