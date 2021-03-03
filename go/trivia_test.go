package main

import (
	"math/rand"
	"testing"
	"time"

	capturer "github.com/kami-zh/go-capturer"
)

var diceArr = []int{4, 1, 6, 3, 4, 5, 5, 2, 4, 2, 2, 3, 1, 3, 3, 2, 3, 1, 1, 2, 6, 1, 2, 4, 3, 3, 3, 3, 3, 1, 2, 2, 2, 6, 4, 1, 2, 4, 3, 3, 6, 5, 1, 6, 3, 1, 5, 2, 1, 1}

var goldenTest = `Chet was added
They are player number 1
Pat was added
They are player number 2
Sue was added
They are player number 3
Chet is the current player
They have rolled a 0
Chet's new location is 0
The category is Pop
Pop Question 0
Question was incorrectly answered
Chet was sent to the penalty box
Pat is the current player
They have rolled a 1
Pat's new location is 1
The category is Science
Science Question 0
Answer was corrent!!!!
Pat now has 1 Gold Coins.
Sue is the current player
They have rolled a 2
Sue's new location is 2
The category is Sports
Sports Question 0
Answer was corrent!!!!
Sue now has 1 Gold Coins.
Chet is the current player
They have rolled a 3
Chet is getting out of the penalty box
Chet's new location is 3
The category is Rock
Rock Question 0
Answer was correct!!!!
Chet now has 1 Gold Coins.
Pat is the current player
They have rolled a 4
Pat's new location is 5
The category is Science
Science Question 1
Answer was corrent!!!!
Pat now has 2 Gold Coins.
Sue is the current player
They have rolled a 5
Sue's new location is 7
The category is Rock
Rock Question 1
Question was incorrectly answered
Sue was sent to the penalty box
Chet is the current player
They have rolled a 6
Chet is not getting out of the penalty box
Pat is the current player
They have rolled a 7
Pat's new location is 0
The category is Pop
Pop Question 1
Answer was corrent!!!!
Pat now has 3 Gold Coins.
Sue is the current player
They have rolled a 8
Sue is not getting out of the penalty box
Chet is the current player
They have rolled a 9
Chet is getting out of the penalty box
Chet's new location is 0
The category is Pop
Pop Question 2
Answer was correct!!!!
Chet now has 2 Gold Coins.
Pat is the current player
They have rolled a 10
Pat's new location is 10
The category is Sports
Sports Question 1
Question was incorrectly answered
Pat was sent to the penalty box
Sue is the current player
They have rolled a 11
Sue is getting out of the penalty box
Sue's new location is 6
The category is Sports
Sports Question 2
Answer was correct!!!!
Sue now has 2 Gold Coins.
Chet is the current player
They have rolled a 12
Chet is not getting out of the penalty box
Pat is the current player
They have rolled a 13
Pat is getting out of the penalty box
Pat's new location is 11
The category is Rock
Rock Question 2
Answer was correct!!!!
Pat now has 4 Gold Coins.
Sue is the current player
They have rolled a 14
Sue is not getting out of the penalty box
Chet is the current player
They have rolled a 15
Chet is getting out of the penalty box
Chet's new location is 3
The category is Rock
Rock Question 3
Question was incorrectly answered
Chet was sent to the penalty box
Pat is the current player
They have rolled a 16
Pat is not getting out of the penalty box
Sue is the current player
They have rolled a 17
Sue is getting out of the penalty box
Sue's new location is 11
The category is Rock
Rock Question 4
Answer was correct!!!!
Sue now has 3 Gold Coins.
Chet is the current player
They have rolled a 18
Chet is not getting out of the penalty box
Pat is the current player
They have rolled a 19
Pat is getting out of the penalty box
Pat's new location is 18
The category is Rock
Rock Question 5
Answer was correct!!!!
Pat now has 5 Gold Coins.
Sue is the current player
They have rolled a 20
Sue is not getting out of the penalty box
Question was incorrectly answered
Sue was sent to the penalty box
Chet is the current player
They have rolled a 21
Chet is getting out of the penalty box
Chet's new location is 12
The category is Rock
Rock Question 6
Answer was correct!!!!
Chet now has 3 Gold Coins.
Pat is the current player
They have rolled a 22
Pat is not getting out of the penalty box
Sue is the current player
They have rolled a 23
Sue is getting out of the penalty box
Sue's new location is 22
The category is Rock
Rock Question 7
Answer was correct!!!!
Sue now has 4 Gold Coins.
Chet is the current player
They have rolled a 24
Chet is not getting out of the penalty box
Pat is the current player
They have rolled a 25
Pat is getting out of the penalty box
Pat's new location is 31
The category is Rock
Rock Question 8
Question was incorrectly answered
Pat was sent to the penalty box
Sue is the current player
They have rolled a 26
Sue is not getting out of the penalty box
Chet is the current player
They have rolled a 27
Chet is getting out of the penalty box
Chet's new location is 27
The category is Rock
Rock Question 9
Answer was correct!!!!
Chet now has 4 Gold Coins.
Pat is the current player
They have rolled a 28
Pat is not getting out of the penalty box
Sue is the current player
They have rolled a 29
Sue is getting out of the penalty box
Sue's new location is 39
The category is Rock
Rock Question 10
Answer was correct!!!!
Sue now has 5 Gold Coins.
Chet is the current player
They have rolled a 30
Chet is not getting out of the penalty box
Question was incorrectly answered
Chet was sent to the penalty box
Pat is the current player
They have rolled a 31
Pat is getting out of the penalty box
Pat's new location is 50
The category is Rock
Rock Question 11
Answer was correct!!!!
Pat now has 6 Gold Coins.
`

func TestGoldenOne(t *testing.T) {

	out := capturer.CaptureStdout(func() {
		notAWinner := false

		game := NewGame([]string{"Chet", "Pat", "Sue"})

		rand.Seed(time.Now().UTC().UnixNano())

		for i := range diceArr {
			game.Roll(i)

			if i%5 == 0 {
				notAWinner = game.WrongAnswer()
			} else {
				notAWinner = game.WasCorrectlyAnswered()

			}

			if !notAWinner {
				break
			}
		}
	})

	if out != goldenTest {
		t.Errorf("Wanted %s but got %s", goldenTest, out)
	}

}
