package domain

type StatePlayer string

const (
	StatePlayerPlaying StatePlayer = ""
	StatePlayerSuccess StatePlayer = "SUCCESS"
	StatePlayerFail    StatePlayer = "FAIL"
	StatePlayerDisqual StatePlayer = "DISQUAL"
)
