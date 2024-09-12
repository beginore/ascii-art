package intern

import (
	"bufio"
	"os"
)

var result string

func MakeMap(font string) (map[rune][]string, error) {
	patch := "ascii-art/assets/" + font + ".txt"
	file, err := os.Open(patch)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := rune(31)
	ch := map[rune][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			dec++
		} else {
			ch[dec] = append(ch[dec], line)
		}
	}
	return ch, nil
}
