package code394

type Alphabet []string

// Digits returns an Alphabet consisting of the digits 0-9.
func Digits() Alphabet {
	return Alphabet{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}

type Condition interface {
	IsValid(code []string) bool
}

type Problem struct {
	Alphabet   Alphabet
	Length     int
	Conditions []Condition
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
		if alphabetIndizes[0] > len(p.Alphabet) {
			return nil
		}
	}
}

type PlacementCondition struct {
	Combination           []string
	CorrectAndWellPlaced  int
	CorrectAndWrongPlaced int
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
