package enums


type RecipeResult int32

const (
    RecipeResultSourceItemDestroyed RecipeResult = RecipeResult(1)
    RecipeResultTargetItemDestroyed RecipeResult = RecipeResult(2)
    RecipeResultSourceItemUsesDecrement RecipeResult = RecipeResult(4)
    RecipeResultTargetItemUsesDecrement RecipeResult = RecipeResult(8)
    RecipeResultSuccessItem1 RecipeResult = RecipeResult(16)
    RecipeResultSuccessItem2 RecipeResult = RecipeResult(32)
    RecipeResultFailureItem1 RecipeResult = RecipeResult(64)
    )
