package recorder

import (
	"testing"
)

const file = "/tmp/test.gob"

type User struct{
	Name string
}

func TestNewCasette(t *testing.T) {
	var datato = &User{"Donald"}
	var datafrom = new(User)

	c, err := NewCasette(file)
	c.Save(datato)

	err := Save(file, datato)
	Check(err)
}