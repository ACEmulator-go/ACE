package enums

type AddFriendResult int32

const (
	AddFriendResultFriendWithSelf        AddFriendResult = AddFriendResult(0)
	AddFriendResultAlreadyInList         AddFriendResult = AddFriendResult(1)
	AddFriendResultCharacterDoesNotExist AddFriendResult = AddFriendResult(2)
)
