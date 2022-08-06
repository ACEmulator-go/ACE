package enums


type RecipeType int32

const (
    RecipeTypeNone RecipeType = RecipeType(0)
    RecipeTypeCreateItem RecipeType = RecipeType(1)
    RecipeTypeHealing RecipeType = RecipeType(2)
    RecipeTypeTinkering RecipeType = RecipeType(3)
    RecipeTypeDyeing RecipeType = RecipeType(4)
    RecipeTypeUnlocking RecipeType = RecipeType(5)
    )
