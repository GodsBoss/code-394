package code394

import (
	"encoding/json"
	"fmt"
)

type Alphabet []string

// Digits returns an Alphabet consisting of the digits 0-9.
func Digits() Alphabet {
	return Alphabet{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

type Condition interface {
	IsValid(code []string) bool
}

type Problem struct {
	Alphabet   Alphabet    `json:"alphabet"`
	Length     int         `json:"length"`
	Conditions []Condition `json:"conditions"`
}

func (p Problem) Solve() []string {
	alphabetIndizes := make([]int, p.Length)
	for {
		currentCode := make([]string, p.Length)
		for i, index := range alphabetIndizes {
			currentCode[i] = p.Alphabet[index]
		}
		ok := true
		for _, condition := range p.Conditions {
			if !condition.IsValid(currentCode) {
				ok = false
				break
			}
		}
		if ok {
			return currentCode
		}
		alphabetIndizes[len(alphabetIndizes)-1]++
		for currentAlphabetIndex := len(alphabetIndizes) - 1; currentAlphabetIndex > 0; currentAlphabetIndex-- {
			if alphabetIndizes[currentAlphabetIndex] >= len(p.Alphabet) {
				alphabetIndizes[currentAlphabetIndex] = 0
				alphabetIndizes[currentAlphabetIndex-1]++
				continue
			}
			break
		}
		if alphabetIndizes[0] >= len(p.Alphabet) {
			return nil
		}
	}
}

func (p *Problem) UnmarshalJSON(data []byte) error {
	temp := struct {
		Alphabet   Alphabet        `json:"alphabet"`
		Length     int             `json:"length"`
		Conditions []jsonCondition `json:"conditions"`
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	for i := range temp.Conditions {
		if temp.Conditions[i].Type != ConditionTypePlacement {
			return fmt.Errorf("unknown condition type '%s'", temp.Conditions[i].Type)
		}
	}
	p.Alphabet = temp.Alphabet
	p.Length = temp.Length
	p.Conditions = make([]Condition, len(temp.Conditions))
	for i := range temp.Conditions {
		tempC := temp.Conditions[i]
		p.Conditions[i] = PlacementCondition{
			Combination:           tempC.Combination,
			CorrectAndWellPlaced:  tempC.CorrectAndWellPlaced,
			CorrectAndWrongPlaced: tempC.CorrectAndWrongPlaced,
		}
	}
	return nil
}

const (
	ConditionTypePlacement = "placement"
)

type jsonCondition struct {
	Type                  string   `json:"type"`
	Combination           []string `json:"combination"`
	CorrectAndWellPlaced  int      `json:"correct_and_well_placed,omitempty"`
	CorrectAndWrongPlaced int      `json:"correct_and_wrong_placed,omitempty"`
}

type PlacementCondition struct {
	Combination           []string `json:"combination"`
	CorrectAndWellPlaced  int      `json:"correct_and_well_placed,omitempty"`
	CorrectAndWrongPlaced int      `json:"correct_and_wrong_placed,omitempty"`
}

func (pc PlacementCondition) IsValid(code []string) bool {
	correctAndWellPlaced := 0
	correctAndWrongPlaced := 0
	for i := range pc.Combination {
		for j := range code {
			if pc.Combination[i] == code[j] {
				if i == j {
					correctAndWellPlaced++
				} else {
					correctAndWrongPlaced++
				}
				break
			}
		}
	}
	return correctAndWellPlaced == pc.CorrectAndWellPlaced && correctAndWrongPlaced == pc.CorrectAndWrongPlaced
}
