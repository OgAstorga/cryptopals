package set1

import (
  "testing"
  "../utils"
)


func TestChallenge2(t *testing.T) {
  data := utils.HexToBytes("1c0111001f010100061a024b53535009181c")
  chiper := utils.HexToBytes("686974207468652062756c6c277320657965")
  output := "746865206b696420646f6e277420706c6179"

  xor := FixedXOR(data, chiper)
  if utils.BytesToHex(xor) != output {
    t.Errorf("Excepted '%s'; got '%s'", output, utils.BytesToHex(xor))
  }
}
