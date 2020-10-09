package encryption

func cesar(character string, seed int) string {
	var characterEn string
	for i := 0; i < len(qwerty); i++ {
		if qwerty[i] == character {
			characterEn = (qwerty[(i+seed)%len(qwerty)]) //a ,s d ,f . if the seed is 3 return you f
		}
	}

	return characterEn
}
func resolveCesar(character string, seed int) string { //cifrade ceaser
	s := seed
	var characterDes string
	for i := 0; i < len(qwerty); i++ {
		if qwerty[i] == character {
			m := i - s
			if m < 0 {
				m *= -1
			}
			characterDes = qwerty[m]
		}
	}
	return characterDes
}
