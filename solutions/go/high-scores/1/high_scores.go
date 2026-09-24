package highscores

import "slices"

type HighScores struct{
    	scores []int
	}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return slices.Max(s.scores)
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	sorted := slices.Clone(s.scores)
	slices.SortFunc(sorted, func(a, b int) int { return b - a }) // по убыванию
	if len(sorted) > 3 {
		return sorted[:3]
	}
	return sorted
}
