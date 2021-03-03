package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	players      []string
	places       []int
	purses       []int
	inPenaltyBox []bool

	popQuestions     []string
	scienceQuestions []string
	sportsQuestions  []string
	rockQuestions    []string

	currentPlayer            int
	isGettingOutOfPenaltyBox bool
}

func NewGame(players []string) *Game {
	game := &Game{}
	for i := 0; i < 6; i++ {
		game.places = append(game.places, 0)
		game.purses = append(game.purses, 0)
		game.inPenaltyBox = append(game.inPenaltyBox, false)
	}

	for i := 0; i < 50; i++ {
		game.popQuestions = append(game.popQuestions,
			fmt.Sprintf("Pop Question %d\n", i))
		game.scienceQuestions = append(game.scienceQuestions,
			fmt.Sprintf("Science Question %d\n", i))
		game.sportsQuestions = append(game.sportsQuestions,
			fmt.Sprintf("Sports Question %d\n", i))
		game.rockQuestions = append(game.rockQuestions,
			game.CreateRockQuestion(i))
	}

	for _, player := range players {
		game.Add(player)
	}

	return game
}

func (g *Game) CreateRockQuestion(index int) string {
	return fmt.Sprintf("Rock Question %d\n", index)
}

func (g *Game) IsPlayable() bool {
	return g.howManyPlayers() >= 2
}

func (g *Game) howManyPlayers() int {
	return len(g.players)
}

func (g *Game) Add(playerName string) bool {
	g.players = append(g.players, playerName)
	g.places[g.howManyPlayers()] = 0
	g.purses[g.howManyPlayers()] = 0
	g.inPenaltyBox[g.howManyPlayers()] = false

	fmt.Printf("%s was added\n", playerName)
	fmt.Printf("They are player number %d\n", len(g.players))

	return true
}

func (g *Game) Roll(roll int) {
	fmt.Printf("%s is the current player\n", g.players[g.currentPlayer])
	fmt.Printf("They have rolled a %d\n", roll)

	if g.inPenaltyBox[g.currentPlayer] {
		if roll%2 != 0 {
			g.isGettingOutOfPenaltyBox = true

			fmt.Printf("%s is getting out of the penalty box\n", g.players[g.currentPlayer])
			g.places[g.currentPlayer] = g.places[g.currentPlayer] + roll
			if g.places[g.currentPlayer] > 11 {
				g.places[g.currentPlayer] = g.places[g.currentPlayer] - 12
			}

			fmt.Printf("%s's new location is %d\n", g.players[g.currentPlayer], g.places[g.currentPlayer])
			fmt.Printf("The category is %s\n", g.currentCategory())
			g.askQuestion()
		} else {
			fmt.Printf("%s is not getting out of the penalty box\n", g.players[g.currentPlayer])
			g.isGettingOutOfPenaltyBox = false
		}
	} else {
		g.places[g.currentPlayer] = g.places[g.currentPlayer] + roll
		if g.places[g.currentPlayer] > 11 {
			g.places[g.currentPlayer] = g.places[g.currentPlayer] - 12
		}

		fmt.Printf("%s's new location is %d\n", g.players[g.currentPlayer], g.places[g.currentPlayer])
		fmt.Printf("The category is %s\n", g.currentCategory())
		g.askQuestion()
	}
}

func (g *Game) askQuestion() {
	if g.currentCategory() == "Pop" {
		question := g.popQuestions[0]
		g.popQuestions = g.popQuestions[1:]
		fmt.Printf(question)
	}
	if g.currentCategory() == "Science" {
		question := g.scienceQuestions[0]
		g.scienceQuestions = g.scienceQuestions[1:]
		fmt.Printf(question)
	}
	if g.currentCategory() == "Sports" {
		question := g.sportsQuestions[0]
		g.sportsQuestions = g.sportsQuestions[1:]
		fmt.Printf(question)
	}
	if g.currentCategory() == "Rock" {
		question := g.rockQuestions[0]
		g.rockQuestions = g.rockQuestions[1:]
		fmt.Printf(question)
	}
}

func (g *Game) currentCategory() string {
	if g.places[g.currentPlayer] == 0 {
		return "Pop"
	}
	if g.places[g.currentPlayer] == 4 {
		return "Pop"
	}
	if g.places[g.currentPlayer] == 8 {
		return "Pop"
	}
	if g.places[g.currentPlayer] == 1 {
		return "Science"
	}
	if g.places[g.currentPlayer] == 5 {
		return "Science"
	}
	if g.places[g.currentPlayer] == 9 {
		return "Science"
	}
	if g.places[g.currentPlayer] == 2 {
		return "Sports"
	}
	if g.places[g.currentPlayer] == 6 {
		return "Sports"
	}
	if g.places[g.currentPlayer] == 10 {
		return "Sports"
	}
	return "Rock"
}

func (g *Game) WasCorrectlyAnswered() bool {
	if g.inPenaltyBox[g.currentPlayer] {
		if g.isGettingOutOfPenaltyBox {
			fmt.Println("Answer was correct!!!!")
			g.purses[g.currentPlayer]++
			fmt.Printf("%s now has %d Gold Coins.\n", g.players[g.currentPlayer], g.purses[g.currentPlayer])

			winner := g.didPlayerWin()
			g.currentPlayer++
			if g.currentPlayer == len(g.players) {
				g.currentPlayer = 0
			}

			return winner
		} else {
			g.currentPlayer++
			if g.currentPlayer == len(g.players) {
				g.currentPlayer = 0
			}
			return true
		}
	} else {

		fmt.Println("Answer was corrent!!!!")
		g.purses[g.currentPlayer]++
		fmt.Printf("%s now has %d Gold Coins.\n", g.players[g.currentPlayer], g.purses[g.currentPlayer])

		winner := g.didPlayerWin()
		g.currentPlayer++
		if g.currentPlayer == len(g.players) {
			g.currentPlayer = 0
		}

		return winner
	}

	return false
}

func (g *Game) didPlayerWin() bool {
	return !(g.purses[g.currentPlayer] == 6)
}

func (g *Game) WrongAnswer() bool {
	fmt.Println("Question was incorrectly answered")
	fmt.Printf("%s was sent to the penalty box\n", g.players[g.currentPlayer])
	g.inPenaltyBox[g.currentPlayer] = true

	g.currentPlayer++
	if g.currentPlayer == len(g.players) {
		g.currentPlayer = 0
	}

	return true
}

func main() {
	notAWinner := false

	game := NewGame([]string{"Chet", "Pat", "Sue"})

	rand.Seed(time.Now().UTC().UnixNano())

	for {
		game.Roll(rand.Intn(5) + 1)

		if rand.Intn(9) == 7 {
			notAWinner = game.WrongAnswer()
		} else {
			notAWinner = game.WasCorrectlyAnswered()

		}

		if !notAWinner {
			break
		}
	}
}
