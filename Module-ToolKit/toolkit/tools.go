package toolkit

import (
	"crypto/rand"
	"fmt"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDRFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Tools struct{}

func (t *Tools) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

func main() {
	fmt.Println("Hello World .... ")
}
