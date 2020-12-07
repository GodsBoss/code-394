package main

import (
	"fmt"
	"strings"

	"github.com/GodsBoss/code-394/pkg/code394"
)

func main() {
	problem := code394.Problem{
		Alphabet: code394.Digits(),
		Length:   3,
		Conditions: []code394.Condition{
			code394.PlacementCondition{
				Combination:          []string{"2", "9", "1"},
				CorrectAndWellPlaced: 1,
			},
			code394.PlacementCondition{
				Combination:           []string{"2", "4", "5"},
				CorrectAndWrongPlaced: 1,
			},
			code394.PlacementCondition{
				Combination:           []string{"4", "6", "3"},
				CorrectAndWrongPlaced: 2,
			},
			code394.PlacementCondition{
				Combination: []string{"5", "7", "8"},
			},
			code394.PlacementCondition{
				Combination:           []string{"5", "6", "9"},
				CorrectAndWrongPlaced: 1,
			},
		},
	}
	solution := problem.Solve()
	if solution == nil {
		fmt.Println("No solution found.")
	}
	fmt.Printf("Solution: %s\n", strings.Join(solution, " "))
}
