package set1

import "../utils"

func FixedXOR(data_hex string, chiper_hex string) string {
  data := utils.HexToBytes(data_hex)
  chiper := utils.HexToBytes(chiper_hex)
  xor := make([]byte, len(data))

  // Assert len(data) == len(chiper)

  for i:=0; i<len(data); i+=1 {
    xor[i] = data[i] ^ chiper[i]
  }

  return utils.BytesToHex(xor)
}
