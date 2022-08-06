package enums

type SkillAdvancementClass uint32

const (
	SkillAdvancementClassInactive  SkillAdvancementClass = SkillAdvancementClass(0)
	SkillAdvancementClassUntrained SkillAdvancementClass = SkillAdvancementClass(1)
	SkillAdvancementClassTrained   SkillAdvancementClass = SkillAdvancementClass(2)
)
