package set1

import (
  "testing"
  "../utils"
)


func TestChallenge3(t *testing.T) {
	const encrypted string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	const decrypred string = "Cooking MC's like a pound of bacon"
	
	output, key, _ := BreakSingleByteXOR(utils.HexToBytes(encrypted))
  if key != byte('X') {
    t.Errorf("Excepted '%c'; got '%s'", 'X', string(key))
  }

  if string(output) != decrypred {
    t.Errorf("Excepted '%s'; got '%s'", decrypred, output)
  }
}
