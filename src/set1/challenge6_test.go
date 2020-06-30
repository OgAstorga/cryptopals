package set1

import (
	"os"
	"bufio"
	"testing"
  "../utils"
)


func TestChallenge6(t *testing.T) {
  exceptedKey := "Terminator X: Bring the noise"

	file, err := os.Open("./challenge6.txt")
	if err != nil {
    t.Errorf("Could not open ./challenge6.txt")
	}
	defer file.Close()

  base64 := ""
	scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    base64 = base64 + scanner.Text()
  }

  bytes := utils.Base64ToBytes(base64)

  _, key := BreakRepeatingXOR(bytes)

  if string(key) != exceptedKey {
    t.Errorf("Excepted key is '%s'; got '%s'", exceptedKey, string(key))
  }
}
