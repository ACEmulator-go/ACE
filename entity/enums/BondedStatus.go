package enums


type BondedStatus int32

const (
    BondedStatusDestroy BondedStatus = BondedStatus(-2)
    BondedStatusSlippery BondedStatus = BondedStatus(-1)
    BondedStatusNormal BondedStatus = BondedStatus(0)
    BondedStatusBonded BondedStatus = BondedStatus(1)
    )
