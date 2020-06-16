package main

import "bytes"


func HexToBuffer(hex string) *bytes.Buffer {
  raw := make([]byte, (len(hex) + 1)/2)

  var word byte = 0
  for i:=0; i<len(hex); i += 1 {
    // Hex to byte
    if byte('0') <= hex[i] && hex[i] <= byte('9') {
      word = hex[i] - byte('0')
    }

    if byte('a') <= hex[i] && hex[i] <= byte('f') {
      word = hex[i] - byte('a') + 10
    }

    // Append byte in the correct order
    if i % 2 == 0 {
      word = word << 4
    }

    raw[i/2] |= word
  }

  return bytes.NewBuffer(raw)
}


func BufferToHex(buffer *bytes.Buffer) string {
  var base string = "0123456789abcdef"
  hex := make([]byte, buffer.Len()*2)
  ix := 0

  for buffer.Len() > 0 {
    word, _ := buffer.ReadByte()
    h1 := word / (1 << 4)
    h2 := word % (1 << 4)
    hex[ix] = base[h1]
    hex[ix + 1] = base[h2]
    ix += 2
  }

  return string(hex)
}


func FixedXOR(data_hex string, chiper_hex string) string {
  data := HexToBuffer(data_hex)
  chiper := HexToBuffer(chiper_hex)
  xor := bytes.NewBufferString("")

  for data.Len() > 0 {
    a, _ := data.ReadByte()
    b, _ := chiper.ReadByte()

    xor.WriteByte(a ^ b)
  }

  return BufferToHex(xor)
}
