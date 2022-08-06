package enums


type CoverageMask uint32

const (
    CoverageMaskUnknown CoverageMask = CoverageMask(0x00000001)
    CoverageMaskUnderwearUpperLegs CoverageMask = CoverageMask(0x00000002)
    CoverageMaskUnderwearLowerLegs CoverageMask = CoverageMask(0x00000004)
    CoverageMaskUnderwearChest CoverageMask = CoverageMask(0x00000008)
    CoverageMaskUnderwearAbdomen CoverageMask = CoverageMask(0x00000010)
    CoverageMaskUnderwearUpperArms CoverageMask = CoverageMask(0x00000020)
    CoverageMaskUnderwearLowerArms CoverageMask = CoverageMask(0x00000040)
    CoverageMaskOuterwearUpperLegs CoverageMask = CoverageMask(0x00000100)
    CoverageMaskOuterwearLowerLegs CoverageMask = CoverageMask(0x00000200)
    CoverageMaskOuterwearChest CoverageMask = CoverageMask(0x00000400)
    CoverageMaskOuterwearAbdomen CoverageMask = CoverageMask(0x00000800)
    CoverageMaskOuterwearUpperArms CoverageMask = CoverageMask(0x00001000)
    CoverageMaskOuterwearLowerArms CoverageMask = CoverageMask(0x00002000)
    CoverageMaskHead CoverageMask = CoverageMask(0x00004000)
    CoverageMaskHands CoverageMask = CoverageMask(0x00008000)
    CoverageMaskFeet CoverageMask = CoverageMask(0x00010000)
    )

type CoverageMaskHelper uint32

const (
    CoverageMaskHelperUnderwear CoverageMaskHelper = CoverageMaskHelper(CoverageMaskUnderwearUpperLegs | CoverageMaskUnderwearLowerLegs | CoverageMaskUnderwearChest | CoverageMaskUnderwearAbdomen | CoverageMaskUnderwearUpperArms | CoverageMaskUnderwearLowerArms)
    CoverageMaskHelperOuterwear CoverageMaskHelper = CoverageMaskHelper(CoverageMaskOuterwearUpperLegs | CoverageMaskOuterwearLowerLegs | CoverageMaskOuterwearChest | CoverageMaskOuterwearAbdomen | CoverageMaskOuterwearUpperArms | CoverageMaskOuterwearLowerArms | CoverageMaskHead | CoverageMaskHands | CoverageMaskFeet)
    CoverageMaskHelperUnderwearLegs CoverageMaskHelper = CoverageMaskHelper(CoverageMaskUnderwearUpperLegs | CoverageMaskUnderwearLowerLegs)
    CoverageMaskHelperUnderwearArms CoverageMaskHelper = CoverageMaskHelper(CoverageMaskUnderwearUpperArms | CoverageMaskUnderwearLowerArms)
    CoverageMaskHelperOuterwearLegs CoverageMaskHelper = CoverageMaskHelper(CoverageMaskOuterwearUpperLegs | CoverageMaskOuterwearLowerLegs)
    CoverageMaskHelperOuterwearArms CoverageMaskHelper = CoverageMaskHelper(CoverageMaskOuterwearUpperArms | CoverageMaskOuterwearLowerArms)
    CoverageMaskHelperUnderwearShirt CoverageMaskHelper = CoverageMaskHelper(CoverageMaskUnderwearChest | CoverageMaskUnderwearUpperArms | CoverageMaskUnderwearLowerArms)
    CoverageMaskHelperUnderwearPants CoverageMaskHelper = CoverageMaskHelper(CoverageMaskUnderwearUpperLegs | CoverageMaskUnderwearLowerLegs)
    CoverageMaskHelperExtremities CoverageMaskHelper = CoverageMaskHelper(CoverageMaskHead | CoverageMaskHands | CoverageMaskFeet)
    )
