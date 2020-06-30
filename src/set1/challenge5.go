package set1


func RepeatingKeyXOR(data []byte, key []byte) []byte {
  encrypted := make([]byte, len(data))

  for i:=0; i<len(data); i+=1 {
    encrypted[i] = data[i] ^ key[i%len(key)]
  }

  return encrypted
}
