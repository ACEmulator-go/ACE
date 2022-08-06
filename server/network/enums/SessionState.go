package enums

type SessionState int32

const (
	SessionStateAuthLoginRequest    SessionState = SessionState(0)
	SessionStateAuthConnectResponse SessionState = SessionState(1)
	SessionStateAuthConnected       SessionState = SessionState(2)
	SessionStateWorldConnected      SessionState = SessionState(3)
	SessionStateTerminationStarted  SessionState = SessionState(4)
)
