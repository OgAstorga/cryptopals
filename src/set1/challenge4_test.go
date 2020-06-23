package set1

import (
	"os"
	"bufio"
	"testing"
)


func TestChallenge4(t *testing.T) {
	const excepted = "Now that the party is jumping\n"

	file, err := os.Open("./challenge4.txt")
	if err != nil {
    t.Errorf("Could not open ./challenge4.txt")
	}
	defer file.Close()

	var hexs []string 
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexs = append(hexs, scanner.Text())
	}

	answer, _ := DetectSingleByteXOR(hexs)

	if excepted != string(answer) {
    t.Errorf("Excepted '%s'; got '%s'", excepted, answer)
	}
}