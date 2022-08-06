package enums


type AppraisalLongDescDecorations int32

const (
    AppraisalLongDescDecorationsNone AppraisalLongDescDecorations = AppraisalLongDescDecorations(0x0)
    AppraisalLongDescDecorationsPrependWorkmanship AppraisalLongDescDecorations = AppraisalLongDescDecorations(0x1)
    AppraisalLongDescDecorationsPrependMaterial AppraisalLongDescDecorations = AppraisalLongDescDecorations(0x2)
    AppraisalLongDescDecorationsAppendGemInfo AppraisalLongDescDecorations = AppraisalLongDescDecorations(0x4)
    )
