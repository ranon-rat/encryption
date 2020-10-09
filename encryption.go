package encryption

import (
	"strings"
)

type encriptado struct {
	resolveList []string
	encryptList []string
}

var (
	resolve string = ""
	encrypt string = ""
	t       []string
	c       []string
	en      []string
	des     []string
	qwerty  []string = strings.Split("qwertyuiopasdfghjklñzxcvbnm 1234567890@!#$%&/()=|°", "")
)

func (e *encriptado) enAndDes(character string, seed int) {
	encrypt += cesar(character, seed)
	m := strings.Split(encrypt, "")
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
			e.encryptList = append(e.encryptList, encrypt)
			e.resolveList = append(e.resolveList, resolve)
			resolve = ""
			encrypt = ""

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
