package domain

type StatePlayer string

const (
	StatePlayerUnspecified StatePlayer = ""
	StatePlayerPlaying     StatePlayer = "PLAYING"
	StatePlayerSuccess     StatePlayer = "SUCCESS"
	StatePlayerFail        StatePlayer = "FAIL"
	StatePlayerDisqual     StatePlayer = "DISQUAL"
)
