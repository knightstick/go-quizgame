package quizgame

import "time"

// WaitQuizTimer is a timer that waits the given duration
type WaitQuizTimer struct {
	SleepTime time.Duration
}

// Sleep waits for the duration
func (timer WaitQuizTimer) Sleep() {
	time.Sleep(timer.SleepTime)
}
