package main


func hexToIx(hex byte) (byte) {
  if byte('0') <= hex && hex <= byte('9') {
    return hex - byte('0')
  }

  if byte('a') <= hex && hex <= byte('f') {
    return hex - byte('a') + 10
  }

  return hex
}

func hexToBytes(hex string) ([]byte) {
  bytes := make([]byte, len(hex) / 2)
  for i := 0; i < len(hex) / 2; i += 1 {
    bytes[i] = hexToIx(hex[2*i]) * 16 + hexToIx(hex[2*i + 1])
  }
  return bytes
}

func bytesToHex(bytes []byte) (string) {
  const base = "0123456789abcdef"

  output := make([]byte, len(bytes) * 2)
  for i := 0; i < len(bytes) * 2; i += 1 {
    bit := i * 4
    index := bit / 8
    carry := bit % 8

    var char byte
    switch carry {
    case 0:
      char = bytes[index] >> 4
    case 4:
      char = bytes[index] % 0b10000
    }

    output[i] = base[char]
  }

  return string(output)
}

func FixedXOR(hex, xor string) (string) {
  word1 := hexToBytes(hex)
  word2 := hexToBytes(xor)
  word3 := make([]byte, len(word1))

  for i:=0; i<len(word1); i+=1 {
    word3[i] = word1[i] ^ word2[i]
  }

  return bytesToHex(word3)
}
