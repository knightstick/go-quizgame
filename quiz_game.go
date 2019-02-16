package quizgame

import (
	"bufio"
	"fmt"
	"io"
)

// QuizGame allows answering the questions
type QuizGame struct {
	in        *bufio.Scanner
	out       io.Writer
	score     int
	questions []Question
	timer     QuizTimer
}

// QuizTimer is used to set the max time for completing the quiz
type QuizTimer interface {
	Sleep()
}

// NewQuizGame initialises a new QuizGame with the given input
func NewQuizGame(in io.Reader, out io.Writer, questions []Question, timer QuizTimer) *QuizGame {
	return &QuizGame{
		in:        bufio.NewScanner(in),
		out:       out,
		questions: questions,
		timer:     timer,
	}
}

// Play takes a list of questions and allows playing the game
func (game *QuizGame) Play() {
	finished := make(chan bool)

	go func() {
		game.timer.Sleep()
		finished <- true
	}()

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

	done := <-finished
	_ = done
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
