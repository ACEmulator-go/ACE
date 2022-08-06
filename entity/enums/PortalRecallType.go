package enums


type PortalRecallType int32

const (
    PortalRecallTypeUndef PortalRecallType = PortalRecallType(0)
    PortalRecallTypeLastLifestone PortalRecallType = PortalRecallType(1)
    PortalRecallTypeLinkedLifestone PortalRecallType = PortalRecallType(2)
    PortalRecallTypeLastPortal PortalRecallType = PortalRecallType(3)
    PortalRecallTypeLinkedPortalOne PortalRecallType = PortalRecallType(4)
    )
