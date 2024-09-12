package intern

import (
	"strings"
)

func PrintArt(userInput string, LettersMap map[rune][]string) string {
	userInput = strings.ReplaceAll(userInput, "\\n", "\n")
	words := strings.Split(userInput, "\n")
	if len(words) > len(userInput) {
		words = words[1:] // условие нужно для отрезания лишнего переноса если аргумент весь состоит только из переносов
	}
	var result string
	for _, oneword := range words {
		if oneword == "" { // for \n\n\ cases
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range oneword {
				if art, ok := LettersMap[rune(ch)]; ok {
					result += art[i]
				}
			}
			result += "\n"
		}
	}
	return result
}
