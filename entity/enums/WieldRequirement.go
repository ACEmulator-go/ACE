package enums


type WieldRequirement int32

const (
    WieldRequirementInvalid WieldRequirement = WieldRequirement(0)
    WieldRequirementSkill WieldRequirement = WieldRequirement(1)
    WieldRequirementRawSkill WieldRequirement = WieldRequirement(2)
    WieldRequirementAttrib WieldRequirement = WieldRequirement(3)
    WieldRequirementRawAttrib WieldRequirement = WieldRequirement(4)
    WieldRequirementSecondaryAttrib WieldRequirement = WieldRequirement(5)
    WieldRequirementRawSecondaryAttrib WieldRequirement = WieldRequirement(6)
    WieldRequirementLevel WieldRequirement = WieldRequirement(7)
    WieldRequirementTraining WieldRequirement = WieldRequirement(8)
    WieldRequirementIntStat WieldRequirement = WieldRequirement(9)
    WieldRequirementBoolStat WieldRequirement = WieldRequirement(10)
    WieldRequirementCreatureType WieldRequirement = WieldRequirement(11)
    )
