package enums

type SessionTerminationPhase int32

const (
	SessionTerminationPhaseInitialized          SessionTerminationPhase = SessionTerminationPhase(0)
	SessionTerminationPhaseSessionWorkCompleted SessionTerminationPhase = SessionTerminationPhase(1)
)
