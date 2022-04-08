package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = generateName()
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, 2)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s%.3d", b, rand.Intn(999))
}
