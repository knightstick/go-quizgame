package quizgame

// QuizGame allows answering the questions
type QuizGame struct{}

// Play takes a list of questions and allows playing the game
func (game *QuizGame) Play([]Question) {

}

// Score returns the user's current score
func (game *QuizGame) Score() int {
	return 0
}

// NumberOfQuestions returns the total number of questions in the game
func (game *QuizGame) NumberOfQuestions() int {
	return 0
}
