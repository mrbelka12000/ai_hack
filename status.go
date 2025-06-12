package ai_hack

type DialogStatus string

const (
	DialogStatusOpen       DialogStatus = "open"
	DialogStatusInProgress DialogStatus = "in_progress"
	DialogStatusClose      DialogStatus = "close"
)
