package enums


type AmmoType uint16

const (
    AmmoTypeNone AmmoType = AmmoType(0x0)
    AmmoTypeArrow AmmoType = AmmoType(0x1)
    AmmoTypeBolt AmmoType = AmmoType(0x2)
    AmmoTypeAtlatl AmmoType = AmmoType(0x4)
    AmmoTypeArrowCrystal AmmoType = AmmoType(0x8)
    AmmoTypeBoltCrystal AmmoType = AmmoType(0x10)
    AmmoTypeAtlatlCrystal AmmoType = AmmoType(0x20)
    AmmoTypeArrowChorizite AmmoType = AmmoType(0x40)
    AmmoTypeBoltChorizite AmmoType = AmmoType(0x80)
    )
