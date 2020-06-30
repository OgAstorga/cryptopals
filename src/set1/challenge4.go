package set1


func DetectSingleByteXOR(hexs [][]byte) ([]byte, float64) {
	guess, _, guess_score := BreakSingleByteXOR(hexs[0])

	for i:=1; i<len(hexs); i+=1 {
		decrypted, _,  score := BreakSingleByteXOR(hexs[i])

		if score < guess_score {
			guess, guess_score = decrypted, score
		}
	}

	return guess, guess_score
}
