package main

import "testing"


func TestSec1(t *testing.T) {
  input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6dk"
  output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

  base64 := HexToBase64(input)
  if base64 != output {
    t.Errorf("Excepted '%s'; got '%s'", output, base64)
  }
}
