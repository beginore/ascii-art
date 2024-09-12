package intern

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

var ValidHash = map[string]string{
	"standard":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

func IsASCII(s string) error {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > unicode.MaxASCII {
			return fmt.Errorf("not valid input, pleace use only ASCII symbols")
		}
	}
	return nil
}

func CheckFont(font string) error {
	path := "./assets/" + font + ".txt"

	currentHash, err := HashFile(path)
	if err != nil {
		return errors.New("Invalid style, valid styles: standard, shadow, thinkertoy")
	}

	if ValidHash[font] != currentHash {
		err := errors.New("file has been modified")
		return err
	}

	return nil
}

func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GetTerminalSize() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	parts := strings.Fields(string(out))
	if len(parts) != 2 {
		return 0, err
	}
	cols, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	return cols, nil
}

func CheckTerminalLength(text string) error {
	terminalWidth, err := GetTerminalSize()
	if err != nil {
		return err
	}
	arr := strings.Split(text, "\n")

	if len(arr[0]) > terminalWidth {
		return fmt.Errorf("invalid length")
	}
	return nil
}
