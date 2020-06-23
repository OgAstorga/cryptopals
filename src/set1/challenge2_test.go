package set1

import "testing"


func TestChallenge2(t *testing.T) {
  data := "1c0111001f010100061a024b53535009181c"
  chiper := "686974207468652062756c6c277320657965"
  output := "746865206b696420646f6e277420706c6179"

  xor := FixedXOR(data, chiper)
  if xor != output {
    t.Errorf("Excepted '%s'; got '%s'", output, xor)
  }
}
