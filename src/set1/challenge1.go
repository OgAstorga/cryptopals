package main

import "../utils"


func HexToBase64(hex string) string {
  return utils.BytesToBase64(utils.HexToBytes(hex))
}
