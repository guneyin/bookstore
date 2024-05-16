package common

import "time"

var (
	gitCommit string
	lastRun   time.Time
)

func GetVersion() string {
	return gitCommit
}

func SetLastRun(t time.Time) {
	lastRun = t
}

func GetLastRun() time.Time {
	return lastRun
}
