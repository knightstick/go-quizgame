package quizgame

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// QuizGame allows answering the questions
type QuizGame struct {
	in        *bufio.Scanner
	out       io.Writer
	score     int
	questions []Question
	timer     *time.Timer
}

// QuizTimer is used to set the max time for completing the quiz
type QuizTimer interface {
	Sleep()
}

// ExitCode is a status that can be used to check how the QuizGame finished, i.e. timing out or
// completing successfully
type ExitCode int

const (
	// Timeout is returned when the timer runs out of time before the game is
	// finished
	Timeout ExitCode = iota
	// GameFinished is returned when the game finishes as all questions have been
	// answered
	GameFinished
)

// NewQuizGame initialises a new QuizGame with the given input
func NewQuizGame(in io.Reader, out io.Writer, questions []Question, timer *time.Timer) *QuizGame {
	return &QuizGame{
		in:        bufio.NewScanner(in),
		out:       out,
		questions: questions,
		timer:     timer,
	}
}

// Play takes a list of questions and allows playing the game
func (game *QuizGame) Play() ExitCode {
	finished := make(chan bool)

	go func() {
		for idx, question := range game.questions {
			questionString := fmt.Sprintf("Problem #%d: %s = ", idx+1, question.Question)
			game.out.Write([]byte(questionString))

			answer := game.readLine()

			if game.correctAnswer(&question, answer) {
				game.score++
			}
		}
		finished <- true
	}()

	select {
	case <-game.timer.C:
		return Timeout
	case <-finished:
		return GameFinished
	}
}

// Score returns the user's current score
func (game *QuizGame) Score() int {
	return game.score
}

// NumberOfQuestions returns the total number of questions in the game
func (game *QuizGame) NumberOfQuestions() int {
	return len(game.questions)
}

func (game *QuizGame) readLine() string {
	game.in.Scan()
	return game.in.Text()
}

func (game *QuizGame) correctAnswer(question *Question, answer string) bool {
	return question.Answer == answer
}
