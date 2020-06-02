package main


func hexToIndex(data byte) byte {
  if byte('0') <= data && data <= byte('9') {
    return data - byte('0')
  }

  if byte('a') <= data && data <= byte('f') {
    return data - byte('a') + 10
  }

  return 0
}


func HexToBase64(hex string) string {
  bytes := make([]byte, len(hex)/2)
  const base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

  for i:=0; i<len(bytes); i += 1 {
    ord1 := hexToIndex(hex[2*i])
    ord2 := hexToIndex(hex[2*i + 1])
    bytes[i] = ord1 * 16 + ord2
  }

  output := make([]byte, len(hex)*4/6)

  for i:=0; i<len(bytes)*8/6; i+=1 {
    bit := i * 6
    index := bit / 8
    carry := bit % 8

    var b64i byte
    switch carry {
    case 0:
      b64i = bytes[index] >> 2
    case 2:
      b64i = bytes[index] % 0b1000000
    case 4:
      b64i = (bytes[index] % 0b10000) << 2 + bytes[index + 1] >> 6
    case 6:
      b64i = (bytes[index] % 0b100) << 4 + bytes[index + 1] >> 4
    }

    output[i] = base[b64i]
  }

  return string(output)
}

