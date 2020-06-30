package set1

import (
	"os"
	"bufio"
	"testing"
  "../utils"
)


func TestChallenge4(t *testing.T) {
	const excepted = "Now that the party is jumping\n"

	file, err := os.Open("./challenge4.txt")
	if err != nil {
    t.Errorf("Could not open ./challenge4.txt")
	}
	defer file.Close()

	var words [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, utils.HexToBytes(scanner.Text()))
	}

	answer, _ := DetectSingleByteXOR(words)

	if excepted != string(answer) {
    t.Errorf("Excepted '%s'; got '%s'", excepted, answer)
	}
}
