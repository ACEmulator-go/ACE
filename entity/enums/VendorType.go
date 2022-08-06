package enums


type VendorType int32

const (
    VendorTypeUndef VendorType = VendorType(0)
    VendorTypeOpen VendorType = VendorType(1)
    VendorTypeClose VendorType = VendorType(2)
    VendorTypeSell VendorType = VendorType(3)
    VendorTypeBuy VendorType = VendorType(4)
    )
