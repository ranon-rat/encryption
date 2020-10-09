package main

import (
	"fmt"
	"strings"
)

type encriptado struct {
	resolveList []string
	encriptList []string
}

var (
	resolve string = ""
	encript string = ""
	t       []string
	c       []string
	en      []string
	des     []string
	qwerty  []string = strings.Split("qwertyuiopasdfghjklñzxcvbnm 1234567890@!#$%&/()=|°", "")
)

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
	for q := 0; q < len(qwerty); q++ {
		if qwerty[q] == character {
			for q < s {
				if q == 0 {
					s--
				} else {
					s -= q
				}
			}
			s %= len(qwerty)
			characterDes = qwerty[q-s]
		}
	}
	return characterDes
}
func (e *encriptado) enAndDes(character string, seed int) {
	encript += cesar(character, seed)
	m := strings.Split(encript, "")
	resolve += resolveCesar(m[len(m)-1], seed)
}
func (e *encriptado) codeCon(text string, patron string) {
	t = strings.Split(text, "")
	c = strings.Split(patron, "")
	x := 0
	for x < len(c) { //password
		for i := 0; i < len(t); i++ { //text
			for y := 0; y < len(qwerty); y++ { // qwerty for return
				if c[x] == qwerty[y] {
					e.enAndDes(t[i], y)
				}
			}
			e.encriptList = append(e.encriptList, encript)
			e.resolveList = append(e.resolveList, resolve)
			resolve = ""
			encript = ""

		}
		x++
	}
}

func unique(text []string) []string {
	occured := map[string]bool{}
	t := []string{}
	for _, eS := range text {
		if occured[eS] != true {
			occured[eS] = true
			t = append(t, eS)
		}
	}
	return t
}
func main() {
	e := encriptado{}
	m := " puta puta ptua puta puta puta"
	e.codeCon(m, "scidroid es homosexual")
	fmt.Println(e.encriptList)
	fmt.Println(e.resolveList)
}
