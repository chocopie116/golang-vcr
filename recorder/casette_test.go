package recorder

import (
	"testing"
	"fmt"
)

func TestNewCasette(t *testing.T) {
	a := NewCasette("tmp/fuga.gob")
	fmt.Printf("%+v", a)

}