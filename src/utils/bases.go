/*
* This file implements pair transformations between (Hex, Base64) and Bytes
*/

package utils


const HEX_ALPHABET string = "0123456789abcdef"
const BASE64_ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"


func HexToBytes(hex string) []byte {
	bytes := make([]byte, (len(hex) + 1) / 2)

	for i:=0; i<len(hex); i+=1 {
		digit := hex[i]

		var digit_value byte
		for j:=0; j<len(HEX_ALPHABET); j+=1 {
			if digit == HEX_ALPHABET[j] {
				digit_value = byte(j)
			}
		}

		if i%2 == 0 {
			digit_value *= (1 << 4)
		}

		bytes[i/2] |= digit_value
	}

	return bytes
}


func BytesToHex(bytes []byte) string {
	chars := make([]byte, len(bytes) * 2)

	for i:=0; i<len(bytes); i+=1 {
		part1 := bytes[i] / (1 << 4)
		part2 := bytes[i] % (1 << 4)

		chars[2*i] = HEX_ALPHABET[part1]
		chars[2*i + 1] = HEX_ALPHABET[part2]
	}

	return string(chars)
}


func Base64ToBytes(base64 string) []byte {
	bytes := make([]byte, len(base64) * 6 / 8)
	padding := 0

	for i:=0; i<len(base64); i+=1 {
		ch := base64[i]

		if ch == byte('=') {
			padding += 1
		}

		var base64Bits byte
		for j:=0; j<len(BASE64_ALPHABET); j+=1 {
			if ch == BASE64_ALPHABET[j] {
				base64Bits = byte(j)
			}
		}

		bits := i * 6
		byteIndex := bits / 8
		byteParity := bits % 8

		switch byteParity {
			case 0:
				bytes[byteIndex] |= base64Bits << 2
			case 2:
				bytes[byteIndex] |= base64Bits
			case 4:
				part1 := base64Bits / (1 << 2)
				part2 := base64Bits % (1 << 2)

				bytes[byteIndex] |= part1
				bytes[byteIndex + 1] |= part2 << 6
			case 6:
				part1 := base64Bits / (1 << 4)
				part2 := base64Bits % (1 << 4)

				bytes[byteIndex] |= part1
				bytes[byteIndex + 1] |= part2 << 4
		}
	}

	switch padding {
		case 1:
			bytes = bytes[0:len(bytes) - 1]
		case 2:
			bytes = bytes[0:len(bytes) - 2]
	}

	return bytes
}


func BytesToBase64(bytes []byte) string {
	bits := len(bytes) * 8
	padding := 0

	for bits%6 != 0 {
		bits += 8
		padding += 1
	}

	chars := make([]byte, bits/6)

	for i:=0; i<bits/6; i+=1 {
		prevBits := i * 6
		byteIndex := prevBits / 8
		byteParity := prevBits % 8

		var base64Bits byte
		if byteIndex < len(bytes) {
			switch byteParity {
				case 0:
					base64Bits = bytes[byteIndex] / (1 << 2)
				case 2:
					base64Bits = bytes[byteIndex] % (1 << 6)
				case 4:
					part1 := bytes[byteIndex] % (1 << 4)
					part2 := byte(0)
					if byteIndex + 1 < len(bytes) {
						part2 = bytes[byteIndex+ 1] / (1 << 6)
					}

					base64Bits = part1 * (1 << 2) + part2
				case 6:
					part1 := bytes[byteIndex] % (1 << 2)
					part2 := byte(0)
					if byteIndex + 1 < len(bytes) {
						part2 = bytes[byteIndex + 1] / (1 << 4)
					}

					base64Bits = part1 * (1 << 4) + part2
			}
		}

		chars[i] = BASE64_ALPHABET[base64Bits]
	}

	for i:=1; i<=padding; i+=1 {
		chars[len(chars) - i] = byte('=')
	}

	return string(chars)
}
