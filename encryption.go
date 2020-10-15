package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"
)

type encriptado struct {
	resolve string
	encrypt string
}

var (
	t      []string
	c      []string
	en     []string
	des    []string
	qwerty []string = strings.Split("qwertyuiopasdfghjklñzxcvbnm QWERTYUIOPASDFGHJKLÑZXCVBNM1234567890!·$%&/()=#@", "")
)

func (e *encriptado) enAndDes(character string, seed int) {
	e.encrypt += cesar(character, seed)
	m := strings.Split(e.encrypt, "")
	e.resolve += resolveCesar(m[len(m)-1], seed)
}
func (e *encriptado) codeCon(text string, patron string) {

	text64 := b64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(text64)
	t = strings.Split(string(text64), "")
	c = strings.Split(patron, "")

	var i int = 0

	for i < len(t) { //text
		for x := 0; x < len(c); x++ { //password
			for y := 0; y < len(qwerty); y++ { // qwerty for return
				if c[x] == qwerty[y] {
					e.enAndDes(t[i%len(c)], y)
					i++
					fmt.Printf("y es %v\n", y)
					break

				}
			}

		}
		if i > len(t) {
			break
		}

	}
	fmt.Println(e.resolve)
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
				m += len(qwerty)
				fmt.Printf("m es %v\n", qwerty[m])
			}

			characterDes = qwerty[m]
		}
	}
	return characterDes
}
func inp() {
	e := encriptado{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("¿Que texto quiere encriptar?  ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("")
	fmt.Print("¿cual clave usara?")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	e.codeCon(text, password)
	fmt.Println(e.encrypt)
	resolve64, _ := b64.StdEncoding.DecodeString(e.resolve)
	fmt.Println(string(resolve64))
}
func main() {
	inp()
}
