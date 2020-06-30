package set1

import "math"


// from https://en.wikipedia.org/wiki/Letter_frequency
var engHistogram = [27]float64 {
  0.08497, 0.01492, 0.02202,
  0.04253, 0.11162, 0.02228,
  0.02015, 0.06094, 0.07546,
  0.00153, 0.01292, 0.04025,
  0.02406, 0.06749, 0.07507,
  0.01929, 0.00095, 0.07587,
  0.06327, 0.09356, 0.02758,
  0.00978, 0.02560, 0.00150,
  0.01994, 0.00077, 0.0,
}


func VectorDistance(a []int, b []int) float64 {
  // Assert len(a) == len(b)

  score := float64(0);

  for i:=0; i<len(a); i++ {
    delta := a[i] - b[i]
    score += float64(delta * delta)
  }

  return math.Sqrt(score)
}


func CalcHistogram(text string) []int {
  histogram := make([]int, 27)

  for i:=0; i<len(text); i+=1 {
    ch := text[i]

    if byte('A') <= ch && ch <= byte('Z') {
      ch += byte('a') - byte('A')
    }

    if byte('a') <= ch && ch <= byte('z') {
      histogram[ch - byte('a')] += 1
    } else if ch == byte(' ') || (byte('0') <= ch && ch <= byte('9')) {
      // Ignore spaces and numbers
    } else {
      // Penalize symbols
      histogram[26] += 1
    }
  }

  return histogram
}


func BreakSingleByteXOR(encrypted []byte) ([]byte, float64) {
  idealHistogram := make([]int, 27)
  for i:=0; i<27; i+=1 {
    idealHistogram[i] = int(math.Round(engHistogram[i] * float64(len(encrypted))))
  }

  var guess []byte
  var guess_score float64 = 1000000.0

  for key:=0; key<256; key+=1 {
    decrypted := make([]byte, len(encrypted))
    for i:=0; i<len(encrypted); i+=1 {
      decrypted[i] = encrypted[i] ^ byte(key)
    }

    score := VectorDistance(idealHistogram, CalcHistogram(string(decrypted)))

    if score < guess_score {
      guess = decrypted
      guess_score = score
    }
  }

  return guess, guess_score
}
