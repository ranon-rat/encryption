package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	qwerty  []string = strings.Split("qwertyuiopasdfghjklñzxcvbnm @!#$%&/()=|°", "")
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
			characterDes = qwerty[q-s]
		}
	}
	return characterDes
}
func (e *encriptado) enAndDes(character string, seed int) {
	fmt.Print(character)
	encript += cesar(character, seed)
	m := strings.Split(encript, "")
	resolve += resolveCesar(m[len(m)-1], seed)
}
func (e *encriptado) codeCon(text string, contraseña string) {
	t = strings.Split(text, "")
	c = strings.Split(contraseña, "")
	x := 0

	for x < len(text)%len(c) { //password

		for y := 0; y < len(qwerty); y++ { // qwerty for return
			for _, i := range t { //text
				if c[x%len(c)] == qwerty[y] {
					e.enAndDes(i, y)
				}
				valor, err := strconv.Atoi(c[x%len(c)])
				if err == nil {
					e.enAndDes(i, valor)
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
func (e *encriptado) unique() {
	///////////////////////////////////////FOR NOT REPEAT A VALUES///////////////////////////////////////////////
	occured := map[string]bool{}

	for _, eS := range e.encriptList {
		if occured[eS] != true {
			occured[eS] = true
			e.encriptList = append(e.encriptList, eS)

		}
	}

}

func main() {

	e := encriptado{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	e.codeCon(text, "aleluya puta 1aasserasddeinalsdeszasxdcfvgbvfcdxsdcfvbghnbgvfbgmxiiasetodas")
	fmt.Println(e.encriptList)
	fmt.Println(e.resolveList)
}
