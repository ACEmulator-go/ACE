package enums

type Usable uint32

const (
	UsableUndef                                   Usable = Usable(0x00)
	UsableNo                                      Usable = Usable(0x01)
	UsableSelf                                    Usable = Usable(0x02)
	UsableWielded                                 Usable = Usable(0x04)
	UsableContained                               Usable = Usable(0x08)
	UsableViewed                                  Usable = Usable(0x10)
	UsableRemote                                  Usable = Usable(0x20)
	UsableNeverWalk                               Usable = Usable(0x40)
	UsableObjSelf                                 Usable = Usable(0x80)
	UsableContainedViewed                         Usable = Usable(UsableContained | UsableViewed)
	UsableContainedViewedRemote                   Usable = Usable(UsableContained | UsableViewed | UsableRemote)
	UsableContainedViewedRemoteNeverWalk          Usable = Usable(UsableContained | UsableViewed | UsableRemote | UsableNeverWalk)
	UsableViewedRemote                            Usable = Usable(UsableViewed | UsableRemote)
	UsableViewedRemoteNeverWalk                   Usable = Usable(UsableViewed | UsableRemote | UsableNeverWalk)
	UsableRemoteNeverWalk                         Usable = Usable(UsableRemote | UsableNeverWalk)
	UsableSourceWieldedTargetWielded              Usable = Usable(0x040004)
	UsableSourceWieldedTargetContained            Usable = Usable(0x080004)
	UsableSourceWieldedTargetViewed               Usable = Usable(0x100004)
	UsableSourceWieldedTargetRemote               Usable = Usable(0x200004)
	UsableSourceWieldedTargetRemoteNeverWalk      Usable = Usable(0x600004)
	UsableSourceContainedTargetWielded            Usable = Usable(0x040008)
	UsableSourceContainedTargetContained          Usable = Usable(0x080008)
	UsableSourceContainedTargetObjselfOrContained Usable = Usable(0x880008)
	UsableSourceContainedTargetSelfOrContained    Usable = Usable(0x0A0008)
	UsableSourceContainedTargetViewed             Usable = Usable(0x100008)
	UsableSourceContainedTargetRemote             Usable = Usable(0x200008)
	UsableSourceContainedTargetRemoteNeverWalk    Usable = Usable(0x600008)
	UsableSourceContainedTargetRemoteOrSelf       Usable = Usable(0x220008)
	UsableSourceViewedTargetWielded               Usable = Usable(0x040010)
	UsableSourceViewedTargetContained             Usable = Usable(0x080010)
	UsableSourceViewedTargetViewed                Usable = Usable(0x100010)
	UsableSourceViewedTargetRemote                Usable = Usable(0x200010)
	UsableSourceRemoteTargetWielded               Usable = Usable(0x040020)
	UsableSourceRemoteTargetContained             Usable = Usable(0x080020)
	UsableSourceRemoteTargetViewed                Usable = Usable(0x100020)
	UsableSourceRemoteTargetRemote                Usable = Usable(0x200020)
	UsableSourceRemoteTargetRemoteNeverWalk       Usable = Usable(0x600020)
	UsableSourceMask                              Usable = Usable(0xFFFF)
	UsableTargetMask                              Usable = Usable(0xFFFF0000)
)
