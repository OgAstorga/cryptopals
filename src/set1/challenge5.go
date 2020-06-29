package set1

import "../utils"


func RepeatingKeyXOR(data string, key string) string {
  encrypted := make([]byte, len(data))

  for i:=0; i<len(data); i+=1 {
    encrypted[i] = data[i] ^ key[i%len(key)]
  }

  return utils.BytesToHex(encrypted)
}
