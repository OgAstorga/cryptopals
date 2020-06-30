package set1


func HammingDistance(A []byte, B []byte) int {
  xor := FixedXOR(A, B)

  bits := 0
  for i:=0; i<len(xor); i+=1 {
    n := xor[i]
    for n > 0 {
      bits += 1
      n -= (n & -n)
    }
  }

  return bits
}


func KeySizeScore(data []byte, keysize int) float64 {
  var sum int
  var count int
  var i int

  for (i+2)*keysize <= len(data) {
    slice1 := data[i*keysize:(i+1)*keysize]
    slice2 := data[(i+1)*keysize:(i+2)*keysize]

    sum += HammingDistance(slice1, slice2)
    count += 1
    i += 1
  }

  return float64(sum) / float64(count)
}


func FindKeySize(data []byte) (int, float64) {
  var key int
  var score float64

  key, score = 1, KeySizeScore(data, 1)

  for i:=1; i<=50 && i<=len(data)/2; i+=1 {
    key_score := KeySizeScore(data, i)

    if key_score / float64(i) < score {
      key = i
      score = key_score / float64(i)
    }
  }

  return key, score
}

func BytesToBlocks(data []byte, size int) [][]byte {
  blocks := make([][]byte, size)

  for i:=0; i<len(data); i+=1 {
    blocks[i%size] = append(blocks[i%size], data[i])
  }

  return blocks
}

func BreakRepeatingXOR(data []byte) ([]byte, []byte) {
  keySize, _ := FindKeySize(data)

  blocks := BytesToBlocks(data, keySize)

  key := make([]byte, keySize)
  for i:=0; i<len(blocks); i+=1 {
    _, keyChar, _ := BreakSingleByteXOR(blocks[i])
    key[i] = keyChar
  }

  return RepeatingKeyXOR(data, key), key
}
