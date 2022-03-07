package main

import (
	"fmt"
	"math/rand"
	"time"
)

const PROBABILITY_GOAL = 0.0001
const PROBABILITY_FIRST_TEAM_GOAL = 0.55
const STAMPS_NUMBER = 3000

type ScoreStamp map[int]Score

type Score struct {
	Home int
	Away int
}

func main() {
	score := fillScores()
	scores := getScore(score, 5000)
	fmt.Println(scores)
}
func fillScores() ScoreStamp {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	scores := make(ScoreStamp)

	for i := 0; i < STAMPS_NUMBER; i++ {
		scoreChanged := random.Float32() < PROBABILITY_GOAL
		home := 0
		away := 0
		if scoreChanged {
			if random.Float32() < PROBABILITY_FIRST_TEAM_GOAL {
				home = 1
				away = 0
			} else {
				home = 0
				away = 1
			}
		}
		var prevScore Score
		if len(scores) == 0 {
			prevScore = Score{
				Home: 0,
				Away: 0,
			}
		} else {
			prevScore = scores[i-1]
		}
		newScore := Score{
			Home: prevScore.Home + home,
			Away: prevScore.Away + away,
		}
		scores[i] = newScore
	}
	return scores
}

func getScore(scores ScoreStamp, offset int) Score {
	return scores[offset]
}
