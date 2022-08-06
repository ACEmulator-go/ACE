package enums


type AccountLookupType int32

const (
    AccountLookupTypeUndef AccountLookupType = AccountLookupType(0)
    AccountLookupTypeSubscription AccountLookupType = AccountLookupType(1)
    AccountLookupTypeCharacter AccountLookupType = AccountLookupType(2)
    AccountLookupTypeIid AccountLookupType = AccountLookupType(3)
    )
