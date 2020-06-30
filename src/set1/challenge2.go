package set1


func FixedXOR(data []byte, chiper []byte) []byte {
  xor := make([]byte, len(data))

  // Assert len(data) == len(chiper)

  for i:=0; i<len(data); i+=1 {
    xor[i] = data[i] ^ chiper[i]
  }

  return xor
}
